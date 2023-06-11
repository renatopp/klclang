package env

import "klc/lang/obj"

type Stack struct {
	Global *Env
	Stacks []*Env
}

func NewStack() *Stack {
	return &Stack{
		Global: NewEnv(nil),
		Stacks: []*Env{},
	}
}

func (s *Stack) Current() *Env {
	if len(s.Stacks) == 0 {
		return s.Global
	}

	return s.Stacks[len(s.Stacks)-1]
}

func (s *Stack) Push(scope *Env) *Env {
	if scope == nil {
		scope = s.Current()
	}

	s.Stacks = append(s.Stacks, NewEnv(scope))
	return s.Current()
}

func (s *Stack) Pop() *Env {
	if len(s.Stacks) > 0 {
		s.Stacks = s.Stacks[:len(s.Stacks)-1]
	}

	return s.Current()
}

func (e *Stack) Get(name string) (obj.Object, bool) {
	return e.Current().Get(name)
}

func (e *Stack) Set(name string, val obj.Object) obj.Object {
	return e.Current().Set(name, val)
}

func (e *Stack) Delete(name string) {
	e.Current().Delete(name)
}

func (e *Stack) ForEach(fn func(string, obj.Object)) {
	e.Global.ForEach(fn)
}
