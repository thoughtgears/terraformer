package terraform

import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

//go:embed templates/*
var content embed.FS

func (t *Terraform) Generate(outFile string) error {

	// Generate the main.tf template file
	mainTemplate, err := content.ReadFile("templates/main.tf.tmpl")
	if err != nil {
		return fmt.Errorf("error reading template: %w", err)
	}

	mainTmpl, err := template.New("main.tf").Parse(string(mainTemplate))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	if err := mainTmpl.Execute(os.Stdout, t); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	// Generate the module templates
	moduleTemplate, err := content.ReadFile("templates/module.tf.tmpl")
	if err != nil {
		return fmt.Errorf("error reading template: %w", err)
	}

	for _, module := range t.Modules {
		moduleTmpl, err := template.New(fmt.Sprintf("%s.tf", module.Name)).Parse(string(moduleTemplate))
		if err != nil {
			return fmt.Errorf("error parsing template: %w", err)
		}

		if err := moduleTmpl.Execute(os.Stdout, module); err != nil {
			return fmt.Errorf("error executing template: %w", err)
		}
	}

	return nil
}

// TODO: ONe single function to generate and output the template
func (t *Terraform) generateAndOutput(input, templateName, output string) error {
	data, err := content.ReadFile(fmt.Sprintf("templates/%s.tf.tmpl", templateName))
	if err != nil {
		return fmt.Errorf("error reading template: %w", err)
	}

	tmpl, err := template.New(fmt.Sprintf("%s.tf", output)).Parse(string(data))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	if err := tmpl.Execute(os.Stdout, t); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
