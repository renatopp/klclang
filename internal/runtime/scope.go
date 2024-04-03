package runtime

import (
	"fmt"
	"strings"
)

type Scope struct {
	depth  int
	store  map[string]Object
	parent *Scope
}

func NewScope() *Scope {
	return &Scope{
		depth:  0,
		store:  make(map[string]Object),
		parent: nil,
	}
}

func (s *Scope) New() *Scope {
	return &Scope{
		depth:  s.depth + 1,
		store:  make(map[string]Object),
		parent: s,
	}
}

func (s *Scope) Set(name string, value Object) {
	s.store[name] = value
}

func (s *Scope) Get(name string) Object {
	if value, ok := s.store[name]; ok {
		return value
	}

	if s.parent != nil {
		return s.parent.Get(name)
	}

	return nil
}

func (s *Scope) GetInScope(name string) Object {
	if value, ok := s.store[name]; ok {
		return value
	}

	return nil
}

func (s *Scope) Keys() []string {
	keys := make([]string, 0, len(s.store))
	for k := range s.store {
		keys = append(keys, k)
	}
	return keys
}

func (s *Scope) Depth() int {
	return s.depth
}

func (s *Scope) Parent() *Scope {
	return s.parent
}

func (s *Scope) DebugString() string {
	result := fmt.Sprintf("Scope at depth %d\n", s.depth)
	for k, v := range s.store {
		result += fmt.Sprintf("- %s: %s\n", k, v.String())
	}
	return result
}

func (s *Scope) DebugStringRecursive() string {
	result := s.DebugString()
	if s.parent != nil {
		result += "- inherits from "
		result += strings.ReplaceAll(s.parent.DebugStringRecursive(), "\n", "\n  ")
	}
	return result
}
