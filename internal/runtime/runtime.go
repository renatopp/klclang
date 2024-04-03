package runtime

import (
	"math"
	"strings"

	"github.com/renatopp/klclang/internal/ast"
	"github.com/renatopp/langtools/asts"
)

type Runtime struct {
	scope  *Scope
	errors []RuntimeError
}

func NewRuntime() *Runtime {
	r := &Runtime{
		scope:  NewScope(),
		errors: make([]RuntimeError, 0),
	}

	registerConstants(r.scope)
	registerFunctions(r.scope)

	return r
}

func (r *Runtime) HasErrors() bool {
	return len(r.errors) > 0
}

func (r *Runtime) Errors() []RuntimeError {
	return r.errors
}

func (r *Runtime) RegisterError(msg string, node asts.Node) Object {
	r.errors = append(r.errors, RuntimeError{
		msg:  msg,
		node: node,
	})
	return nil
}

func (r *Runtime) ClearErrors() {
	r.errors = make([]RuntimeError, 0)
}

func (r *Runtime) Scope() *Scope {
	return r.scope
}

func (r *Runtime) Eval(node asts.Node) Object {
	obj := r.eval(r.scope, node)

	// if len(r.errors) > 0 {
	// 	return nil, r.errors[0]
	// }

	return obj
}

func (r *Runtime) eval(env *Scope, node asts.Node) Object {
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
		r.RegisterError("unknown node type", node)
		return nil

	}
}

func (r *Runtime) evalBlock(env *Scope, node ast.Block) Object {
	var result Object
	for _, statement := range node.Statements {
		result = r.eval(env, statement)
		if result == nil {
			break
		}
	}
	return result
}

func (r *Runtime) evalNumber(_ *Scope, node ast.Number) Object {
	return NewNumber(node.Value)
}

func (r *Runtime) evalIdentifier(env *Scope, node ast.Identifier) Object {
	value := env.Get(node.Name)
	if value == nil {
		return r.RegisterError("undefined identifier", node)
	}
	return value
}

func (r *Runtime) evalUnaryOperator(env *Scope, node ast.UnaryOperator) Object {
	rightExpr := r.eval(env, node.Expression)
	if rightExpr == nil {
		return r.RegisterError("undefined expression", node)
	}

	right := rightExpr.Number()

	switch node.Operator {
	case "+":
		return NewNumber(right)
	case "-":
		return NewNumber(-right)
	default:
		return r.RegisterError("unknown unary operator", node)
	}
}

func (r *Runtime) evalBinaryOperator(env *Scope, node ast.BinaryOperator) Object {
	leftExpr := r.eval(env, node.Left)
	if leftExpr == nil {
		return r.RegisterError("undefined left expression", node)
	}

	rightExpr := r.eval(env, node.Right)
	if rightExpr == nil {
		return r.RegisterError("undefined right expression", node)
	}

	left := leftExpr.Number()
	right := rightExpr.Number()

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
	default:
		return r.RegisterError("unknown binary operator", node)
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
		return r.RegisterError("undefined function", node.Target)
	}

	args := make([]Object, len(node.Arguments))
	for i, arg := range node.Arguments {
		args[i] = r.eval(env, arg)
		if args[i] == nil {
			return nil
		}
	}

	switch fun := fun.(type) {
	case *BuiltinFunction:
		obj := fun.Fn(env, args...)
		if obj.Type() == ERROR {
			return r.RegisterError(obj.String(), node)
		}
		return obj

	case *Function:
		scope := fun.Scope.New()
		matchIdx := -1

		// Pattern matching
		// println("pattern matching:", node.Target.Name)
		argsStr := ""
		for _, arg := range args {
			argsStr += arg.String() + ", "
		}
		// println(">>>", argsStr)

		for idx, match := range fun.Matches {
			print(idx, " | ", padLeft(match.DebugString(), 100), " | ")
			if len(match.Params) != len(args) {
				// println("different number of arguments")
				continue
			}

			accepted := true
			for i, arg := range match.Params {
				_, ok := arg.(ast.Identifier)
				if ok {
					continue
				}

				n, ok := arg.(ast.Number)
				if ok {
					if n.Value == args[i].Number() {
						continue
					}
				}

				// println("unmatched value", n.Value, "!=", args[i].Number())
				accepted = false
				break
			}

			if accepted {
				// println("accepted!")
				matchIdx = idx
				break
			}
		}

		// Function call
		if matchIdx >= 0 {
			for i, arg := range fun.Matches[matchIdx].Params {
				a := args[i]
				identifier, ok := arg.(ast.Identifier)
				if ok {
					scope.Set(identifier.Name, a)
				}
			}
			return r.eval(scope, fun.Matches[matchIdx].Body)
		}

		return r.RegisterError("no matching function", node.Target)

	default:
		return r.RegisterError("unknown function type", node)
	}
}

func (r *Runtime) evalFunctionDef(env *Scope, node ast.FunctionDef) Object {
	var fun *Function
	storedFun := env.GetInScope(node.Name)
	if storedFun != nil {
		fn, ok := storedFun.(*Function)
		if ok {
			fun = fn
		}
	}

	if fun == nil {
		fun = NewFunction(env)
	}

	fun.AddMatch(node.Params, node.Body)
	// TODO: add matching validation here?
	env.Set(node.Name, fun)
	return fun
}

func padLeft(msg string, length int) string {
	if len(msg) >= length {
		return msg
	}

	return strings.Repeat(" ", length-len(msg)) + msg
}
