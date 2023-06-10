package eval

import (
	"errors"
	"fmt"
	"klc/lang/ast"
	"klc/lang/builtins"
	"klc/lang/lexer"
	"klc/lang/obj"
	"klc/lang/parser"
	"math"
	"strings"
)

const DEBUG = false

const CHAIN_KEY = "0_chain"
const RETURN_KEY = "0_return"

type Evaluator struct {
	Stack *EnvironmentStack
}

func Run(code string) {
	eval := New()
	l := lexer.New(code)
	p := parser.New(l)
	prg, err := p.Parse()

	if err != nil {
		fmt.Println(err)
		return
	}

	r := eval.Eval(prg.Root)
	if r != nil {
		fmt.Println(r.AsString())
	}
}

func New() *Evaluator {
	ev := &Evaluator{
		Stack: NewStack(),
	}

	RegisterBuiltins(ev.Stack)

	return ev
}

func (e *Evaluator) debug(a ...any) {
	if DEBUG {
		fmt.Println(a...)
	}
}

func (e *Evaluator) throw(msg string, a ...any) {
	panic(fmt.Sprintf(msg, a...))
}

func (e *Evaluator) SafeEval(n ast.Node) (r obj.Object, err error) {
	defer func() {
		if e := recover(); e != nil {
			r = nil
			err = errors.New(e.(string))
		}
	}()

	return e.Eval(n), nil
}

func (e *Evaluator) Eval(n ast.Node) obj.Object {
	switch node := n.(type) {
	case *ast.Block:
		return e.evalBlock(node)

	case *ast.Number:
		return e.evalNumber(node)

	case *ast.List:
		return e.evalList(node)

	case *ast.String:
		return e.evalString(node)

	case *ast.Assignment:
		return e.evalAssignment(node)

	case *ast.Identifier:
		return e.evalIdentifier(node)

	case *ast.Index:
		return e.evalIndex(node)

	case *ast.UnaryOperation:
		right := e.Eval(node.Right)
		return e.evalUnaryOperation(node, right)

	case *ast.BinaryOperation:
		left := e.Eval(node.Left)
		right := e.Eval(node.Right)
		return e.evalBinaryOperation(node, left, right)

	case *ast.IfReturn:
		return e.evalIfReturn(node)

	case *ast.Conditional:
		return e.evalIfTrueFalse(node)

	case *ast.FunctionDef:
		return e.evalFuncDef(node)

	case *ast.FunctionCall:
		return e.evalFuncCall(node)

	case *ast.Chain:
		return e.evalChain(node)

	default:
		e.throw("node not implemented: " + n.String())
	}

	return nil
}

func (e *Evaluator) Call(fn obj.Callable, args []obj.Object) obj.Object {
	params := fn.GetParams()

	e.Stack.Push()
	defer e.Stack.Pop()

	tArgs := len(args)
	tParams := len(params)
	g := 0
	for i, v := range params {
		if !v.Spread {
			var value obj.Object
			if g >= tArgs {
				if v.Default == nil {
					e.throw("missing argument: %s", v.Name)
					return nil
				}
				value = v.Default
			} else {
				value = args[g]
			}

			g++
			e.Stack.Set(v.Name, value)
		} else {
			missing := tParams - i - 1

			total := 0
			sv := make([]obj.Object, 0)
			for j := i; j < (tArgs - missing); j++ {
				t := args[j]
				if t.Type() == obj.TList {
					sv = append(sv, t.(*obj.List).Values...)
				} else {
					sv = append(sv, t)
				}
				total++
			}

			g += total
			e.Stack.Set(v.Name, &obj.List{Values: sv})
		}
	}

	switch fun := fn.(type) {
	case *obj.BuiltinFunction:
		return fun.Fn(e, args...)

	case *obj.Function:
		return e.Eval(fun.Body)
	}

	e.throw("invalid function type")
	return nil
}

//

func (e *Evaluator) evalBlock(n *ast.Block) obj.Object {
	var result obj.Object
	for _, stmt := range n.Statements {
		result = e.Eval(stmt)

		if _, ok := e.Stack.Get(RETURN_KEY); ok {
			break
		}
	}
	e.Stack.Delete(RETURN_KEY)
	return result
}

func (e *Evaluator) evalNumber(n *ast.Number) obj.Object {
	return &obj.Number{Value: n.Value}
}

func (e *Evaluator) evalList(n *ast.List) obj.Object {
	values := make([]obj.Object, len(n.Values))
	for i, v := range n.Values {
		values[i] = e.Eval(v)
	}
	return &obj.List{Values: values}
}

