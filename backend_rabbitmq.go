package main

import "fmt"
import rabbit "github.com/michaelklishin/rabbit-hole"
import log "github.com/Sirupsen/logrus"

// RabbitMQConfig struct
type RabbitMQConfig struct {
	Type     string
	Address  string
	Username string
	Password string
}

// RabbitMQBackend struct
type RabbitMQBackend struct {
	Name       string
	Config     RabbitMQConfig
	Connection *rabbit.Client
}

// NewRabbitMQBackend will create a new RabbitMQ Client
func NewRabbitMQBackend(name string, config RabbitMQConfig) (*RabbitMQBackend, error) {
	if config.Address == "" {
		return nil, fmt.Errorf("Missing address key")
	}

	// create the rabbitmq http client
	client, err := rabbit.NewClient(config.Address, config.Username, config.Password)
	if err != nil {
		return nil, err
	}

	// ensure we can connect to the backend and that it works
	whoami, permErr := client.Whoami()
	if permErr != nil {
		log.Errorf("Sad: %+v", permErr)
		return nil, permErr
	}

	if whoami.Name == "" {
		return nil, fmt.Errorf("Possible authentication issue, could not get information about my own account")
	}

	backend := &RabbitMQBackend{}
	backend.Name = name
	backend.Config = config
	backend.Connection = client

	return backend, nil
}
