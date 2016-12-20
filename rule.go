package main

import log "github.com/Sirupsen/logrus"
import "fmt"

// Work for derp
func (r *Rule) Work() error {
	if r.BackendInstance == nil {
		return fmt.Errorf("No BackendInstance set")
	}

	value, err := r.BackendInstance.GetValue(*r)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}

	log.Infof("Value: %f", value)
	return nil
}

// GetConfigString ...
func (r *Rule) GetConfigString(key string, defaultValue string) string {
	if r.Config[key] == nil {
		return defaultValue
	}

	value, ok := r.Config[key].(string)
	if !ok {
		return defaultValue
	}

	return value
}
