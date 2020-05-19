package eval

import (
	"math"

	"github.com/juandroid007/palacinke/pkg/object"
)

func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
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
		return newError("Unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	if operator != "+" {
		return newError("Unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	return &object.String{Value: leftVal + rightVal}
}
