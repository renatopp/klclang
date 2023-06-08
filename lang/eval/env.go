package eval

import (
	"fmt"
	"klc/lang/obj"
)

type Environment struct {
	store map[string]obj.Object
}

func NewEnvironment() *Environment {
	s := make(map[string]obj.Object)
	return &Environment{store: s}
}

func (e *Environment) Get(name string) (obj.Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val obj.Object) obj.Object {
	e.store[name] = val
	return val
}

func (e *Environment) Delete(name string) {
	delete(e.store, name)
}

func (e *Environment) Clear() {
	e.store = make(map[string]obj.Object)
}

func (e *Environment) ForEach(fn func(string, obj.Object)) {
	for k, v := range e.store {
		fn(k, v)
	}
}

// --

type EnvironmentStack struct {
	Stack []*Environment
}

func NewStack() *EnvironmentStack {
	return &EnvironmentStack{Stack: []*Environment{
		NewEnvironment(),
	}}
}

func (s *EnvironmentStack) Current() *Environment {
	return s.Stack[len(s.Stack)-1]
}

func (s *EnvironmentStack) Push() {
	s.Stack = append(s.Stack, NewEnvironment())
}

func (s *EnvironmentStack) Pop() {
	s.Stack = s.Stack[:len(s.Stack)-1]
}

func (s *EnvironmentStack) Get(name string) (obj.Object, bool) {
	for i := len(s.Stack) - 1; i >= 0; i-- {
		obj, ok := s.Stack[i].Get(name)
		if ok {
			return obj, ok
		}
	}

	return nil, false
}

func (s *EnvironmentStack) Set(name string, val obj.Object) obj.Object {
	s.Current().Set(name, val)
	return val
}

func (s *EnvironmentStack) Delete(name string) {
	s.Current().Delete(name)
}

func (s *EnvironmentStack) Clear() {
	s.Stack = []*Environment{
		NewEnvironment(),
	}
}

func (s *EnvironmentStack) ForEach(fn func(string, obj.Object)) {
	for i := len(s.Stack) - 1; i >= 0; i-- {
		s.Stack[i].ForEach(fn)
	}
}

func (s *EnvironmentStack) Print() {
	for i := len(s.Stack) - 1; i >= 0; i-- {
		fmt.Printf("Environment %d:\n", i)
		s.Stack[i].ForEach(func(name string, val obj.Object) {
			fmt.Printf("  %s: %s\n", name, val.AsString())
		})
	}
}
