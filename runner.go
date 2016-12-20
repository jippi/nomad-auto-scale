package main

type Runner struct {
}

// NewRunner will create a new task runner
func NewRunner(job Job, group Group) (Runner, error) {
	return Runner{}, nil
}
