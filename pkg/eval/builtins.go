package eval

import (
	"fmt"

	"github.com/juandroid007/palacinke/pkg/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(
			env *object.Environment,
			args ...object.Object,
		) object.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments. Got: %d, want: 1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Argument to `len` not supported, got: %s",
					args[0].Type())
			}
		},
	},
	"print": {
		Fn: func(
			env *object.Environment,
			args ...object.Object,
		) object.Object {
			for _, arg := range args {
				fmt.Fprint(env.GetOutput(), arg.Inspect())
			}
			return NIL
		},
	},
	"println": {
		Fn: func(
			env *object.Environment,
			args ...object.Object,
		) object.Object {
			for _, arg := range args {
				fmt.Fprint(env.GetOutput(), arg.Inspect())
			}
			fmt.Println()
			return NIL
		},
	},
}
