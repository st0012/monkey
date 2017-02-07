package evaluator

import "github.com/st0012/monkey/object"

var builtins = map[string]*object.BuiltInFunction{
	"len": &object.BuiltInFunction{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				newError("wrong number of arguments. got=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
}