func (e *Evaluator) evalString(n *ast.String) obj.Object {
	return &obj.String{Value: n.Value}
}

// ----------------------------------------------------------------------------
// FUNCTIONS
// ----------------------------------------------------------------------------
func (e *Evaluator) evalFuncDef(n *ast.FunctionDef) obj.Object {
	pms := make([]*obj.FunctionParam, len(n.Params))
	hasSpread := false
	hasDefault := false
	for i, v := range n.Params {
		p := &obj.FunctionParam{}
		f := false
		for !f {
			switch node := v.(type) {
			case *ast.Identifier:
				p.Name = node.Value
				f = true

			case *ast.DefaultArg:
				p.Default = e.Eval(node.Value)
				v = node.Identifier

			case *ast.SpreadArg:
				p.Spread = true
				v = node.Identifier

			default:
				f = true
			}
		}

		if p.Spread && p.Default != nil {
			e.throw("spread arguments cannot have default values: '%s'", p.Name)
			return nil
		}

		if p.Spread {
			if hasSpread {
				e.throw("only one spread argument is allowed: '%s'", p.Name)
				return nil
			}

			hasSpread = true
		}

		if p.Default != nil {
			if hasSpread {
				e.throw("default arguments cannot proceed spread arguments: '%s'", p.Name)
				return nil
			}
			hasDefault = true
		} else if hasDefault {
			e.throw("default arguments must be at the end")
			return nil
		}

		pms[i] = p
	}

	return &obj.Function{
		Params: pms,
		Body:   n.Block,
	}
}

func (e *Evaluator) evalFuncCall(n *ast.FunctionCall) obj.Object {
	target := e.Eval(n.Function)

	if target == nil {
		e.throw("undefined identifier: %s", n.Function)
		return nil
	}

	if target.Type() != obj.TFunction {
		e.throw("calling a non-function")
		return nil
	}

	args := make([]obj.Object, len(n.Arguments))
	for i, v := range n.Arguments {
		args[i] = e.Eval(v)
	}

	fn := target.(obj.Callable)
	params := fn.GetParams()

	e.Stack.Push()
	defer e.Stack.Pop()

	tArgs := len(args)
	tParams := len(params)
	g := 0
	for i, v := range params {
		if !v.Spread {
			var value obj.Object
			if g >= tArgs {
				if v.Default == nil {
					e.throw("missing argument: %s", v.Name)
					return nil
				}
				value = v.Default
			} else {
				value = args[g]
			}

			g++
			e.Stack.Set(v.Name, value)
		} else {
			missing := tParams - i - 1

			total := 0
			sv := make([]obj.Object, 0)
			for j := i; j < (tArgs - missing); j++ {
				t := args[j]
				if t.Type() == obj.TList {
					sv = append(sv, t.(*obj.List).Values...)
				} else {
					sv = append(sv, t)
				}
				total++
			}

			g += total
			e.Stack.Set(v.Name, &obj.List{Values: sv})
		}
	}

	switch fun := fn.(type) {
	case *obj.BuiltinFunction:
		return fun.Fn(e, args...)

	case *obj.Function:
		return e.Eval(fun.Body)
	}

	e.throw("invalid function type")
	return nil
}

func (e *Evaluator) evalChain(n *ast.Chain) obj.Object {
	prv, ok := e.Stack.Get(CHAIN_KEY)

	if ok { // piped
		switch call := n.Left.(type) {
		case *ast.FunctionCall:
			fn := &ast.FunctionCall{
				Function: call.Function,
				Arguments: append([]ast.Node{
					&ast.Identifier{Value: CHAIN_KEY},
				}, call.Arguments...),
			}
			prv = e.Eval(fn)
		default:
			e.throw("unexpected chaining operation")
			return nil
		}
	} else {
		prv = e.Eval(n.Left)
	}

	e.Stack.Set(CHAIN_KEY, prv)
	defer e.Stack.Delete(CHAIN_KEY)

	switch call := n.Right.(type) {
	case *ast.FunctionCall:
		fn := &ast.FunctionCall{
			Function: call.Function,
			Arguments: append([]ast.Node{
				&ast.Identifier{Value: CHAIN_KEY},
			}, call.Arguments...),
		}
		return e.Eval(fn)

	case *ast.Chain:
		return e.evalChain(call)

	default:
		e.throw("invalid chaining operation")
		return nil
	}
}

