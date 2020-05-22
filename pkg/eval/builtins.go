package eval

import (
	"fmt"

	"github.com/juandroid007/palacinke/pkg/object"
	"github.com/juandroid007/palacinke/pkg/token"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(
			env *object.Environment,
			pos token.TokenPos,
			args ...object.Object,
		) object.Object {
			if len(args) != 1 {
				return newError(pos, "Wrong number of arguments. Got: %d, want: 1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError(pos, "Argument to `len` not supported, got: %s",
					args[0].Type())
			}
		},
	},
	"print": {
		Fn: func(
			env *object.Environment,
			pos token.TokenPos,
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
			pos token.TokenPos,
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
