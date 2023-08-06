package main

import (
	"flag"
	"os"

	"github.com/thoughtgears/terraformer/pkg/terraform"

	"github.com/rs/zerolog/log"
)

func main() {
	fileName := flag.String("file", "input.yaml", "Name of file to terraform - default: input.yaml")
	fileType := flag.String("type", "yaml", "Type of file to terraform - default: yaml")
	outputDir := flag.String("output", "./", "Output directory for terraform files - default: current directory")
	flag.Parse()

	f, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatal().Err(err).Msg("error reading file")
	}

	t, err := terraform.Parse(*fileType, f)
	if err != nil {
		log.Fatal().Err(err).Msg("error parsing file")
	}

	if err := t.Generate(*outputDir); err != nil {
		log.Fatal().Err(err).Msg("error generating terraform files")
	}
}
