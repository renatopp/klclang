package env

type Stack struct {
	Global *Env
	Stacks []*Env
}

func NewStack() *Stack {
	return &Stack{
		Global: &Env{},
		Stacks: []*Env{},
	}
}

func (s *Stack) Current() *Env {
	if len(s.Stacks) == 0 {
		return s.Global
	}

	return s.Stacks[len(s.Stacks)-1]
}

func (s *Stack) Push() *Env {
	s.Stacks = append(s.Stacks, s.Current().New())
	return s.Current()
}

func (s *Stack) Pop() *Env {
	if len(s.Stacks) > 0 {
		s.Stacks = s.Stacks[:len(s.Stacks)-1]
	}

	return s.Current()
}
