package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/chr4/pwgen"
	"github.com/cycloidio/terracognita/errcode"
	"github.com/cycloidio/terracognita/filter"
	"github.com/cycloidio/terracognita/log"
	"github.com/cycloidio/terracognita/tag"
	"github.com/cycloidio/terracognita/util"
	"github.com/cycloidio/terracognita/writer"
	awsdocs "github.com/cycloidio/tfdocs/providers/aws"
	azuredocs "github.com/cycloidio/tfdocs/providers/azurerm"
	googledocs "github.com/cycloidio/tfdocs/providers/google"
	"github.com/cycloidio/tfdocs/providers/vsphere"
	tfdocs "github.com/cycloidio/tfdocs/resource"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform/providers"
	"github.com/hashicorp/terraform/states"
	"github.com/pascaldekloe/name"
	"github.com/pkg/errors"
)

//go:generate go tool mockgen -destination=../mock/resource.go -mock_names=Resource=Resource -package mock github.com/cycloidio/terracognita/provider Resource

// Resource represents the minimal information needed to
// define a Provider resource
type Resource interface {
	// ID is the ID of the Resource
	ID() string

	// Type is the type of resource (ex: aws_instance)
	Type() string

	// InstanceState is the Terraform state of the resource
	// it contains important elements like `Attributes`
	InstanceState() *terraform.InstanceState

	// Name is the resource name given by Terracognita
	Name() string

	// TFResource is the definition of that resource
	TFResource() *schema.Resource

	// SetImporter set schema.Resource.Importer
	// It defines the ResourceImporter implementation for this resource
	SetImporter(*schema.ResourceImporter)

	// Data is the actual data of the Resource
	Data() *schema.ResourceData

	// Provider is the Provider of that Resource
	Provider() Provider

	// ImportState imports the Resource state
	// to the Resource and could return []Resource if
	// it imported more than one state, this list does not
	// include the actual Resource on parameters, so if
	// len([]Resource) == 0 means only the Resource is imported
	ImportState(ctx context.Context) ([]Resource, error)

	// Read read the remote information of the Resource to the
	// state and calculates the ResourceInstanceObject
	Read(ctx context.Context, f *filter.Filter) error

	// State calculates the state of the Resource and
	// writes it to w
	State(w writer.Writer) error

	// HCL returns the HCL configuration of the Resource and
	// writes it to HCL
	HCL(w writer.Writer) error

	// InstanceInfo returns the InstanceInfo of this Resource
	InstanceInfo() *terraform.InstanceInfo

	// ImpliedType returns the cty.Type of the
	// Resource
	ImpliedType() cty.Type

	// ResourceInstanceObject is the calculated states.ResourceInstanceObject
	// after 'Read' has been called
	ResourceInstanceObject() *states.ResourceInstanceObject

	// AttributesReference return the list of possible value
	// to be interpolated with the resource
	AttributesReference() ([]string, error)

	// SetIgnoreTagFilter is to mark the resource as already filtered, so we do not try
	// to filter it again when reading it.
	// This is mostly used in case the API does not support filtering and we have to
	// do it manually
	SetIgnoreTagFilter(b bool)
}

// resources is a general implementation of Resource interface
// that fulfills all the usecases for Terracognita
type resource struct {
	id string

	// resourceType is the type of resource (ex: aws_instance)
	// as type is a reserved word resourceType is used on private
	resourceType string

	tfResource *schema.Resource

	data       *schema.ResourceData
	state      *terraform.InstanceState
	stateValue cty.Value

	provider Provider

	// The name it has on the config
	// so it can be the same on the HCL
	// and State
	configName string

	resourceInstanceObject *states.ResourceInstanceObject

	client *GRPCClient

	ignoreTagFilter bool
}

var (
	providerResources = map[string]getResource{
		"aws":     awsdocs.GetResource,
		"azurerm": azuredocs.GetResource,
		"google":  googledocs.GetResource,
		"vsphere": vsphere.GetResource,
	}
)

type getResource func(rt string) (*tfdocs.Resource, error)

// NewResource returns an implementation of the Resource
func NewResource(id, rt string, p Provider) Resource {
	return &resource{
		id:           id,
		resourceType: rt,
		provider:     p,
		client:       NewGRPCClient(p.TFProvider()),
	}
}

