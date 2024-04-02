package runtime

import (
	"math"

	"github.com/renatopp/klclang/internal/ast"
	"github.com/renatopp/langtools/asts"
)

type Runtime struct {
	scope *Scope
}

func NewRuntime() *Runtime {
	r := &Runtime{
		scope: NewScope(),
	}

	registerConstants(r.scope)

	return r
}

func (r *Runtime) Eval(node asts.Node) Object {
	return r.eval(r.scope, node)
}

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

	// case ast.Assignment:
	// 	return r.evalAssignment(env, node)

	// case ast.FunctionCall:
	// 	return r.evalFunctionCall(env, node)

	// case ast.FunctionDef:
	// 	return r.evalFunctionDef(env, node)

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