// ----------------------------------------------------------------------------
// EVAL ASSIGNMENT
// ----------------------------------------------------------------------------
func (e *Evaluator) evalAssignment(n *ast.Assignment) obj.Object {
	exp := e.Eval(n.Expression)

	if n.Token.Literal != "=" {
		exp = e.bop(n.Token.Literal, e.Eval(n.Identifier), exp)
	}

	switch node := n.Identifier.(type) {
	case *ast.Identifier:
		return e.assignToIdentifier(node.Value, exp)

	case *ast.Index:
		// return e.assignToIndex(node, exp)
	}

	e.throw("invalid assignment")
	return nil
}

func (e *Evaluator) assignToIdentifier(id string, value obj.Object) obj.Object {
	e.Stack.Set(id, value)
	return value
}

func (e *Evaluator) evalIdentifier(n *ast.Identifier) obj.Object {
	r, ok := e.Stack.Get(n.Value)

	if !ok {
		e.throw("undefined identifier: " + n.Value)
		return nil
	}

	return r
}

func (e *Evaluator) evalIndex(n *ast.Index) obj.Object {
	target := e.Eval(n.Target)
	index := e.Eval(n.Value)

	if index.Type() != obj.TNumber {
		e.throw("invalid index type: %s", index.Type())
		return nil
	}

	i := int(math.Floor(index.AsNumber()))

	switch node := target.(type) {
	case *obj.List:
		if i < 0 || i >= len(node.Values) {
			e.throw("index out of range: %d", i)
			return nil
		}
		return node.Values[i]
	case *obj.String:
		if i < 0 || i >= len(node.Value) {
			e.throw("index out of range: %d", i)
			return nil
		}
		return &obj.String{Value: string(node.Value[i])}
	}

	e.throw("index on invalid type: %s", target.Type())
	return nil
}

// ----------------------------------------------------------------------------
// EVAL UNARY OPERATION
// ----------------------------------------------------------------------------
func (e *Evaluator) evalUnaryOperation(n *ast.UnaryOperation, right obj.Object) obj.Object {
	switch n.Token.Literal {
	case "!":
		return e.negate(right)

	case "-":
		return e.bop("*", right, &obj.Number{Value: -1})

	case "+":
		return right
	}

	return nil
}

func (e *Evaluator) negate(n obj.Object) obj.Object {
	switch node := n.(type) {
	case *obj.Number:
		return e.negateNumber(node)

	case *obj.List:
		return e.negateList(node)

	case *obj.String:
		return e.negateString(node)
	}

	e.throw("invalid type to negate %s", n.Type())
	return nil
}

func (e *Evaluator) negateNumber(n *obj.Number) obj.Object {
	if n.AsBool() {
		return builtins.FALSE
	}
	return builtins.TRUE
}

func (e *Evaluator) negateList(n *obj.List) obj.Object {
	for i, v := range n.Values {
		n.Values[i] = e.negate(v)
	}

	return n
}

func (e *Evaluator) negateString(n *obj.String) obj.Object {
	if n.AsBool() {
		return builtins.FALSE
	}
	return builtins.TRUE
}

// ----------------------------------------------------------------------------
// EVAL BINARY OPERATION
// ----------------------------------------------------------------------------
func (e *Evaluator) evalBinaryOperation(n *ast.BinaryOperation, left, right obj.Object) obj.Object {
	e.debug("evalBinaryOperation", n.Token.Literal, left, right)
	return e.bop(n.Token.Literal, left, right)
}

func (e *Evaluator) bop(op string, left, right obj.Object) obj.Object {
	e.debug("bot", op, left, right)
	lt := left.Type()
	rt := right.Type()

	if lt == obj.TNumber && rt == obj.TNumber {
		return e.bopNumberToNumber(op, left.(*obj.Number), right.(*obj.Number))
	} else if lt == obj.TNumber && rt == obj.TList {
		return e.bopNumberToList(op, left.(*obj.Number), right.(*obj.List), false)
	} else if lt == obj.TList && rt == obj.TNumber {
		return e.bopNumberToList(op, right.(*obj.Number), left.(*obj.List), true)
	} else if lt == obj.TList && rt == obj.TList {
		return e.bopListToList(op, left.(*obj.List), right.(*obj.List))
	} else if lt == obj.TString && rt == obj.TString {
		return e.bopStringToString(op, left.(*obj.String), right.(*obj.String))
	}

	e.throw("invalid types: " + string(lt) + " " + op + " " + string(rt))
	return nil
}