func (r *resource) AttributesReference() ([]string, error) {
	resourceFunc, ok := providerResources[r.provider.String()]
	if !ok {
		return nil, errors.New(fmt.Sprintf("provider %s is not supported", r.provider.String()))
	}
	res, err := resourceFunc(r.resourceType)
	if err != nil {
		return nil, errors.Wrap(err, "unable to extract provider attributes")
	}
	result := make([]string, len(res.Attributes))
	for i, attribute := range res.Attributes {
		result[i] = attribute.Name
	}
	// add "id" to the exported attributes
	result = append(result, "id")
	return result, nil
}

func (r *resource) ID() string { return r.id }

func (r *resource) Type() string { return r.resourceType }

func (r *resource) Name() string { return r.configName }

func (r *resource) InstanceState() *terraform.InstanceState { return r.state }

func (r *resource) TFResource() *schema.Resource {
	if r.tfResource != nil {
		return r.tfResource
	}

	tfr, ok := r.provider.TFProvider().ResourcesMap[r.resourceType]
	if !ok {
		panic(fmt.Sprintf("the resource %s does not exists on TF", r.resourceType))
	}
	r.tfResource = tfr
	return r.tfResource
}

func (r *resource) SetImporter(i *schema.ResourceImporter) {
	r.TFResource().Importer = i
}

func (r *resource) SetIgnoreTagFilter(b bool) {
	r.ignoreTagFilter = b
}

func (r *resource) Data() *schema.ResourceData {
	if r.data == nil {
		r.data = r.TFResource().Data(nil)
	}
	return r.data
}

func (r *resource) Provider() Provider { return r.provider }

func (r *resource) ImportState(ctx context.Context) ([]Resource, error) {
	logger := log.Get()
	// If it does not support import do not try
	if r.TFResource().Importer == nil {
		logger.Debug("This resource it's not Importable", "func", "ImportState", "resource", r.Type())
		return nil, nil
	}

	irsresp := r.client.ImportResourceState(ctx, ImportResourceStateRequest{
		TypeName: r.resourceType,
		ID:       r.id,
	})
	if err := irsresp.Diagnostics.Err(); err != nil {
		return nil, errors.Wrapf(err, "could not import resource %s with id %s", r.resourceType, r.id)
	}
	// This converts value to state so we can follow the 2 API
	newInstanceStates := make([]*terraform.InstanceState, 0, len(irsresp.ImportedResources))
	for _, ir := range irsresp.ImportedResources {
		is := terraform.NewInstanceStateShimmedFromValue(ir.State, r.TFResource().SchemaVersion)
		// It's not set from the NewInstanceStateShimmedFromValue
		// and we need it later on to create a new Resource
		// as this may be different from the current Resource
		is.Ephemeral = terraform.EphemeralState{
			Type: ir.TypeName,
		}
		newInstanceStates = append(newInstanceStates, is)
	}

	resources := make([]Resource, 0, len(newInstanceStates)-1)
	// We assume that the first state is the one for this Resource
	// so the following ones after the first one are from other types
	// and are, for that reason, other Resources which should be returned
	for i, is := range newInstanceStates {
		// copy the ID again just to be sure it wasn't missed
		is.Attributes["id"] = is.ID

		resourceType := is.Ephemeral.Type
		if resourceType == "" {
			resourceType = r.Type()
		}

		if i != 0 {
			res := NewResource(is.ID, resourceType, r.provider)
			res.(*resource).state = is
			resources = append(resources, res)
		} else {
			r.state = is
			r.stateValue = irsresp.ImportedResources[i].State
		}

	}

	return resources, nil
}

