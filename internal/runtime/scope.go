package runtime

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

// TODO: Find variable in the current scope and all parent scopes
func (s *Scope) Get(name string) Object {
	if value, ok := s.store[name]; ok {
		return value
	}

	if s.parent != nil {
		return s.parent.Get(name)
	}

	return nil
}

func (s *Scope) Depth() int {
	return s.depth
}

func (s *Scope) Parent() *Scope {
	return s.parent
}
