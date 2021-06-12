package object

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		"len",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				switch arg := args[0].(type) {
				case *Array:
					return &Integer{Value: int64(len(arg.Elements))}
				case *String:
					return &Integer{Value: int64(len(arg.Value))}
				default:
					return newError("argument to `len` not supported, got %s", args[0].Type())
				}
			},
		},
	},
	{
		"first",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)

				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				}

				return nil
			},
		},
	},
	{
		"last",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)

				if length > 0 {
					return arr.Elements[length-1]
				}

				return nil
			},
		},
	},
	{
		"rest", &Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)

				if length > 0 {
					newElements := make([]Object, length-1, length-1)
					copy(newElements, arr.Elements[1:length])
					return &Array{Elements: newElements}
				}

				return nil
			},
		},
	},
	{
		"push", &Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)

				newElements := make([]Object, length+1, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &Array{Elements: newElements}
			},
		},
	},
	{
		"puts", &Builtin{
			Fn: func(args ...Object) Object {
				for _, arg := range args {
					fmt.Println(arg.Inspect())
				}

				return nil
			},
		},
	},
	{
		"gets", &Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				if args[0].Type() != STRING_OBJ {
					return newError("argument to `gets` must be STRING, got %s", args[0].Type())
				}
				fmt.Print(args[0].(*String).Value)
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					return newError("Error while reading data from Stdin")
				}
				return &String{Value: input}
			},
		},
	},
	{
		"str", &Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				if args[0].Type() != INTEGER_OBJ {
					return newError("argument to `str` must be INTEGER, got %s", args[0].Type())
				}
				return &String{Value: strconv.Itoa(int(args[0].(*Integer).Value))}
			},
		},
	},
	{
		"int", &Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				if args[0].Type() != STRING_OBJ {
					return newError("argument to `int` must be STRING, got %s", args[0].Type())
				}
				strVal := args[0].(*String)
				val, err := strconv.Atoi(strVal.Value)
				if err != nil {
					return newError("Could not convert %s to int", strVal.Value)
				}
				return &Integer{Value: int64(val)}
			},
		},
	},
	{
		"type", &Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got %d, want=1", len(args))
				}

				return &String{Value: string(args[0].Type())}
			},
		},
	},
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}
