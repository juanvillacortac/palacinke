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
				return NewError(pos, "Wrong number of arguments. Got: %d, want: 1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.Hash:
				return &object.Integer{Value: int64(len(arg.Pairs))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return NewError(pos, "Argument to `len` not supported, got: %s",
					args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(
			env *object.Environment,
			pos token.TokenPos,
			args ...object.Object,
		) object.Object {
			if len(args) != 1 {
				return NewError(pos, "Wrong number of arguments. Got: %d, want: 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return NewError(pos, "Argument to `first` must be %s, got: %s",
					object.ARRAY_OBJ, args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NIL
		},
	},
	"last": {
		Fn: func(
			env *object.Environment,
			pos token.TokenPos,
			args ...object.Object,
		) object.Object {
			if len(args) != 1 {
				return NewError(pos, "Wrong number of arguments. Got: %d, want: 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return NewError(pos, "Argument to `last` must be %s, got: %s",
					object.ARRAY_OBJ, args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NIL
		},
	},
	"tail": {
		Fn: func(
			env *object.Environment,
			pos token.TokenPos,
			args ...object.Object,
		) object.Object {
			if len(args) != 1 {
				return NewError(pos, "Wrong number of arguments. Got: %d, want: 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return NewError(pos, "Argument to `last` must be %s, got: %s",
					object.ARRAY_OBJ, args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NIL
		},
	},
	"append": {
		Fn: func(
			env *object.Environment,
			pos token.TokenPos,
			args ...object.Object,
		) object.Object {
			if len(args) <= 1 {
				return NewError(pos, "Arguments to `append` must have least 2 arguments, got: %d", len(args))
			}
			switch args[0].Type() {
			case object.ARRAY_OBJ:
				arr := args[0].(*object.Array)
				length := len(arr.Elements)

				if args[1].Type() != object.ARRAY_OBJ {
					return NewError(
						pos,
						"Second argument to `append` must be ARRAY, got %s",
						args[1].Type(),
					)
				}

				newElements := make([]object.Object, length+1, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &object.Array{Elements: newElements}
			case object.HASH_OBJ:
				hash := args[0].(*object.Hash)

				toAppend, ok := args[1].(*object.Hash)
				if !ok {
					return NewError(
						pos,
						"Second argument to `append` must be HASH, got %s",
						args[1].Type(),
					)
				}

				newPairs := map[object.HashKey]object.HashPair{}

				for key, pair := range hash.Pairs {
					newPairs[key] = pair
				}
				for key, pair := range toAppend.Pairs {
					newPairs[key] = pair
				}

				return &object.Hash{Pairs: newPairs}
			default:
				return NewError(pos, "First argument to `append` must be ARRAY or HASH, got %s",
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
				switch arg.(type) {
				case *object.String:
					length := len(arg.Inspect())
					fmt.Fprint(env.GetOutput(), arg.Inspect()[1:length-1])
				default:
					fmt.Fprint(env.GetOutput(), arg.Inspect())
				}
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
				switch arg.(type) {
				case *object.String:
					length := len(arg.Inspect())
					fmt.Fprint(env.GetOutput(), arg.Inspect()[1:length-1])
				default:
					fmt.Fprint(env.GetOutput(), arg.Inspect())
				}
			}
			fmt.Println()
			return NIL
		},
	},
}
