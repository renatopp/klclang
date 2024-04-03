package runtime

import (
	"math"

	"github.com/renatopp/klclang/internal/ast"
	"github.com/renatopp/langtools/asts"
)

type RuntimeError struct {
	node asts.Node
	msg  string
}

func (e RuntimeError) Error() string {
	return e.msg
}

func (e RuntimeError) At() (int, int) {
	token := e.node.GetToken()
	return token.Line, token.Column
}

type Runtime struct {
	scope *Scope
	// errors []error
}

func NewRuntime() *Runtime {
	r := &Runtime{
		scope: NewScope(),
	}

	registerConstants(r.scope)
	registerFunctions(r.scope)

	return r
}

func (r *Runtime) Eval(node asts.Node) Object {
	obj := r.eval(r.scope, node)

	// if len(r.errors) > 0 {
	// 	return nil, r.errors[0]
	// }

	return obj
}

// func (r *Runtime) registerError(msg string, node asts.Node) {
// 	r.errors = append(r.errors, RuntimeError{
// 		msg:  msg,
// 		node: node,
// 	})
// }

func (r *Runtime) eval(env *Scope, node asts.Node) Object {
	// println("evaluating:", node.String(), reflect.TypeOf(node).String())
	// Push node in the scope

	switch node := node.(type) {
	case ast.Number:
		return r.evalNumber(env, node)

	case ast.Identifier:
		return r.evalIdentifier(env, node)

	case ast.Block:
		return r.evalBlock(env, node)

	case ast.UnaryOperator:
		return r.evalUnaryOperator(env, node)

	case ast.BinaryOperator:
		return r.evalBinaryOperator(env, node)

	case ast.Assignment:
		return r.evalAssignment(env, node)

	case ast.FunctionCall:
		return r.evalFunctionCall(env, node)

	case ast.FunctionDef:
		return r.evalFunctionDef(env, node)

	default:
		// TODO: handle error
		return nil

	}
}

func (r *Runtime) evalBlock(env *Scope, node ast.Block) Object {
	var result Object
	for _, statement := range node.Statements {
		result = r.eval(env, statement)
	}
	return result
}

func (r *Runtime) evalNumber(_ *Scope, node ast.Number) Object {
	return NewNumber(node.Value)
}

func (r *Runtime) evalIdentifier(env *Scope, node ast.Identifier) Object {
	value := env.Get(node.Name)
	if value == nil {

		// TODO: Panic
	}
	return value
}

func (r *Runtime) evalUnaryOperator(env *Scope, node ast.UnaryOperator) Object {
	// TODO: Consider panic cases
	// TODO: Consider type conversion
	right := r.eval(env, node.Expression).Number()

	switch node.Operator {
	case "+":
		return NewNumber(right)
	case "-":
		return NewNumber(-right)
	default:
		// TODO: Panic
		return nil
	}
}

func (r *Runtime) evalBinaryOperator(env *Scope, node ast.BinaryOperator) Object {
	// TODO: Consider panic cases
	// TODO: Consider type conversion
	left := r.eval(env, node.Left).Number()
	right := r.eval(env, node.Right).Number()

	var result float64 = 0
	switch node.Operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	case "%":
		result = float64(int(left) % int(right))
	case "^":
		result = math.Pow(left, right)
	}

	return NewNumber(result)
}

func (r *Runtime) evalAssignment(env *Scope, node ast.Assignment) Object {
	value := r.eval(env, node.Expression)
	env.Set(node.Identifier.Name, value)
	return value
}

func (r *Runtime) evalFunctionCall(env *Scope, node ast.FunctionCall) Object {
	fun := env.Get(node.Target.Name)

	if fun == nil {
		// TODO: Undefined function
	}

	args := make([]Object, len(node.Arguments))
	for i, arg := range node.Arguments {
		args[i] = r.eval(env, arg)
	}

	switch fun := fun.(type) {
	case *BuiltinFunction:
		return fun.Fn(env, args...)

	case *Function:
		scope := fun.Scope.New()
		// TODO: apply argument matching here
		// TODO: apply identifiers to scope here
		for i, arg := range fun.Matches[0].Args {
			a := args[i]
			identifier, ok := arg.(ast.Identifier)
			if ok {
				scope.Set(identifier.Name, a)
			}
		}

		return r.eval(scope, fun.Matches[0].Body)
	}

	// TODO: Undefined function

	return nil
}

func (r *Runtime) evalFunctionDef(env *Scope, node ast.FunctionDef) Object {
	// TODO: add matching validation here?
	var fun *Function
	storedFun := env.GetInScope(node.Name)
	if storedFun != nil {
		fn, ok := storedFun.(*Function)
		if ok {
			fun = fn
		}
	}

	if fun == nil {
		fun = NewFunction()
		fun.Scope = env
	}

	fun.AddMatch(node.Params, node.Body)
	env.Set(node.Name, fun)
	return fun
}
