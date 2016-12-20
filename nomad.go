package main

import (
	api "github.com/hashicorp/nomad/api"
)

// NewNomad will create a instance of a nomad API Client
func NewNomad(config NomadConfig) (*api.Client, error) {
	nomadDefaultConfig := api.DefaultConfig()
	nomadDefaultConfig.Address = config.Address

	client, err := api.NewClient(nomadDefaultConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