func (r *resource) Read(ctx context.Context, f *filter.Filter) error {
	var err error
	rrreq := ReadResourceRequest{
		TypeName:   r.Type(),
		PriorState: r.stateValue,
	}
	rrres := r.client.ReadResource(ctx, rrreq)
	if err = rrres.Diagnostics.Err(); err != nil {
		return errors.Wrapf(err, "could not read resource %s with id %s", r.resourceType, r.id)
	}

	// After getting the resource data we call the provider to fix any potential
	// issues it may have like invalid configurations
	newStateValue, err := r.Provider().FixResource(r.Type(), rrres.NewState)
	if err != nil {
		return errors.Wrapf(err, "failed to fix resource %s", r.Type())
	}

	newInstanceState := terraform.NewInstanceStateShimmedFromValue(newStateValue, r.TFResource().SchemaVersion)
	r.state = newInstanceState
	r.stateValue = newStateValue

	// The old provider API used an empty id to signal that the remote
	// object appears to have been deleted, but our new protocol expects
	// to see a null value (in the cty sense) in that case.
	if rrres.NewState.IsNull() || newInstanceState.ID == "" {
		return errors.Wrapf(errcode.ErrProviderResourceNotRead, "the resource %q with ID %q did not return an ID", r.resourceType, r.id)
	}

	r.data = r.TFResource().Data(r.state)

	// Some resources can not be filtered by tags,
	// so we have to do it manually
	// it's not all of them though.
	// We don't do it if explicitly set on
	// the resource to not do it as it means
	// it has been done already
	if !r.ignoreTagFilter {
		for _, t := range f.Tags {
			// Default match key
			if v, ok := r.data.GetOk(fmt.Sprintf("%s.%s", r.Provider().TagKey(), t.Name)); ok && v.(string) == t.Value {
				continue
			}

			// Check if the filter tag match any other tags found
			// https://github.com/cycloidio/terracognita/issues/223
			if v, ok := tag.GetOtherTags(r.Provider().String(), r.data, t); ok && v == t.Value {
				continue
			}

			return errors.WithStack(errcode.ErrProviderResourceDoNotMatchTag)
		}
	}

	// Filter out autogenerated resources from AWS
	if v, ok := r.data.GetOk(r.Provider().TagKey()); ok {
		if err = r.provider.FilterByTags(v); err != nil {
			return err
		}
	}

	meta, err := json.Marshal(r.state.Meta)
	if err != nil {
		return err
	}

	zstate, err := util.HashicorpToZclonfValue(newStateValue, r.tfResource.CoreConfigSchema().ImpliedType())
	if err != nil {
		return err
	}

	rio := providers.ImportedResource{
		TypeName: r.resourceType,
		Private:  meta,
		State:    zstate,
	}.AsInstanceObject()

	r.resourceInstanceObject = rio

	return nil
}

// State calculates the state of the Resource and
// writes it to w
func (r *resource) State(w writer.Writer) error {
	if importer := r.tfResource.Importer; importer != nil {
		// If it does not have any configName we will generate one
		// and store it, so net time it'll use that one on any config
		if r.configName == "" {
			configName := tag.GetNameFromTag(r.provider.TagKey(), r.data, r.id)
			if ok, err := w.Has(fmt.Sprintf("%s.%s", r.resourceType, configName)); err != nil {
				return err
			} else if ok {
				configName = pwgen.Alpha(5)
			}

			err := w.Write(fmt.Sprintf("%s.%s", r.resourceType, configName), r)
			if err != nil {
				return fmt.Errorf("error writing resource %s: %w", r.resourceType, err)
			}

			r.configName = configName
		} else {
			err := w.Write(fmt.Sprintf("%s.%s", r.resourceType, r.configName), r)
			if err != nil {
				return fmt.Errorf("error writing resource config: %s %s: %w", r.configName, r.resourceType, err)
			}

			return nil
		}
	}
	return nil
}

