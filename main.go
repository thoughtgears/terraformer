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

	tf, err := terraform.Parse(*fileType, f)
	if err != nil {
		log.Fatal().Err(err).Msg("error parsing file")
	}

	if err := tf.Generate(*outputDir); err != nil {
		log.Fatal().Err(err).Msg("error generating terraform files")
	}

	client, err := terraform.NewClient(tf, *outputDir)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating terraform client")
	}

	if err := client.Format(); err != nil {
		log.Fatal().Err(err).Msg("error formatting terraform files")
	}
}
