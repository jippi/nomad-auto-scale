package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	mapstructure "github.com/mitchellh/mapstructure"
)

// IniitalizeBackends will derp
func IniitalizeBackends(backends map[string]Backend) (ConfiguredBackends, error) {
	configuredBackends := make(ConfiguredBackends, 0)

	for name, backend := range backends {
		backendType := backend["type"]

		if backendType == "" {
			return nil, fmt.Errorf("Missing backend type for '%s'", name)
		}

		switch backendType {
		case "rabbitmq":
			var config RabbitMQConfig
			err := mapstructure.Decode(backend, &config)
			if err != nil {
				return nil, fmt.Errorf("Could not decode configuration %s: %s", name, err)
			}

			connection, err := NewRabbitMQBackend(name, config)
			if err != nil {
				return nil, fmt.Errorf("Bad configuration for %s: %s", name, err)
			}

			configuredBackends[name] = connection

		default:
			log.Fatalf("unknown backend type '%s' for backend %s", backendType, name)
			return nil, fmt.Errorf("unknown backend %s", backendType)
		}
	}

	return configuredBackends, nil
}
