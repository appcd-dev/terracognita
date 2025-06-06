package google

import (
	"context"
	"fmt"

	"github.com/cycloidio/terracognita/cache"
	"github.com/cycloidio/terracognita/errcode"
	"github.com/cycloidio/terracognita/filter"
	"github.com/cycloidio/terracognita/log"
	"github.com/cycloidio/terracognita/provider"
	"github.com/hashicorp/go-cty/cty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfprovider "github.com/hashicorp/terraform-provider-google/google/provider"
	"github.com/hashicorp/terraform-provider-google/google/transport"
	googleapi "google.golang.org/api/googleapi"

	"github.com/pkg/errors"
)

// version of the Terraform provider, this is automatically changed with the 'make update-terraform-provider'
const version = "4.9.0"

// skippableCodes is a list of codes
// which won't make Terracognita failed
// but they will be printed on the output
// they are based on the err.Code() content
// of the GCP error
var skippableCodes = map[string]struct{}{
	"accessNotConfigured": struct{}{},
}

type google struct {
	tfGoogleClient interface{}
	tfProvider     *schema.Provider
	gcpr           *GCPReader

	cache cache.Cache
}

// NewProvider returns a Gooogle Provider
func NewProvider(ctx context.Context, maxResults uint64, project, region, credentials, accessToken string) (provider.Provider, error) {
	cfg := transport.Config{
		Credentials: credentials,
		Project:     project,
		Region:      region,
		AccessToken: accessToken,
	}

	// tfgoogle.ConfigureBasePaths(&cfg)
	transport.ConfigureBasePaths(&cfg)
	log.Get().Debug("loading TF client", "func", "google.NewProvider")
	if err := cfg.LoadAndValidate(ctx); err != nil {
		return nil, fmt.Errorf("could not initialize 'terraform/google.Config.LoadAndValidate()' because: %s", err)
	}

	tfp := tfprovider.Provider()
	tfp.SetMeta(&cfg)

	log.Get().Debug("loading GCP client", "func", "google.NewProvider")
	reader, err := NewGcpReader(ctx, maxResults, project, region, credentials)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize GCPReader: %v", err)
	}

	return &google{
		tfGoogleClient: &cfg,
		tfProvider:     tfp,
		gcpr:           reader,
		cache:          cache.New(),
	}, nil
}

func (g *google) HasResourceType(t string) bool {
	_, err := ResourceTypeString(t)
	return err == nil
}

func (g *google) Region() string                        { return g.tfGoogleClient.(*transport.Config).Region }
func (g *google) Project() string                       { return g.tfGoogleClient.(*transport.Config).Project }
func (g *google) String() string                        { return "google" }
func (g *google) TagKey() string                        { return "labels" }
func (g *google) Source() string                        { return "hashicorp/google" }
func (g *google) Version() string                       { return version }
func (g *google) Configuration() map[string]interface{} { return make(map[string]interface{}) }

func (g *google) ResourceTypes() []string {
	return ResourceTypeStrings()
}

func (g *google) Resources(ctx context.Context, t string, f *filter.Filter) ([]provider.Resource, error) {
	rt, err := ResourceTypeString(t)
	if err != nil {
		return nil, err
	}

	rfn, ok := resources[rt]
	if !ok {
		return nil, errors.Errorf("the resource %q it's not implemented", t)
	}

	resources, err := rfn(ctx, g, t, f)
	if err != nil {
		// we filter the error from GCP and return a custom error
		// type if it's an error that we want to skip
		// Remove all wrap layer to get the right type
		unwrapErr := err
		for errors.Unwrap(unwrapErr) != nil {
			unwrapErr = errors.Unwrap(unwrapErr)
		}
		if reqErr, ok := unwrapErr.(*googleapi.Error); ok {
			// https://pkg.go.dev/google.golang.org/api@v0.68.0/googleapi#Error
			for _, gerr := range reqErr.Errors {
				if _, ok := skippableCodes[gerr.Reason]; ok {
					return nil, fmt.Errorf("%w: %v", errcode.ErrProviderAPI, reqErr)
				}
			}

		}

		return nil, errors.Wrapf(unwrapErr, " ")
	}

	return resources, nil
}

func (g *google) TFClient() interface{} {
	return g.tfGoogleClient
}

func (g *google) TFProvider() *schema.Provider {
	return g.tfProvider
}
func (g *google) FixResource(t string, v cty.Value) (cty.Value, error) { return v, nil }
func (g *google) FilterByTags(tags interface{}) error                  { return nil }
