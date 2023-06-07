package eval

import (
	"errors"
	"fmt"
	"klc/lang/ast"
	"klc/lang/obj"
	"math"
	"strings"
)

var TRUE = &obj.Number{Value: 1}
var FALSE = &obj.Number{Value: 0}

const DEBUG = false

type Evaluator struct {
	Env *Environment
}

func New() *Evaluator {
	return &Evaluator{Env: NewEnvironment()}
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

	default:
		e.throw("node not implemented: " + n.String())
	}

	return nil
}

//

func (e *Evaluator) evalBlock(n *ast.Block) obj.Object {
	var result obj.Object
	for _, stmt := range n.Statements {
		result = e.Eval(stmt)
	}
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
	e.Env.Set(id, value)
	return value
}

func (e *Evaluator) evalIdentifier(n *ast.Identifier) obj.Object {
	r, ok := e.Env.Get(n.Value)

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
		return FALSE
	}
	return TRUE
}

func (e *Evaluator) negateList(n *obj.List) obj.Object {
	for i, v := range n.Values {
		n.Values[i] = e.negate(v)
	}

	return n
}

func (e *Evaluator) negateString(n *obj.String) obj.Object {
	if n.AsBool() {
		return FALSE
	}
	return TRUE
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
		return ifReturn(left.Value == right.Value, TRUE, FALSE)

	case "!=":
		return ifReturn(left.Value != right.Value, TRUE, FALSE)

	case ">":
		return ifReturn(left.Value > right.Value, TRUE, FALSE)

	case "<":
		return ifReturn(left.Value < right.Value, TRUE, FALSE)

	case ">=":
		return ifReturn(left.Value >= right.Value, TRUE, FALSE)

	case "<=":
		return ifReturn(left.Value <= right.Value, TRUE, FALSE)

	case "&&":
		return ifReturn(left.AsBool() && right.AsBool(), TRUE, FALSE)

	case "||":
		return ifReturn(left.AsBool() || right.AsBool(), TRUE, FALSE)

	case "^^":
		return ifReturn(left.AsBool() != right.AsBool(), TRUE, FALSE)

	case "!|":
		return ifReturn(!left.AsBool() && !right.AsBool(), TRUE, FALSE)

	case "!&":
		return ifReturn(!left.AsBool() || !right.AsBool(), TRUE, FALSE)

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

func ifReturn[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}