// HCL returns the HCL configuration of the Resource and
// writes it to HCL
func (r *resource) HCL(w writer.Writer) error {
	cfg := mergeFullConfig(r.data, r.tfResource.Schema, "")

	resourceFunc, ok := providerResources[r.provider.String()]
	if !ok {
		return errors.New(fmt.Sprintf("provider %s is not supported", r.provider.String()))
	}

	tfdoc, err := resourceFunc(r.Type())
	if err != nil {
		return errors.New(fmt.Sprintf("provider %s with resource %s is not supported on the docs", r.provider.String(), r.Type()))
	}
	// This will convert all Category into snake_case
	cfg[writer.ResourceCategoryKey] = strings.ToLower(name.Delimit(tfdoc.Category, '_'))

	// If it does not have any configName we will generate one
	// and store it, so net time it'll use that one on any config
	if r.configName == "" {
		configName := tag.GetNameFromTag(r.provider.TagKey(), r.data, r.id)
		if ok, err := w.Has(fmt.Sprintf("%s.%s", r.resourceType, configName)); err != nil {
			return err
		} else if ok {
			configName = pwgen.Alpha(5)
		}

		err := w.Write(fmt.Sprintf("%s.%s", r.resourceType, configName), cfg)
		if err != nil {
			return err
		}

		r.configName = configName
	} else {
		err := w.Write(fmt.Sprintf("%s.%s", r.resourceType, r.configName), cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *resource) InstanceInfo() *terraform.InstanceInfo {
	return &terraform.InstanceInfo{
		Id:   r.id,
		Type: r.resourceType,
	}
}

func (r *resource) ImpliedType() cty.Type {
	return r.tfResource.CoreConfigSchema().ImpliedType()
}

func (r *resource) ResourceInstanceObject() *states.ResourceInstanceObject {
	return r.resourceInstanceObject
}

// mergeFullConfig creates the key to the map and if it had a value before set it, if
func mergeFullConfig(cfgr *schema.ResourceData, sch map[string]*schema.Schema, key string) map[string]interface{} {
	res := make(map[string]interface{})
	// conflicts and exactly have all the possible conflicts that we can have at this level
	// of the schema, so we do not add a conflicted attribute after the other.
	// The structure is:
	// * key: Value that is Conflicted with
	// * value: Attribute that has Clonflicts
	conflicts := make(map[string]string)
	exactly := make(map[string]string)
	for k, v := range sch {
		// If it's just a Computed value, do not add it to the output
		if !isConfig(v) {
			continue
		}

		// Basically calculates the needed
		// key to the current access
		var kk string
		if key != "" {
			kk = key + "." + k
		} else {
			kk = k
		}

		// schema.Resource means that it has nested fields
		if sr, ok := v.Elem.(*schema.Resource); ok {
			// Example would be aws_security_group
			if v.Type == schema.TypeSet {
				s, ok := cfgr.GetOk(kk)
				// If the value is Required we need to add it
				// even if it's not sent
				if (!ok || s == nil) && !v.Required {
					continue
				}

				res[k] = normalizeSetList(sr.Schema, s.(*schema.Set).List())
			} else if v.Type == schema.TypeList {
				var ar interface{} = make([]interface{}, 0)

				l, ok := cfgr.GetOk(kk)
				if !ok {
					continue
				}

				list := l.([]interface{})
				for i := range list {
					fc := mergeFullConfig(cfgr, sr.Schema, fmt.Sprintf("%s.%d", kk, i))
					// It can be possible that there are no actual values insight, so to not add
					// an empty entity we validate it
					if len(fc) > 0 {
						ar = append(ar.([]interface{}), fc)
					}
				}

				// If no element on the ar, then we just
				// ignore it to not add empty values
				if len(ar.([]interface{})) > 0 {
					res[k] = ar
				}
			} else {
				res[k] = mergeFullConfig(cfgr, sr.Schema, kk)
			}
			// As it's a nested element it does not require any of
			// the other code as it's for single value schemas
			continue
		}

		// This sets the single values that we see on the
		// end result

		vv, ok := cfgr.GetOk(kk)
		// If the value is Required we need to add it
		// even if it's not sent
		if (!ok || vv == nil) && !v.Required {
			continue
		}

		// Format the ConflictsWith and ExactlyOneOf so they are "standard"
		conflictsWith := formatAttributePaths(v.ConflictsWith)
		exactlyOneOf := formatAttributePaths(v.ExactlyOneOf)

		// A value in which this one conflicts has been set before
		// so we should no set this one as it'll raise an error of
		// `conflicts with *` on Terraform
		if mapHasKeys(res, conflictsWith) {
			continue
		}
		if mapHasKeys(res, exactlyOneOf) {
			continue
		}

		// If any of the attributes added before has
		// a conflict with the current key we have to delete
		// the old one and use this one
		if c, ok := conflicts[k]; ok {
			delete(res, c)
		}
		if e, ok := exactly[k]; ok {
			delete(res, e)
		}

		// If the ConflictsWith has values we store them on the
		// conflicts map so none of those attributes is added after
		// this one has been added
		if len(conflictsWith) != 0 {
			for _, c := range conflictsWith {
				conflicts[c] = k
			}
		}
		if len(exactlyOneOf) != 0 {
			for _, e := range exactlyOneOf {
				exactly[e] = k
			}
		}

		// If it's a type map, we'll prefix the key with '=tc=' to know
		// that it's an attribute so we have to keep the '=' when formatting
		// the HCL
		if v.Type == schema.TypeMap {
			k = "=tc=" + k
		}
		if s, ok := vv.(*schema.Set); ok {
			res[k] = s.List()
		} else {
			res[k] = normalizeInterpolation(normalizeValue(vv))
		}
	}
	return res
}

// formatAttributePaths get's the last element of the string, the
// cws look like this sometimes '["name", "a.name", "a.0.name"]'
// and as we always need the "name" from all of them we just have
// to abstract the last element
func formatAttributePaths(cws []string) []string {
	for i, cw := range cws {
		aux := strings.Split(cw, ".")
		cws[i] = aux[len(aux)-1]
	}
	return cws
}

// mapHasKeys checks if any of the keys is present on the res
func mapHasKeys(res map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		if _, ok := res[key]; ok {
			return true
		}
	}
	return false
}

// normalizeValue removes the \n from the value now
func normalizeValue(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		return strings.Replace(s, "\n", "", -1)
	}
	return v
}

