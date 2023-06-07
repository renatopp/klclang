package eval

import (
	"errors"
	"fmt"
	"klc/lang/ast"
	"klc/lang/obj"
	"math"
)

var TRUE = &obj.Number{Value: 1}
var FALSE = &obj.Number{Value: 0}

const DEBUG = false

func debug(a ...any) {
	if DEBUG {
		fmt.Println(a...)
	}
}

func throw(msg string, a ...any) {
	panic(fmt.Sprintf(msg, a...))
}

func SafeEval(n ast.Node) (r obj.Object, err error) {
	defer func() {
		if e := recover(); e != nil {
			r = nil
			err = errors.New(e.(string))
		}
	}()

	return Eval(n), nil
}

func Eval(n ast.Node) obj.Object {
	switch node := n.(type) {
	case *ast.Block:
		return evalBlock(node)

	case *ast.Number:
		return evalNumber(node)

	case *ast.List:
		return evalList(node)

	case *ast.String:
		return evalString(node)

	case *ast.UnaryOperation:
		right := Eval(node.Right)
		return evalUnaryOperation(node, right)

	case *ast.BinaryOperation:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalBinaryOperation(node, left, right)

	default:
		throw("node not implemented: " + n.String())
	}

	return nil
}

//

func evalBlock(n *ast.Block) obj.Object {
	var result obj.Object
	for _, stmt := range n.Statements {
		result = Eval(stmt)
	}
	return result
}

func evalNumber(n *ast.Number) obj.Object {
	return &obj.Number{Value: n.Value}
}

func evalList(n *ast.List) obj.Object {
	values := make([]obj.Object, len(n.Values))
	for i, v := range n.Values {
		values[i] = Eval(v)
	}
	return &obj.List{Values: values}
}

func evalString(n *ast.String) obj.Object {
	return &obj.String{Value: n.Value}
}

func evalUnaryOperation(n *ast.UnaryOperation, right obj.Object) obj.Object {
	switch n.Token.Literal {
	case "!":
		return negate(right)

	case "-":
		return bop("*", right, &obj.Number{Value: -1})

	case "+":
		return right
	}

	return nil
}

func evalBinaryOperation(n *ast.BinaryOperation, left, right obj.Object) obj.Object {
	debug("evalBinaryOperation", n.Token.Literal, left, right)
	return bop(n.Token.Literal, left, right)
}

func negate(n obj.Object) obj.Object {
	switch node := n.(type) {
	case *obj.Number:
		return negateNumber(node)

	case *obj.List:
		return negateList(node)

	case *obj.String:
		return negateString(node)
	}

	throw("invalid type to negate %s", n.Type())
	return nil
}

func negateNumber(n *obj.Number) obj.Object {
	if n.AsBool() {
		return FALSE
	}
	return TRUE
}

func negateList(n *obj.List) obj.Object {
	for i, v := range n.Values {
		n.Values[i] = negate(v)
	}

	return n
}

func negateString(n *obj.String) obj.Object {
	if n.AsBool() {
		return FALSE
	}
	return TRUE
}

func bop(op string, left, right obj.Object) obj.Object {
	debug("bot", op, left, right)
	lt := left.Type()
	rt := right.Type()

	if lt == obj.TNumber && rt == obj.TNumber {
		return bopNumberToNumber(op, left.(*obj.Number), right.(*obj.Number))
	} else if lt == obj.TNumber && rt == obj.TList {
		return bopNumberToList(op, left.(*obj.Number), right.(*obj.List), false)
	} else if lt == obj.TList && rt == obj.TNumber {
		return bopNumberToList(op, right.(*obj.Number), left.(*obj.List), true)
	} else if lt == obj.TList && rt == obj.TList {
		return bopListToList(op, left.(*obj.List), right.(*obj.List))
	}

	throw("invalid types: " + string(lt) + " " + op + " " + string(rt))
	return nil
}

func bopNumberToNumber(op string, left, right *obj.Number) obj.Object {
	debug("bopNumberToNumber", op, left, right)
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

	case "<":
		return ifReturn(left.Value < right.Value, TRUE, FALSE)

	case "<=":
		return ifReturn(left.Value <= right.Value, TRUE, FALSE)

	case ">":
		return ifReturn(left.Value > right.Value, TRUE, FALSE)

	case ">=":
		return ifReturn(left.Value >= right.Value, TRUE, FALSE)

	case "==":
		return ifReturn(left.Value == right.Value, TRUE, FALSE)

	case "!=":
		return ifReturn(left.Value != right.Value, TRUE, FALSE)
	}

	throw("invalid operator: " + op)
	return nil
}

func bopNumberToList(op string, left *obj.Number, right *obj.List, reverse bool) obj.Object {
	debug("bopNumberToList", op, left, right)
	ret := &obj.List{Values: make([]obj.Object, len(right.Values))}

	for i, v := range right.Values {
		if reverse {
			debug("reverse op", v)
			ret.Values[i] = bop(op, v, left)
		} else {
			debug("normal op", v)
			ret.Values[i] = bop(op, left, v)
		}
	}

	return ret
}

func bopListToList(op string, left, right *obj.List) obj.Object {
	debug("bopListToList", op, left, right)
	if len(left.Values) != len(right.Values) {
		throw("invalid list size")
		return nil
	}

	// TODO Concat operator

	ret := &obj.List{Values: make([]obj.Object, len(left.Values))}

	for i, v := range left.Values {
		ret.Values[i] = bop(op, v, right.Values[i])
	}

	return ret
}

func bopStringToString(op string, left, right *obj.String) obj.Object {
	// TODO Concat operator
	throw("Not implemented")
	return nil
}

func ifReturn[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}
