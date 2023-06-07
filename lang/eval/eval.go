package eval

import (
	"errors"
	"fmt"
	"klc/lang/ast"
	"klc/lang/obj"
)

var TRUE = &obj.Number{Value: 1}
var FALSE = &obj.Number{Value: 0}

func SafeEval(n ast.Node) (r obj.Object, err error) {
	defer func() {
		if e := recover(); e != nil {
			r = nil
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()

	return Eval(n), nil
}

func Eval(n ast.Node) obj.Object {
	switch node := n.(type) {
	case *ast.Number:
		return &obj.Number{Value: node.Value}
	case *ast.String:
		return &obj.String{Value: node.Value}
	case *ast.UnaryOperation:
		right := Eval(node.Right)
		return evalUnaryOperation(node, right)
	case *ast.BinaryOperation:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalBinaryOperation(node, left, right)
	case *ast.Block:
		return evalBlock(node)
	}

	return nil
}

func evalBlock(n *ast.Block) obj.Object {
	var result obj.Object
	for _, stmt := range n.Statements {
		result = Eval(stmt)
	}
	return result
}

func evalUnaryOperation(n *ast.UnaryOperation, right obj.Object) obj.Object {
	switch n.Token.Literal {
	case "!":
		v := right.Bool()
		if v {
			return FALSE
		} else {
			return TRUE
		}

	case "-":
		if right.Type() != obj.TNumber {
			panic("invalid type")
		}

		return &obj.Number{Value: -right.(*obj.Number).Value}

	case "+":
		if right.Type() != obj.TNumber {
			panic("invalid type")
		}

		return &obj.Number{Value: right.(*obj.Number).Value}
	}

	return nil
}

func evalBinaryOperation(n *ast.BinaryOperation, left, right obj.Object) obj.Object {
	switch n.Token.Literal {
	case "+":
		if left.Type() == obj.TNumber && right.Type() == obj.TNumber {
			l := left.(*obj.Number).Value
			r := right.(*obj.Number).Value
			return &obj.Number{Value: l + r}
		}
	}

	return nil
}
