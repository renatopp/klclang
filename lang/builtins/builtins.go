package builtins

import "klc/lang/obj"

var TRUE = &obj.Number{Value: 1}
var FALSE = &obj.Number{Value: 0}
var EMPTY_STRING = &obj.String{Value: ""}

func NewFunction(fn obj.BuiltinFunctionFn, params ...*obj.FunctionParam) *obj.BuiltinFunction {
	return &obj.BuiltinFunction{
		Fn:     fn,
		Params: params,
	}
}

func NewParam(name string, def obj.Object, spread bool) *obj.FunctionParam {
	return &obj.FunctionParam{
		Name:    name,
		Default: def,
		Spread:  spread,
	}
}

func NewString(value string) *obj.String {
	return &obj.String{Value: value}
}

func NewNumber(value float64) *obj.Number {
	return &obj.Number{Value: value}
}

func NewNumberList(values ...float64) *obj.List {
	v := make([]obj.Object, len(values))

	for i, x := range values {
		v[i] = NewNumber(x)
	}

	return &obj.List{Values: v}
}

func WithDoc(obj obj.Object, doc string) obj.Object {
	obj.SetDoc(doc)
	return obj
}