var iamInternpolationRe = regexp.MustCompile(`(\$\{[^}]+\})`)

// normalizeInterpolation fixes the https://github.com/hashicorp/terraform/issues/18937
// on reading
func normalizeInterpolation(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		return iamInternpolationRe.ReplaceAllString(s, `$$$1`)
	}
	return v
}

// normalizeSetList returns the normalization of a schema.Set.List
// it could be a simple list or a embedded structure.
// The sch it's used to also add required values if needed
func normalizeSetList(sch map[string]*schema.Schema, list []interface{}) interface{} {
	var ar interface{} = make([]interface{}, 0)

	for _, set := range list {
		switch val := set.(type) {
		case map[string]interface{}:
			// This case it's when a TypeSet has
			// a nested structure,
			// example: aws_security_group.ingress
			res := make(map[string]interface{})
			for k, v := range val {
				switch vv := v.(type) {
				case *schema.Set:
					nsch := make(map[string]*schema.Schema)
					if sc, ok := sch[k]; ok {
						if rs, ok := sc.Elem.(*schema.Resource); ok {
							nsch = rs.Schema
						}
					}
					ns := normalizeSetList(nsch, vv.List())
					if !isDefault(sch[k], ns) {
						res[k] = ns
					}
				case []interface{}:
					nsch := make(map[string]*schema.Schema)
					if sc, ok := sch[k]; ok {
						if rs, ok := sc.Elem.(*schema.Resource); ok {
							nsch = rs.Schema
						}
					}
					ns := normalizeSetList(nsch, vv)
					if !isDefault(sch[k], ns) {
						res[k] = ns
					}
				case map[string]interface{}:
					if !isDefault(sch[k], v) {
						res[fmt.Sprintf("=tc=%s", k)] = v
					}
				case interface{}:
					if !isDefault(sch[k], v) {
						res[k] = v
					}
				}
			}
			ar = append(ar.([]interface{}), res)
		case []interface{}:
			ns := normalizeSetList(sch, val)
			if !isDefault(nil, ns) {
				ar = append(ar.([]interface{}), ns)
			}
		case interface{}:
			// This case is normally for the
			// "Type: schema.TypeSet, Elm: schema.Schema{Type: schema.TypeString}"
			// definitions on TF,
			// example: aws_security_group.ingress.security_groups
			if !isDefault(nil, val) {
				ar = append(ar.([]interface{}), val)
			}
		}
	}

	return ar
}

var (
	// Ideally this could be generated using "enumer", it
	// would be a better idea as then we do not have
	// to maintain this list
	tfTypes = []schema.ValueType{
		schema.TypeBool,
		schema.TypeInt,
		schema.TypeFloat,
		schema.TypeString,
		schema.TypeList,
		schema.TypeMap,
		schema.TypeSet,
	}
)

// isDefault is used on normalizSet as the Sets do not use the normal
// TF strucure (access by key) and are stored as raw maps with some
// default values that we don't want on the HCL output.
// example: [], false, "", 0 ...
func isDefault(sch *schema.Schema, v interface{}) bool {
	if sch != nil {
		if sch.Required {
			return false
		}

		// This way values that are not suppose
		// to be on the config are also not added
		if !isConfig(sch) {
			return true
		}
	}

	for _, t := range tfTypes {
		if reflect.DeepEqual(t.Zero(), v) {
			// If it has a default value which is different
			// than the one v then v has to be setted.
			// Example: Default => true, v => false
			// the v = false has to be setted
			if sch.Default != nil {
				if v != sch.Default {
					return false
				}
			}
			return true
		}
	}
	return false
}

// isConfig  checks if the sch has to be
// set to a config opt or not
func isConfig(sch *schema.Schema) bool {
	if (sch.Computed && !sch.Optional && !sch.Required) || sch.Deprecated != "" {
		return false
	}
	return true
}
