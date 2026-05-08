package internal

import "github.com/renatopp/klclang/internal/core"

type Runner struct {
	Cache map[string]*core.Script
	// runtime here
}

func NewRunner() *Runner {
	return &Runner{
		Cache: map[string]*core.Script{},
	}
}

func (r *Runner) Run(entry *core.Script) (any, error) {
	return nil, nil
}

func (r *Runner) Eval(entry *core.Script) (any, error) {
	return nil, nil
}
