package main

import log "github.com/Sirupsen/logrus"

func main() {
	config, err := NewConfig("config.hcl")
	if err != nil {
		log.Errorf("Failed to read or parse config file: %s", err)
		return
	}
	log.Info("Loaded and parsed configuration file")

	log.Info("")
	nomad, err := NewNomad(config.Nomad)
	if err != nil {
		log.Fatalf("Failed to create Nomad Client: %s", err)
		return
	}
	log.Info("Successfully created Nomad Client")

	dc, err := nomad.Agent().Datacenter()
	if err != nil {
		log.Fatalf("  Failed to get Nomad DC: %s", err)
		return
	}
	log.Infof("  -> DC: %s", dc)

	backends, err := IniitalizeBackends(config.Backends)
	if err != nil {
		log.Fatalf("%s", err)
		return
	}

	log.Info("")
	log.Infof("Found %d backends", len(backends))
	for name := range backends {
		log.Infof("  -> %s", name)
	}
	log.Info("")
	log.Infof("Found %d jobs", len(config.Jobs))
	for _, jobConfig := range config.Jobs {
		log.Infof("  -> Job: %s", jobConfig.Name)

		for _, groupConfig := range jobConfig.Groups {
			log.Infof("  --> Group: %s", groupConfig.Name)
			log.Infof("      min_count = %d", groupConfig.MinCount)
			log.Infof("      max_count = %d", groupConfig.MaxCount)

			// _ := NewRunner(job, groupConfig)

			for _, ruleConfig := range groupConfig.Rules {
				log.Infof("  ----> Rule: %s", ruleConfig.Name)

				log.Infof("      Backend 	 = %s", ruleConfig.Backend)
				log.Infof("      CheckType 	 = %s", ruleConfig.CheckType)
				log.Infof("      Comparison 	 = %s", ruleConfig.Comparison)
				log.Infof("      ComparisonValue = %f", ruleConfig.ComparisonValue)
				log.Infof("      IfTrue          = %+v", ruleConfig.IfTrue)
				log.Infof("      IfFalse         = %+v", ruleConfig.IfFalse)
			}
		}
	}
}
