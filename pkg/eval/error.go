package eval

import (
	"fmt"

	"github.com/juandroid007/palacinke/pkg/object"
	"github.com/juandroid007/palacinke/pkg/token"
)

func newError(pos token.TokenPos, format string, a ...interface{}) *object.Error {
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