func (e *Evaluator) bopNumberToNumber(op string, left, right *obj.Number) obj.Object {
	e.debug("bopNumberToNumber", op, left, right)
	switch op {
	case "+":
		return &obj.Number{Value: left.Value + right.Value}

	case "-":
		return &obj.Number{Value: left.Value - right.Value}

	case "*":
		return &obj.Number{Value: left.Value * right.Value}

	case "/":
		return &obj.Number{Value: left.Value / right.Value}

	case "%":
		return &obj.Number{Value: math.Mod(left.Value, right.Value)}

	case "**":
		return &obj.Number{Value: math.Pow(left.Value, right.Value)}

	case "//":
		return &obj.Number{Value: math.Floor(left.Value / right.Value)}

	case "==":
		return ifReturn(left.Value == right.Value, builtins.TRUE, builtins.FALSE)

	case "!=":
		return ifReturn(left.Value != right.Value, builtins.TRUE, builtins.FALSE)

	case ">":
		return ifReturn(left.Value > right.Value, builtins.TRUE, builtins.FALSE)

	case "<":
		return ifReturn(left.Value < right.Value, builtins.TRUE, builtins.FALSE)

	case ">=":
		return ifReturn(left.Value >= right.Value, builtins.TRUE, builtins.FALSE)

	case "<=":
		return ifReturn(left.Value <= right.Value, builtins.TRUE, builtins.FALSE)

	case "&&":
		return ifReturn(left.AsBool() && right.AsBool(), builtins.TRUE, builtins.FALSE)

	case "||":
		return ifReturn(left.AsBool() || right.AsBool(), builtins.TRUE, builtins.FALSE)

	case "^^":
		return ifReturn(left.AsBool() != right.AsBool(), builtins.TRUE, builtins.FALSE)

	case "!|":
		return ifReturn(!left.AsBool() && !right.AsBool(), builtins.TRUE, builtins.FALSE)

	case "!&":
		return ifReturn(!left.AsBool() || !right.AsBool(), builtins.TRUE, builtins.FALSE)

	case "++":
		return &obj.List{Values: []obj.Object{left, right}}
	}

	e.throw("invalid operator: " + op)
	return nil
}

func (e *Evaluator) bopNumberToList(op string, left *obj.Number, right *obj.List, reverse bool) obj.Object {
	e.debug("bopNumberToList", op, left, right)
	if op == "++" {
		ret := &obj.List{Values: make([]obj.Object, len(right.Values)+1)}
		if reverse {
			copy(ret.Values, right.Values)
			ret.Values[len(right.Values)] = left
		} else {
			copy(ret.Values[1:], right.Values)
			ret.Values[0] = left
		}
		return ret
	}

	ret := &obj.List{Values: make([]obj.Object, len(right.Values))}

	for i, v := range right.Values {
		if reverse {
			e.debug("reverse op", v)
			ret.Values[i] = e.bop(op, v, left)
		} else {
			e.debug("normal op", v)
			ret.Values[i] = e.bop(op, left, v)
		}
	}

	return ret
}

func (e *Evaluator) bopListToList(op string, left, right *obj.List) obj.Object {
	e.debug("bopListToList", op, left, right)

	if op == "++" {
		ret := &obj.List{Values: make([]obj.Object, len(left.Values)+len(right.Values))}
		copy(ret.Values, left.Values)
		copy(ret.Values[len(left.Values):], right.Values)
		return ret
	}

	if len(left.Values) != len(right.Values) {
		e.throw("different list size: %d against %d", len(left.Values), len(right.Values))
		return nil
	}

	ret := &obj.List{Values: make([]obj.Object, len(left.Values))}

	for i, v := range left.Values {
		ret.Values[i] = e.bop(op, v, right.Values[i])
	}

	return ret
}

func (e *Evaluator) bopStringToString(op string, left, right *obj.String) obj.Object {
	e.debug("bopStringToString", op, left, right)

	if op == "++" {
		return &obj.String{Value: strings.Join([]string{left.Value, right.Value}, "")}
	}

	e.throw("invalid operation over strings: " + op)
	return nil
}

// ----------------------------------------------------------------------------
// EVAL CONDITIONAL
// ----------------------------------------------------------------------------
func (e *Evaluator) evalIfReturn(n *ast.IfReturn) obj.Object {
	cond := e.Eval(n.Condition)
	if cond.AsBool() {
		v := e.Eval(n.Return)
		e.Stack.Set(RETURN_KEY, v)
		return v
	}

	return nil
}

func (e *Evaluator) evalIfTrueFalse(n *ast.Conditional) obj.Object {
	cond := e.Eval(n.Condition)
	if cond.AsBool() {
		return e.Eval(n.True)
	}

	return e.Eval(n.False)
}

// ----------------------------------------------------------------------------
// UTILS
// ----------------------------------------------------------------------------
func ifReturn[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}
