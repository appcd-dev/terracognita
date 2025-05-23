package hcl_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/cycloidio/mxwriter"
	"github.com/cycloidio/terracognita/errcode"
	"github.com/cycloidio/terracognita/hcl"
	"github.com/cycloidio/terracognita/interpolator"
	"github.com/cycloidio/terracognita/mock"
	"github.com/cycloidio/terracognita/writer"
	"github.com/golang/mock/gomock"
	aws "github.com/hashicorp/terraform-provider-aws/provider"
	azurerm "github.com/hashicorp/terraform-provider-azurerm/provider"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHCLWriter(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
		)
		p.EXPECT().String().Return("aws").Times(2)
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")
		provider, err := aws.New(t.Context())
		require.NoError(t, err)
		p.EXPECT().TFProvider().Return(provider)
		p.EXPECT().Configuration().Return(map[string]interface{}{
			"region": "eu-west-1",
		})

		hw := hcl.NewWriter(nil, p, &writer.Options{HCLProviderBlock: true})
		assert.Equal(t, map[string]map[string]interface{}{
			"hcl": map[string]interface{}{
				"provider": map[string]interface{}{"aws": map[string]interface{}{}},
				"resource": map[string]map[string]interface{}{},
				"terraform": map[string]interface{}{
					"required_providers": map[string]interface{}{
						"=tc=aws": map[string]interface{}{
							"source":  "hashicorp/aws",
							"version": "=4.9.0",
						},
					},
					"required_version": ">= 1.0",
				},
			},
		}, hw.Config)
	})
	t.Run("SuccessWithSplitConfig", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
		)
		p.EXPECT().String().Return("aws").Times(2)
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")
		provider, err := aws.New(t.Context())
		require.NoError(t, err)
		p.EXPECT().TFProvider().Return(provider)
		p.EXPECT().Configuration().Return(map[string]interface{}{
			"region": "eu-west-1",
		})

		hw := hcl.NewWriter(nil, p, &writer.Options{
			HCLProviderBlock:     true,
			TerraformCategoryKey: "config",
		})

		assert.True(t, reflect.DeepEqual(map[string]map[string]interface{}{
			"hcl": map[string]interface{}{
				"resource": map[string]map[string]interface{}{},
			},
			"config": map[string]interface{}{
				"provider": map[string]interface{}{"aws": map[string]interface{}{}},
				"terraform": map[string]interface{}{
					"required_providers": map[string]interface{}{
						"=tc=aws": map[string]interface{}{
							"source":  "hashicorp/aws",
							"version": "=4.9.0",
						},
					},
					"required_version": ">= 1.0",
				},
			},
		}, hw.Config))
	})
	t.Run("SuccessWithoutProviderBLock", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(nil, p, &writer.Options{})
		assert.Equal(t, map[string]map[string]interface{}{
			"hcl": map[string]interface{}{
				"resource": map[string]map[string]interface{}{},
				"terraform": map[string]interface{}{
					"required_providers": map[string]interface{}{
						"=tc=aws": map[string]interface{}{
							"source":  "hashicorp/aws",
							"version": "=4.9.0",
						},
					},
					"required_version": ">= 1.0",
				},
			},
		}, hw.Config)
	})
	t.Run("SuccessWithModule", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(nil, p, &writer.Options{Module: "my-module"})
		assert.Equal(t, map[string]map[string]interface{}{
			"tc_module": map[string]interface{}{
				"module": map[string]interface{}{
					"my-module": map[string]interface{}{
						"source": "./module-my-module",
					},
				},
				"terraform": map[string]interface{}{
					"required_providers": map[string]interface{}{
						"=tc=aws": map[string]interface{}{
							"source":  "hashicorp/aws",
							"version": "=4.9.0",
						},
					},
					"required_version": ">= 1.0",
				},
			},
		}, hw.Config)
	})
	t.Run("SuccessWithModuleAndConfig", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(nil, p, &writer.Options{
			Module:               "my-module",
			TerraformCategoryKey: "config",
		})
		assert.Equal(t, map[string]map[string]interface{}{
			"tc_module": map[string]interface{}{
				"module": map[string]interface{}{
					"my-module": map[string]interface{}{
						"source": "./module-my-module",
					},
				},
			},
			"config": map[string]interface{}{
				"terraform": map[string]interface{}{
					"required_providers": map[string]interface{}{
						"=tc=aws": map[string]interface{}{
							"source":  "hashicorp/aws",
							"version": "=4.9.0",
						},
					},
					"required_version": ">= 1.0",
				},
			},
		}, hw.Config)
	})
}

