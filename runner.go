package main

import (
	"fmt"
)

// NewRunner will create a new task runner
func NewRunner(job *Job, group *Group) *Runner {
	runner := &Runner{}
	runner.Job = job
	runner.Group = group

	return runner
}

// LoadRules ...
func (r *Runner) LoadRules(backends ConfiguredBackends) error {
	for name, rule := range r.Group.Rules {
		if backends[rule.Backend] == nil {
			return fmt.Errorf("Unknown backend: %s (%s)", rule.Backend, name)
		}

		rule.BackendInstance = backends[rule.Backend]
	}

	return nil
}

// Validate runner configuration
func (r *Runner) Validate() error {
	return nil
}

// Start runner
func (r *Runner) Start() {
	r.Running = true
}

// Stop runner
func (r *Runner) Stop() {
	r.Running = false
}

// Work runner
func (r *Runner) Work() {
	for _, rule := range r.Group.Rules {
		_ = rule.Work()
	}
}
