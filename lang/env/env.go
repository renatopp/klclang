package env

import "klc/lang/obj"

type Env struct {
	Parent *Env
	Store  map[string]obj.Object
}

func NewEnv(parent *Env) *Env {
	return &Env{
		Parent: parent,
		Store:  make(map[string]obj.Object),
	}
}

func (e *Env) Get(name string) (obj.Object, bool) {
	obj, ok := e.Store[name]
	if !ok && e.Parent != nil {
		obj, ok = e.Parent.Get(name)
	}

	return obj, ok
}

func (e *Env) Set(name string, val obj.Object) obj.Object {
	e.Store[name] = val
	return val
}

func (e *Env) Delete(name string) {
	delete(e.Store, name)
}

func (e *Env) Clear() {
	e.Store = make(map[string]obj.Object)
}

func (e *Env) ForEach(fn func(string, obj.Object)) {
	for k, v := range e.Store {
		fn(k, v)
	}
}
