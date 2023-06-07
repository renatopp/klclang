package eval

import "klc/lang/obj"

func NewEnvironment() *Environment {
	s := make(map[string]obj.Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]obj.Object
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
