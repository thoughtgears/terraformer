package terraform

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

// Terraform is the struct that will be used to parse the terraform file.
// Should contain all the information needed to generate terraform code.
type Terraform struct {
	General General `yaml:"general" json:"general"`
	Service Service `yaml:"service" json:"service"`
}

// General contains general information about the terraform project.
// This information is used to generate the main.tf file.
// Organisation and Workspace are optional. If they are not provided, it will not connect to a remote state.
// TerraformVersion, Region and ProjectID are required.
type General struct {
	Organisation     string `yaml:"organisation,omitempty" json:"organisation,omitempty"`
	Workspace        string `yaml:"workspace,omitempty" json:"workspace,omitempty"`
	TerraformVersion string `yaml:"terraformVersion" json:"terraformVersion" validate:"required"`
	Region           string `yaml:"region" json:"region" validate:"required"`
	ProjectID        string `yaml:"projectId" json:"projectId" validate:"required"`
}

// Service contains information about the service.
// This information is used to generate the information across the files.
// Name is required.
// Description and Modules is optional.
type Service struct {
	Name        string   `yaml:"name" json:"name" validate:"required"`
	Description string   `yaml:"description,omitempty" json:"description,omitempty"`
	Modules     []Module `yaml:"modules,omitempty" json:"modules,omitempty"`
}

// Module contains information about the module.
// This information is used to generate the information across the module files.
// Name and Version are required.
// Description is optional.
type Module struct {
	Name            string `yaml:"name" json:"name" validate:"required"`
	GitOrganisation string `yaml:"organisation" json:"organisation" validate:"required"`
	Description     string `yaml:"description,omitempty" json:"description,omitempty"`
	Version         string `yaml:"version" json:"version" validate:"required"`
}

// Parse will parse data into a Terraform struct.
// It will return an error if the file type is not supported.
// Currently only json and yaml is supported.
// It will also return an error if the data is not valid.
// The data is not valid if it does not match the Terraform struct validation rules.
func Parse(fileType string, data []byte) (*Terraform, error) {
	var t Terraform
	if fileType == "yaml" {
		if err := yaml.Unmarshal(data, &t); err != nil {
			return nil, fmt.Errorf("error unmarshalling yaml: %w", err)
		}
	} else if fileType == "json" {
		if err := json.Unmarshal(data, &t); err != nil {
			return nil, fmt.Errorf("error unmarshalling json: %w", err)
		}
	} else {
		return nil, fmt.Errorf("invalid file type %s, filetypes supported are: json, yaml", fileType)
	}

	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				return nil, fmt.Errorf("field error: %s, rule: %s", e.Field(), e.ActualTag())
			}
		} else {
			return nil, fmt.Errorf("error validating terraform struct: %w", err)
		}
	}

	return &t, nil
}
