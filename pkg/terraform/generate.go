package terraform

import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

//go:embed templates/*
var content embed.FS

// Generate will generate terraform files from the Terraform struct.
// It will return an error if the template cannot be read, parsed or executed.
// It will also return an error if the file cannot be created.
// The file will be created in the outputDir. Defaults to the current directory.
// The file name will be main.tf for the main file and <module_name>.tf for each module.
func (t *Terraform) Generate(outputDir string) error {
	if err := generateAndOutput("main", fmt.Sprintf("%s/main.tf", outputDir), t); err != nil {
		return err
	}

	for _, module := range t.Service.Modules {
		if err := generateAndOutput("module", fmt.Sprintf("%s/%s.tf", outputDir, module.Name), module); err != nil {
			return err
		}
	}

	return nil
}

func generateAndOutput(input, output string, data interface{}) error {
	fmt.Println(data)
	templateFile, err := content.ReadFile(fmt.Sprintf("templates/%s.tf.tmpl", input))
	if err != nil {
		return fmt.Errorf("error reading template: %w", err)
	}

	tmpl, err := template.New(input).Parse(string(templateFile))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
