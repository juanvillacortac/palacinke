package eval

import (
	"fmt"

	"github.com/juanvillacortac/palacinke/pkg/object"
	"github.com/juanvillacortac/palacinke/pkg/token"
)

func NewError(pos token.TokenPos, format string, a ...interface{}) *object.Error {
	msg := fmt.Sprintf(format, a...)
	return &object.Error{
		Message: fmt.Sprintf("[%d:%d] %s", pos.Line, pos.Col, msg),
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
