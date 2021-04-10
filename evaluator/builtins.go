package evaluator

import (
	"bufio"
	"fmt"
	"github.com/dpakach/pongo/object"
	"os"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)

			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},

	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"gets": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `gets` must be STRING, got %s", args[0].Type())
			}
			fmt.Print(args[0].(*object.String).Value)
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return newError("Error while reading data from Stdin")
			}
			return &object.String{Value: input}
		},
	},
	"str": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to `str` must be INTEGER, got %s", args[0].Type())
			}
			return &object.String{Value: strconv.Itoa(int(args[0].(*object.Integer).Value))}
		},
	},
	"int": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `int` must be STRING, got %s", args[0].Type())
			}
			strVal := args[0].(*object.String)
			val, err := strconv.Atoi(strVal.Value)
			if err != nil {
				return newError("Could not convert %s to int", strVal.Value)
			}
			return &object.Integer{Value: int64(val)}
		},
	},
	"type": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want=1", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	},
}
