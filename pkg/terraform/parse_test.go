package terraform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thoughtgears/terraformer/pkg/terraform"
)

func TestParseYaml(t *testing.T) {
	data := `
general:
  organisation: test-org
  workspace: test-workspace
  terraformVersion: 1.3.0
  region: europe-west1
  projectId: test-123456
service:
  name: test-project
  description: this is a test project
  modules:
  - name: test-module1
    organisation: test-org
    description: this is a test module
    version: v1.0.0
  - name: test-module2
    organisation: test-org
    description: this is a test module
    version: v2.0.0
`
	tf, err := terraform.Parse("yaml", []byte(data))

	assert.NoError(t, err)
	assert.Equal(t, "test-org", tf.General.Organisation)
	assert.Equal(t, "test-workspace", tf.General.Workspace)
	assert.Equal(t, "1.3.0", tf.General.TerraformVersion)
	assert.Equal(t, "europe-west1", tf.General.Region)
	assert.Equal(t, "test-123456", tf.General.ProjectID)
	assert.Equal(t, "test-project", tf.Service.Name)
	assert.Equal(t, "this is a test project", tf.Service.Description)

	assert.Equal(t, []terraform.Module{
		{
			Name:        "test-module1",
			Description: "this is a test module",
			Version:     "v1.0.0",
		}, {
			Name:        "test-module2",
			Description: "this is a test module",
			Version:     "v2.0.0",
		},
	},
		tf.Service.Modules)
}

func TestParseJson(t *testing.T) {
	data := `
{
  "general":{
    "organisation":"test-org",
    "workspace":"test-workspace",
    "terraformVersion":"1.3.0",
    "region":"europe-west1",
    "projectId":"test-123456"
  },
  "service":{
    "name":"test-project",
    "description":"this is a test project",
    "modules":[
      {
        "name":"test-module1",
		"organisation":"test-org",
        "description":"this is a test module",
        "version":"v1.0.0"
      },
      {
        "name":"test-module2",
		"organisation":"test-org",
        "description":"this is a test module",
        "version":"v2.0.0"
      }
    ]
  }
}
`
	tf, err := terraform.Parse("json", []byte(data))

	assert.NoError(t, err)
	assert.Equal(t, "test-org", tf.General.Organisation)
	assert.Equal(t, "test-workspace", tf.General.Workspace)
	assert.Equal(t, "1.3.0", tf.General.TerraformVersion)
	assert.Equal(t, "europe-west1", tf.General.Region)
	assert.Equal(t, "test-123456", tf.General.ProjectID)
	assert.Equal(t, "test-project", tf.Service.Name)
	assert.Equal(t, "this is a test project", tf.Service.Description)

	assert.Equal(t, []terraform.Module{
		{
			Name:        "test-module1",
			Description: "this is a test module",
			Version:     "v1.0.0",
		}, {
			Name:        "test-module2",
			Description: "this is a test module",
			Version:     "v2.0.0",
		},
	},
		tf.Service.Modules)
}

func TestParseErrorFileFormat(t *testing.T) {
	data := ``
	_, err := terraform.Parse("invalid", []byte(data))

	assert.Error(t, err)
	assert.Equal(t, "invalid file type invalid, filetypes supported are: json, yaml", err.Error())
}

func TestParseValidationError(t *testing.T) {
	data := `
description: this is a test project
projectId: test-123456
`
	_, err := terraform.Parse("yaml", []byte(data))

	assert.Error(t, err)
	assert.Equal(t, "field error: TerraformVersion, rule: required", err.Error())
}
