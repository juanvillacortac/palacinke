package eval

import (
	"github.com/juanvillacortac/palacinke/pkg/object"
	"github.com/juanvillacortac/palacinke/pkg/token"
)

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NIL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object, pos token.TokenPos) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NewError(pos, "Unknown operator: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}
