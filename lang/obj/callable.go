package obj

type Callable interface {
	GetParams() []*FunctionParam
	GetScope() ScopedStore
}