func TestHCLWriter_Write(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			mw    = mxwriter.NewMux()
			value = map[string]interface{}{
				"key": "value",
			}
			key  = "type.name"
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
		)
		defer ctrl.Finish()

		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write(key, value)
		require.NoError(t, err)

		assert.Equal(t, map[string]map[string]interface{}{
			"hcl": map[string]interface{}{
				"resource": map[string]map[string]interface{}{
					"type": map[string]interface{}{
						"name": map[string]interface{}{
							"key": "value",
						},
					},
				},
				"terraform": map[string]interface{}{
					"required_providers": map[string]interface{}{
						"=tc=aws": map[string]interface{}{
							"source":  "hashicorp/aws",
							"version": "=4.9.0",
						},
					},
					"required_version": ">= 1.0",
				},
			},
		}, hw.Config)
		t.Run("Has", func(t *testing.T) {
			ok, err := hw.Has(key)
			require.NoError(t, err)
			assert.True(t, ok)

			ok, err = hw.Has("type.new")
			require.NoError(t, err)
			assert.False(t, ok)

			t.Run("Empty", func(t *testing.T) {
				var (
					ctrl = gomock.NewController(t)
					p    = mock.NewProvider(ctrl)
				)

				p.EXPECT().String().Return("aws")
				p.EXPECT().Source().Return("hashicorp/aws")
				p.EXPECT().Version().Return("4.9.0")

				mw = mxwriter.NewMux()
				hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true, Module: "s"})

				ok, err = hw.Has("type.new")
				require.NoError(t, err)
				assert.False(t, ok)
			})
		})
	})
	t.Run("ErrRequiredKey", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
			mw   = mxwriter.NewMux()
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write("", nil)
		assert.Equal(t, errcode.ErrWriterRequiredKey, errors.Cause(err))
	})
	t.Run("ErrRequiredValue", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
			mw   = mxwriter.NewMux()
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write("type.name", nil)
		assert.Equal(t, errcode.ErrWriterRequiredValue, errors.Cause(err))
	})
	t.Run("ErrInvalidKey", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
			mw   = mxwriter.NewMux()
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write("type.name.name", "")
		assert.Equal(t, errcode.ErrWriterInvalidKey, errors.Cause(err))

		err = hw.Write("type", "")
		assert.Equal(t, errcode.ErrWriterInvalidKey, errors.Cause(err))

		err = hw.Write("type.", "")
		assert.Equal(t, errcode.ErrWriterInvalidKey, errors.Cause(err))
	})
	t.Run("ErrAlreadyExistsKey", func(t *testing.T) {
		var (
			ctrl = gomock.NewController(t)
			p    = mock.NewProvider(ctrl)
			mw   = mxwriter.NewMux()
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write("type.name", map[string]interface{}{})
		require.NoError(t, err)

		err = hw.Write("type.name", map[string]interface{}{})
		assert.Equal(t, errcode.ErrWriterAlreadyExistsKey, errors.Cause(err))
	})
}

