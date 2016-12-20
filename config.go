package main

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/hcl"
)

// NewConfig will return a Config struct
func NewConfig(file string) (*Config, error) {
	config, err := ioutil.ReadFile(file)
	if err != nil {
		log.Errorf("Failed to read file: %s", err)
		return nil, err
	}

	var out Config
	hclErr := hcl.Decode(&out, string(config))
	if hclErr != nil {
		log.Errorf("HCL Error: %s", hclErr)
		return nil, hclErr
	}

	for jobName, jobConfig := range out.Jobs {
		jobConfig.Name = jobName

		for groupName, groupConfig := range jobConfig.Groups {
			groupConfig.Name = groupName

			for ruleName, ruleConfig := range groupConfig.Rules {
				ruleConfig.Name = ruleName
			}
		}
	}

	return &out, nil
}
