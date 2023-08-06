package terraform

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
)

type Client struct {
	tf  *tfexec.Terraform
	ctx context.Context
}

func NewClient(t *Terraform, outputDir string) (*Client, error) {
	ctx := context.Background()
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion(t.General.TerraformVersion)),
	}

	execPath, err := installer.Install(ctx)
	if err != nil {
		return nil, fmt.Errorf("error installing terraform: %w", err)
	}

	tf, err := tfexec.NewTerraform(outputDir, execPath)
	if err != nil {
		return nil, fmt.Errorf("error creating terraform: %w", err)
	}

	return &Client{
		tf:  tf,
		ctx: ctx,
	}, nil
}