func TestHCLWriter_Sync(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mx    = mxwriter.NewMux()
			value = map[string]interface{}{
				"key":         "value",
				"tc_category": "some-category",
			}
			ehcl = `
provider "aws" { }

terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}

resource "type" "name" {
  key = "value"
}

`
		)

		p.EXPECT().String().Return("aws").Times(2)
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")
		provider, err := aws.New(t.Context())
		require.NoError(t, err)
		p.EXPECT().TFProvider().Return(provider)
		p.EXPECT().Configuration().Return(map[string]interface{}{
			"region": "eu-west-1",
		})

		hw := hcl.NewWriter(mx, p, &writer.Options{HCLProviderBlock: true, Interpolate: true})

		err = hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mx)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
	t.Run("SuccessWithoutProviderBlock", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mx    = mxwriter.NewMux()
			value = map[string]interface{}{
				"key":         "value",
				"tc_category": "some-category",
			}
			ehcl = `
terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}

resource "type" "name" {
  key = "value"
}

`
		)

		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mx, p, &writer.Options{Interpolate: true})

		err := hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mx)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
	t.Run("Module", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mx    = mxwriter.NewMux()
			value = map[string]interface{}{
				"key": "value",
			}
			value2 = map[string]interface{}{
				"key":  "value",
				"key2": "value",
				"key3": []interface{}{},
				"key4": map[string]interface{}{
					"nested_key4": "value4",
				},
				"key5": []interface{}{
					map[string]interface{}{
						"nested_key5": "value5.0",
					},
					map[string]interface{}{
						"nested_key5": "value5.1",
					},
				},
			}
			ehcl = `
resource "type" "name" {
  key = var.type_name_key
}

resource "type" "name2" {
  key = var.type_name2_key
	key2 = var.type_name2_key2
	key3 = var.type_name2_key3
	key4 {
		nested_key4 = var.type_name2_key4_nested_key4
	}

	key5 {
		nested_key5 = var.type_name2_key5_0_nested_key5
	}

	key5 {
		nested_key5 = var.type_name2_key5_1_nested_key5
	}
}

module "test" {
  source = "./module-test"
	type_name2_key = "value"
	type_name2_key2 = "value"
	type_name2_key3 = []
	type_name2_key4_nested_key4 = "value4"
	type_name2_key5_0_nested_key5 = "value5.0"
	type_name2_key5_1_nested_key5 = "value5.1"
	type_name_key = "value"
}

provider "aws" { }

terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}

variable "type_name2_key" {
	default = "value"
}

variable "type_name2_key2" {
	default = "value"
}

variable "type_name2_key3" {
	default = []
}

variable "type_name2_key4_nested_key4" {
	default = "value4"
}

variable "type_name2_key5_0_nested_key5" {
	default = "value5.0"
}

variable "type_name2_key5_1_nested_key5" {
	default = "value5.1"
}

variable "type_name_key" {
	default = "value"
}
`
		)
		p.EXPECT().String().Return("aws").Times(2)
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")
		provider, err := aws.New(t.Context())
		require.NoError(t, err)
		p.EXPECT().TFProvider().Return(provider)
		p.EXPECT().Configuration().Return(map[string]interface{}{
			"region": "eu-west-1",
		})

		hw := hcl.NewWriter(mx, p, &writer.Options{Interpolate: true, HCLProviderBlock: true, Module: "test"})

		err = hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Write("type.name2", value2)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mx)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
	t.Run("ModuleVariables", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mx    = mxwriter.NewMux()
			value = map[string]interface{}{
				"key": "value",
				"key3": map[string]interface{}{
					"nested_key3": "nvalue3",
				},
				"key4": []interface{}{
					map[string]interface{}{
						"nested_key4": "nvalue4.0",
					},
					map[string]interface{}{
						"nested_key4": "nvalue4.1",
					},
				},
				"=tc=tags": map[string]interface{}{
					"tagk": "tagv",
				},
			}
			value2 = map[string]interface{}{
				"key":  "value",
				"key2": "value",
			}
			ehcl = `
resource "type" "name" {
	tags = var.type_name_tags
  key = var.type_name_key
	key3 {
		nested_key3 = var.type_name_key3_nested_key3
	}
	key4 {
		nested_key4 = var.type_name_key4_0_nested_key4
	}
	key4 {
		nested_key4 = var.type_name_key4_1_nested_key4
	}
}

resource "type" "name2" {
  key = var.type_name2_key
	key2 = "value"
}

module "test" {
	type_name_tags = {
		tagk = "tagv"
	}
  source = "./module-test"
	type_name2_key = "value"
	type_name_key = "value"
	type_name_key3_nested_key3 = "nvalue3"
	type_name_key4_0_nested_key4 = "nvalue4.0"
	type_name_key4_1_nested_key4 = "nvalue4.1"
}

provider "aws" { }

terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}

variable "type_name2_key" {
	default = "value"
}

variable "type_name_key" {
	default = "value"
}

variable "type_name_key3_nested_key3" {
	default = "nvalue3"
}

variable "type_name_key4_0_nested_key4" {
	default = "nvalue4.0"
}

variable "type_name_key4_1_nested_key4" {
	default = "nvalue4.1"
}

variable "type_name_tags" {
	default = {
		tagk = "tagv"
	}
}
`
		)
		p.EXPECT().String().Return("aws").Times(2)
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")
		provider, err := aws.New(t.Context())
		require.NoError(t, err)
		p.EXPECT().TFProvider().Return(provider)
		p.EXPECT().Configuration().Return(map[string]interface{}{
			"region": "eu-west-1",
		})

		hw := hcl.NewWriter(mx, p, &writer.Options{Interpolate: true, HCLProviderBlock: true, Module: "test", ModuleVariables: map[string]struct{}{
			"type.key":              struct{}{},
			"type.tags":             struct{}{},
			"type.key3.nested_key3": struct{}{},
			"type.key4.nested_key4": struct{}{},
		}})

		err = hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Write("type.name2", value2)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mx)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
	t.Run("ModuleWithProviderDefaults", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mx    = mxwriter.NewMux()
			value = map[string]interface{}{
				"key": "value",
			}
			ehcl = `
resource "type" "name" {
  key = var.type_name_key
}

module "test" {
  source = "./module-test"
	type_name_key = "value"
}

provider "azurerm" {
	environment   = var.environment
	features      = var.features
	metadata_host = var.metadata_host
}

terraform {
	required_providers {
		azurerm = {
			source = "hashicorp/azurerm"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}

variable "environment" {
  default = "public"
}

variable "features" {
}

variable "metadata_host" {
	default = "host"
}

variable "type_name_key" {
	default = "value"
}
`
		)
		p.EXPECT().String().Return("azurerm").Times(5)
		p.EXPECT().Source().Return("hashicorp/azurerm")
		p.EXPECT().Version().Return("4.9.0")
		p.EXPECT().TFProvider().Return(azurerm.AzureProvider())
		p.EXPECT().Configuration().Return(map[string]interface{}{
			"metadata_host": "host",
		})

		hw := hcl.NewWriter(mx, p, &writer.Options{Interpolate: true, HCLProviderBlock: true, Module: "test"})

		err := hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mx)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
	t.Run("Slice", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mw    = mxwriter.NewMux()
			value = map[string]interface{}{
				"ingress": []map[string]interface{}{
					{
						"cidr_blocks": []string{"0.0.0.0/0"},
						"from_port":   80,
					},
					{
						"cidr_blocks": []string{"0.0.0.0/1"},
						"from_port":   81,
					},
				},
			}
			ehcl = `
resource "type" "name" {
  ingress {
		cidr_blocks = ["0.0.0.0/0"]
		from_port = 80
	}

  ingress {
		cidr_blocks = ["0.0.0.0/1"]
		from_port = 81
	}
}

terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}
`
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
	t.Run("EmptySlice", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mw    = mxwriter.NewMux()
			value = map[string]interface{}{
				"ingress": []map[string]interface{}{
					{
						"cidr_blocks": []string{},
						"from_port":   80,
					},
				},
			}
			ehcl = `
resource "type" "name" {
  ingress {
		cidr_blocks = []
		from_port = 80
	}
}

terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}
`
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
	t.Run("NestedMap", func(t *testing.T) {
		var (
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			mw    = mxwriter.NewMux()
			value = map[string]interface{}{
				"ingress": []map[string]interface{}{
					{
						"cidr_blocks": []string{},
						"from_port": map[string]interface{}{
							"in":  "vin",
							"out": "vout",
						},
					},
				},
			}
			ehcl = `
resource "type" "name" {
  ingress {
		cidr_blocks = []
		from_port {
			in = "vin"
			out = "vout"
		}
	}
}

terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "=4.9.0"
		}
	}
	required_version = ">= 1.0"
}
`
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})

		err := hw.Write("type.name", value)
		require.NoError(t, err)

		err = hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.Equal(t, strings.Join(strings.Fields(ehcl), " "), strings.Join(strings.Fields(string(b)), " "))
	})
}

