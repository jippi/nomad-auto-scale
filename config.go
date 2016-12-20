package main

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/hcl"
)

// NewConfig will return a Config struct
func NewConfig(file string) (*Config, error) {
	var out Config

	config, err := ioutil.ReadFile(file)
	if err != nil {
		log.Errorf("Failed to read file: %s", err)
		return nil, err
	}

	hclErr := hcl.Decode(&out, string(config))
	if hclErr != nil {
		log.Errorf("HCL Error: %s", hclErr)
		return nil, hclErr
	}

	return &out, nil
}
