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

// GetValue derp
func (b *RabbitMQBackend) GetValue(rule Rule) (float64, error) {
	checkType := rule.GetConfigString("check_type", "")
	if checkType == "" {
		return 0.0, fmt.Errorf("Missing check_type insie config{} stanza")
	}

	queueName := rule.GetConfigString("queue_name", "")
	if queueName == "" {
		return 0.0, fmt.Errorf("Missing queue_name inside config{} stanza")
	}

	vhost := rule.GetConfigString("vhost", "/")

	switch checkType {
	case "queue_length":
		return b.GetQueueLength(vhost, queueName)
	case "queue_utilization":
		return b.GetQueueUtilization(vhost, queueName)
	default:
		return 0.0, fmt.Errorf("Unknown check_type: %s", checkType)
	}
}

// GetQueueLength ...
func (b *RabbitMQBackend) GetQueueLength(vhost string, queue string) (float64, error) {
	q, err := b.Connection.GetQueue(vhost, queue)
	if err != nil {
		return 0.0, err
	}

	return float64(q.Messages), nil
}

// GetQueueUtilization ...
func (b *RabbitMQBackend) GetQueueUtilization(vhost string, queue string) (float64, error) {
	q, err := b.Connection.GetQueue(vhost, queue)
	if err != nil {
		return 0.0, err
	}

	return q.ConsumerUtilisation, nil
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
