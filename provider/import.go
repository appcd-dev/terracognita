package provider

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/cycloidio/terracognita/errcode"
	"github.com/cycloidio/terracognita/filter"
	"github.com/cycloidio/terracognita/interpolator"
	"github.com/cycloidio/terracognita/log"
	"github.com/cycloidio/terracognita/writer"
	"github.com/pkg/errors"
)

func readResource(ctx context.Context,
	re Resource,
	t string,
	hcl, tfstate writer.Writer,
	interpolation *interpolator.Interpolator,
	f *filter.Filter,
	logger *slog.Logger,
) error {
	res, err := re.ImportState(ctx)
	if err != nil {
		return err
	}

	// If the InstanceState is nil after the ImportState it
	// means that nothing was imported (potentially is not even Importable)
	// so we have to skip the resource
	if re.InstanceState() == nil {
		return nil
	}

	// In case there is more than one State to import
	// we create a new slice with those elements and iterate
	// over it
	for _, r := range append([]Resource{re}, res...) {
		err = r.Read(ctx, f)
		if err != nil {
			// Errors are ignored. If a resource is invalid we assume it can be skipped, it can be related to inconsistencies in deployed resources.
			// So instead of failing and stopping execution we ignore them and continue (we log them if -v is specified)

			logger.Debug("error reading resource", "error", err)

			continue
		}

		if hcl != nil {
			logger.Debug("calculating HCL")
			err = r.HCL(hcl)
			if err != nil {
				return errors.Wrapf(err, "error while calculating the Config of resource %q", t)
			}
		}

		if tfstate != nil {
			logger.Debug("calculating TFState")
			err = r.State(tfstate)
			if err != nil {
				return errors.Wrapf(err, "error while calculating the state of resource %q", t)
			}
		}
		state := r.InstanceState()

		if state != nil {
			attributes, err := re.AttributesReference()
			if err != nil {
				return errors.Wrapf(err, "unable to fetch attributes of resource")
			}
			attrs := make(map[string]string)
			for _, attribute := range attributes {
				value, ok := state.Attributes[attribute]
				if !ok || len(value) == 0 {
					continue
				}
				attrs[attribute] = value
			}
			interpolation.AddResourceAttributes(fmt.Sprintf("%s.%s", r.Type(), r.Name()), attrs)
		}
	}
	return nil
}

func getResourceTypes(p Provider, f *filter.Filter) ([]string, error) {
	// Validate if the Exclude filter is right
	if len(f.Exclude) != 0 {
		for _, e := range f.Exclude {
			if !p.HasResourceType(e) {
				return nil, errors.Wrapf(errcode.ErrProviderResourceNotSupported, "type %s on Exclude filter", e)
			}
		}
	}
	types := make([]string, 0)
	if len(f.Targets) != 0 {
		typesWithIDs := f.TargetsTypesWithIDs()
		for k := range typesWithIDs {
			if !p.HasResourceType(k) {
				return nil, errors.Wrapf(errcode.ErrProviderResourceNotSupported, "type %s on Target filter", k)
			}
			types = append(types, k)
		}
		return types, nil
	}
	// Validate if the Include filter is right
	if len(f.Include) != 0 {
		for _, i := range f.Include {
			if !p.HasResourceType(i) {
				return nil, errors.Wrapf(errcode.ErrProviderResourceNotSupported, "type %s on Include filter", i)
			}
		}
		return f.Include, nil
	}
	return p.ResourceTypes(), nil

}

// GetResources returns all the resources of the Provider p filtered by f
func GetResources(ctx context.Context, p Provider, f *filter.Filter) (result []Resource, err error) {
	logger := log.Get().With("func", "provider.GetResources")
	types, err := getResourceTypes(p, f)
	if err != nil {
		return nil, err
	}
	typesWithIDs := f.TargetsTypesWithIDs()
	logger.Debug("current filter", "filters", f.String(), "types", types)
	resTypeErrGroup, rtCtx := errgroup.WithContext(ctx)
	resTypeErrGroup.SetLimit(runtime.NumCPU())
	for _, t := range types {
		t := t
		resTypeErrGroup.Go(func() error {
			logger := logger.With("resource", t)
			if f.IsExcluded(t) {
				logger.Debug("excluded")
				return nil
			}
			if len(typesWithIDs) != 0 {
				for _, ID := range typesWithIDs[t] {
					result = append(result, NewResource(ID, t, p))
				}
				return nil
			}

			logger.Debug("fetching the list of resources")

			resources, err := p.Resources(rtCtx, t, f)
			if err != nil {
				return fmt.Errorf("error while fetching the resources of type: %s: %w", t, err)
			}
			logger.Debug("fetched the list of resources", "count", len(resources))

			result = append(result, resources...)
			return nil
		})
	}
	err = resTypeErrGroup.Wait()
	if err != nil {
		return nil, fmt.Errorf("error while reading the resources: %w", err)
	}
	logger.Debug("Scanning done")
	return result, nil
}

