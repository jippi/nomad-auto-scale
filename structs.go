package main

// RootConfig struct
// This is the main configuration, it contain Jobs and other auxiallary configuration
type RootConfig struct {
	Jobs     map[string]*Job    `hcl:"job"`
	Nomad    NomadConfig        `hcl:"nomad"`
	Backends map[string]Backend `hcl:"backend"`
}

// Backend struct
type Backend map[string]string

// ConfiguredBackends struct
type ConfiguredBackends map[string]Backender

// NomadConfig struct
type NomadConfig struct {
	Address string `hcl:"address"`
}

// Job Struct
type Job struct {
	Name   string
	Groups map[string]*Group `hcl:"group"`
}

// Group struct
type Group struct {
	Name     string
	MinCount int              `hcl:"min_count"`
	MaxCount int              `hcl:"max_count"`
	Rules    map[string]*Rule `hcl:"rule"`
}

// Rule struct
type Rule struct {
	Name            string
	Backend         string `hcl:"backend"`
	BackendInstance Backender
	Config          map[string]interface{}
	Comparison      string                 `hcl:"comparison"`
	ComparisonValue float64                `hcl:"comparison_value,float"`
	Action          map[string]string      `hcl:"action"`
	IfTrue          map[string]interface{} `hcl:"if_true"`
	IfFalse         map[string]interface{} `hcl:"if_false"`
}

// Runner struct
type Runner struct {
	Job     *Job
	Group   *Group
	Running bool
	stopCh  chan interface{}
}

// Runners list
type Runners []*Runner