func TestHCLWriter_Interpolate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var (
			mw    = mxwriter.NewMux()
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			value = map[string]interface{}{
				"network": "to-be-interpolated",
			}
			network = map[string]interface{}{
				"id": "interpolated",
			}
			i = interpolator.New("aws")
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})
		i.AddResourceAttributes("aType.aName", map[string]string{
			"id": "to-be-interpolated",
		})
		hw.Write("type.name", value)
		hw.Write("aType.aName", network)

		hw.Interpolate(i)

		err := hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.Contains(t, string(b), "network = aType.aName.id")
	})
	t.Run("SuccessWithModule", func(t *testing.T) {
		var (
			mw    = mxwriter.NewMux()
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			value = map[string]interface{}{
				"network": "to-be-interpolated",
			}
			network = map[string]interface{}{
				"id": "interpolated",
			}
			i = interpolator.New("aws")
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Module: "test", Interpolate: true})
		i.AddResourceAttributes("aType.aName", map[string]string{
			"id": "to-be-interpolated",
		})
		hw.Write("type.name", value)
		hw.Write("aType.aName", network)

		hw.Interpolate(i)

		err := hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.NotContains(t, string(b), "network = aType.aName.id")
	})
	t.Run("SuccessWithModuleVariablesNotSelected", func(t *testing.T) {
		var (
			mw    = mxwriter.NewMux()
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			value = map[string]interface{}{
				"network": "to-be-interpolated",
			}
			network = map[string]interface{}{
				"id": "interpolated",
			}
			i = interpolator.New("aws")
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Module: "test", ModuleVariables: map[string]struct{}{"type.name": struct{}{}}, Interpolate: true})
		i.AddResourceAttributes("aType.aName", map[string]string{
			"id": "to-be-interpolated",
		})
		hw.Write("type.name", value)
		hw.Write("aType.aName", network)

		hw.Interpolate(i)

		err := hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.Contains(t, string(b), "network = aType.aName.id")
	})
	t.Run("SuccessWithModuleVariablesSelected", func(t *testing.T) {
		var (
			mw    = mxwriter.NewMux()
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			value = map[string]interface{}{
				"network": "to-be-interpolated",
			}
			network = map[string]interface{}{
				"id": "interpolated",
			}
			i = interpolator.New("aws")
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Module: "test", ModuleVariables: map[string]struct{}{"type.network": struct{}{}}, Interpolate: true})
		i.AddResourceAttributes("aType.aName", map[string]string{
			"id": "to-be-interpolated",
		})
		hw.Write("type.name", value)
		hw.Write("aType.aName", network)

		hw.Interpolate(i)

		err := hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.NotContains(t, string(b), "network = aType.aName.id")
	})
	t.Run("SuccessAvoidInterpolaception", func(t *testing.T) {
		var (
			mw    = mxwriter.NewMux()
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			value = map[string]interface{}{
				"network": "to-be-interpolated",
			}
			network = map[string]interface{}{
				"id":   "interpolated",
				"name": "to-be-interpolated",
			}
			i = interpolator.New("aws")
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})
		i.AddResourceAttributes("aType.aName", map[string]string{
			"id": "to-be-interpolated",
		})
		hw.Write("type.name", value)
		hw.Write("aType.aName", network)

		hw.Interpolate(i)

		err := hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.Contains(t, string(b), "to-be-interpolated")
	})
	t.Run("SuccessMutualInterpolation", func(t *testing.T) {
		var (
			mw       = mxwriter.NewMux()
			ctrl     = gomock.NewController(t)
			p        = mock.NewProvider(ctrl)
			instance = map[string]interface{}{
				"subnet_id": "1234",
			}
			subnet = map[string]interface{}{
				"id":                "subnet-1",
				"availability_zone": "a-zone",
			}
			i = interpolator.New("aws")
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: true})
		i.AddResourceAttributes("aws_instance.instance", map[string]string{
			"availability_zone": "a-zone",
		})
		i.AddResourceAttributes("aws_subnet.subnet", map[string]string{
			"id": "1234",
		})
		hw.Write("aws_subnet.subnet", subnet)
		hw.Write("aws_instance.instance", instance)

		hw.Interpolate(i)

		err := hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		// the only way to assert that there is one interpolation is to
		// check if we have exactly one value starting by `aws_`
		assert.Equal(t, 1, strings.Count(string(b), "= aws_"))
	})
	t.Run("SuccessNoInterpolation", func(t *testing.T) {
		var (
			mw    = mxwriter.NewMux()
			ctrl  = gomock.NewController(t)
			p     = mock.NewProvider(ctrl)
			value = map[string]interface{}{
				"network": "should-not-be-interpolated",
			}
			network = map[string]interface{}{
				"id":   "interpolated",
				"name": "to-be-interpolated",
			}
			i = interpolator.New("aws")
		)
		p.EXPECT().String().Return("aws")
		p.EXPECT().Source().Return("hashicorp/aws")
		p.EXPECT().Version().Return("4.9.0")

		hw := hcl.NewWriter(mw, p, &writer.Options{Interpolate: false})
		i.AddResourceAttributes("aType.aName", map[string]string{
			"id": "should-not-be-interpolated",
		})
		hw.Write("type.name", value)
		hw.Write("aType.aName", network)

		hw.Interpolate(i)

		err := hw.Sync()
		require.NoError(t, err)

		b, err := io.ReadAll(mw)
		require.NoError(t, err)

		assert.Contains(t, string(b), "network = \"should-not-be-interpolated\"")
	})
}
