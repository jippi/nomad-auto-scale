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

	runners := make(Runners, 0)

	log.Info("")
	log.Infof("Found %d jobs", len(config.Jobs))
	for _, job := range config.Jobs {
		log.Infof("  -> Job: %s", job.Name)

		for _, group := range job.Groups {
			log.Infof("  --> Group: %s", group.Name)
			log.Infof("      min_count = %d", group.MinCount)
			log.Infof("      max_count = %d", group.MaxCount)

			runner := NewRunner(job, group)
			if err := runner.LoadRules(backends); err != nil {
				log.Errorf("%s", err)
				return
			}

			if err := runner.Validate(); err != nil {
				log.Errorf("%s", err)
				return
			}

			runners = append(runners, runner)

			// for _, rule := range group.Rules {
			// 	log.Infof("  ----> Rule: %s", rule.Name)

			// 	log.Infof("      Backend 	 = %s", rule.Backend)
			// 	log.Infof("      CheckType 	 = %s", rule.CheckType)
			// 	log.Infof("      Comparison 	 = %s", rule.Comparison)
			// 	log.Infof("      ComparisonValue = %f", rule.ComparisonValue)
			// 	log.Infof("      IfTrue          = %+v", rule.IfTrue)
			// 	log.Infof("      IfFalse         = %+v", rule.IfFalse)
			// }
		}
	}

	log.Info("")
	log.Infof("Numebr of runners: %d", len(runners))

	for _, runner := range runners {
		runner.Work()
	}
}
