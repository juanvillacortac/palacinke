package eval

import (
	"math"

	"github.com/juandroid007/palacinke/pkg/object"
	"github.com/juandroid007/palacinke/pkg/token"
)

func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
	pos token.TokenPos,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "^":
		val := int64(math.Pow(float64(leftVal), float64(rightVal)))
		return &object.Integer{Value: val}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return booleanInstances[leftVal < rightVal]
	case ">":
		return booleanInstances[leftVal > rightVal]
	case "==":
		return booleanInstances[leftVal == rightVal]
	case "!=":
		return booleanInstances[leftVal != rightVal]
	case "<=":
		return booleanInstances[leftVal <= rightVal]
	case ">=":
		return booleanInstances[leftVal >= rightVal]
	default:
		return NewError(pos, "Unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
	pos token.TokenPos,
) object.Object {
	if operator != "+" {
		return NewError(pos, "Unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	return &object.String{Value: leftVal + rightVal}
}
