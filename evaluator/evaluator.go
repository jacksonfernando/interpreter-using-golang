package evaluator

import (
	"github.com/v2/golang-intrepeter/ast"
	"github.com/v2/golang-intrepeter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