// Import imports from the Provider p all the resources filtered by f and writes
// the result to the hcl or tfstate if those are not nil
func Import(ctx context.Context, p Provider, hcl, tfstate writer.Writer, f *filter.Filter, out io.Writer) error {
	logger := log.Get().With("func", "provider.Import")

	if err := f.Validate(); err != nil {
		return err
	}
	types, err := getResourceTypes(p, f)
	if err != nil {
		return err
	}

	typesWithIDs := f.TargetsTypesWithIDs()

	fmt.Fprintf(out, "Scanning with filters: %s", f)
	logger.Debug("current filter", "filters", f.String())

	interpolation := interpolator.New(p.String())
	resTypeErrGroup, rtCtx := errgroup.WithContext(ctx)
	resTypeErrGroup.SetLimit(3)
	for _, t := range types {
		t := t
		resTypeErrGroup.Go(func() error {
			logger := logger.With("resource", t)

			if f.IsExcluded(t) {
				logger.Debug("excluded")
				return nil
			}

			logger.Debug("fetching the list of resources")

			var resources []Resource

			if len(typesWithIDs) != 0 {
				for _, ID := range typesWithIDs[t] {
					resources = append(resources, NewResource(ID, t, p))
				}
			} else {
				resources, err = p.Resources(rtCtx, t, f)
				if err != nil {
					// we filter the error: if it's an error provider side, we continue
					// the import but we print the error.
					if errors.Is(err, errcode.ErrProviderAPI) {
						logger.Debug("unable to import resource", "error", err)
					} else if strings.Contains(err.Error(), "AccessDenied") {
						// skip access denied errors, since we might not have access to all resources when trying to import based on tags
						logger.Debug("unable to import resource, access denied", "error", err)
					} else {
						logger.Warn("failed to fetch the resource type", "error", err, "resource_type", t)
					}
				}
			}

			resourceLen := len(resources)
			resErrGroup, ectx := errgroup.WithContext(rtCtx)
			resErrGroup.SetLimit(runtime.NumCPU())
			for i, re := range resources {
				logger := logger.With("id", re.ID(), "total", resourceLen, "current", i+1)
				fmt.Fprintf(out, "\rScanning %s [%d/%d]", t, i+1, resourceLen)

				logger.Debug("reading from TF")
				resErrGroup.Go(func() error {
					return readResource(ectx, re, t, hcl, tfstate, interpolation, f, logger)
				})
			}
			err := resErrGroup.Wait()
			if err != nil {
				return fmt.Errorf("error while reading the resources of type: %s: %w", t, err)
			}
			if resourceLen > 0 {
				fmt.Fprintf(out, "\rScanning %s [%d/%d] Done!\n", t, resourceLen, resourceLen)
			}
			return nil
		})

	}
	err = resTypeErrGroup.Wait()
	if err != nil {
		return fmt.Errorf("error while reading the resources: %w", err)
	}
	logger.Debug("Scanning done")

	if hcl != nil {
		hcl.Interpolate(interpolation)
		fmt.Fprintf(out, "\rWriting HCL ...")
		logger.Debug("writing the HCL")

		err = hcl.Sync()
		if err != nil {
			return errors.Wrapf(err, "error while Sync Config")
		}

		fmt.Fprintf(out, "\rWriting HCL Done!\n")
		logger.Debug("writing the HCL done")
	}

	if tfstate != nil {
		tfstate.Interpolate(interpolation)
		fmt.Fprintf(out, "\rWriting TFState ...")
		logger.Debug("writing the TFState")

		err := tfstate.Sync()
		if err != nil {
			return errors.Wrapf(err, "error while Sync State")
		}

		fmt.Fprintf(out, "\rWriting TFState Done!\n")
		logger.Debug("writing the TFState done")
	}

	return nil
}
