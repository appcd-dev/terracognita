package tag

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	neptunetypes "github.com/aws/aws-sdk-go-v2/service/neptune/types"
	rdstypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/chr4/pwgen"
	"github.com/cycloidio/terracognita/errcode"
	"github.com/cycloidio/terracognita/util"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// nameRegexp is the new regexp used to validate the names
// of the resources on TF (defined on configs/configschema/internal_validate.go)
var nameRegexp = regexp.MustCompile(`^[a-z0-9_]+$`)

// Tag it's an easy representation of
// a ec2.Filter for tags
type Tag struct {
	Name  string
	Value string
}

// New initializes a tag with the format NAME:VALUE that we use
func New(t string) (tag Tag, err error) {
	for _, d := range []string{"=", ":"} {
		tag, err = new(t, d)
		if err == nil {
			return tag, nil
		}
	}
	return Tag{}, err
}

func new(t string, delimitter string) (Tag, error) {
	values := strings.Split(t, delimitter)
	if len(values) != 2 {
		return Tag{}, errcode.ErrTagInvalidForamt
	}
	return Tag{Name: values[0], Value: values[1]}, nil
}

// ToEC2Filter transforms the Tag to a ec2.Filter
// to use on AWS filters
func (t Tag) ToEC2Filter() ec2types.Filter {
	return ec2types.Filter{
		Name:   aws.String(fmt.Sprintf("tag:%s", t.Name)),
		Values: []string{t.Value},
	}
}

// ToRDSFilter transforms the Tag to a rds.Filter
// to use on AWS filters
func (t Tag) ToRDSFilter() *rdstypes.Filter {
	return &rdstypes.Filter{
		Name:   aws.String(fmt.Sprintf("tag:%s", t.Name)),
		Values: []string{t.Value},
	}
}

// ToNeptuneFilter transforms the Tag to a Neptune.Filter
// to use on AWS filters
func (t Tag) ToNeptuneFilter() *neptunetypes.Filter {
	return &neptunetypes.Filter{
		Name:   aws.String(fmt.Sprintf("tag:%s", t.Name)),
		Values: []string{t.Value},
	}
}

// GetNameFromTag returns the 'tags.Name' from the src or the fallback
// if it's not defined.
// Also validates that the 'tags.Name' and fallback are valid, if not it
// generates a random one
func GetNameFromTag(key string, srd *schema.ResourceData, fallback string) string {
	fallback = strings.ToLower(fallback)

	var n string
	if name, ok := srd.GetOk(fmt.Sprintf("%s.Name", key)); ok {
		n = strings.ToLower(name.(string))
	}

	forcedN := util.NormalizeName(n)
	forcedFallback := util.NormalizeName(fallback)

	if isValidResourceName(n) && hclsyntax.ValidIdentifier(n) {
		return n
	} else if isValidResourceName(forcedN) && hclsyntax.ValidIdentifier(forcedN) && forcedN != "___" {
		return forcedN
	} else if isValidResourceName(fallback) && hclsyntax.ValidIdentifier(fallback) {
		return fallback
	} else if isValidResourceName(forcedFallback) && hclsyntax.ValidIdentifier(forcedFallback) && forcedFallback != "___" {
		return forcedFallback
	} else {
		return pwgen.Alpha(5)
	}
}

// isValidResourceName checks with the TF regex
// for names to validate if it's valid
func isValidResourceName(name string) bool {
	return nameRegexp.MatchString(name)
}

// GetOtherTags used to check other possible tag attributes on resources
func GetOtherTags(provider string, srd *schema.ResourceData, filterTag Tag) (string, bool) {
	// keep the same logic as r.data.GetOk
	otherTagsMap := make(map[string]string)

	// Special Tag attribute
	if provider == "aws" {
		// Some resource like aws_autoscaling_group do not have tags map but tag schema.Set
		// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_group#tag
		// Structure is not the same as "basic tags" and need specific code to validates
		// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/autoscaling_group#tag-and-tags
		v, _ := srd.GetOk("tag")
		// Get Set to List tag
		tag, okset := v.(*schema.Set)
		if okset {
			// convert list to map
			for _, i := range tag.List() {
				// Cast to get tag key and value
				m, okmap := i.(map[string]interface{})
				if okmap {
					otherTagsMap[m["key"].(string)] = m["value"].(string)
				}
			}
		}
	}

	// return the tag value if key (tag name) found
	if val, ok := otherTagsMap[filterTag.Name]; ok {
		return val, true
	}

	return "", false
}
