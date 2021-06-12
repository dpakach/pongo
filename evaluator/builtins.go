package evaluator

import (
	"github.com/dpakach/pongo/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"last":  object.GetBuiltinByName("last"),
	"rest":  object.GetBuiltinByName("rest"),
	"push":  object.GetBuiltinByName("push"),
	"puts":  object.GetBuiltinByName("puts"),
	"gets":  object.GetBuiltinByName("gets"),
	"str":   object.GetBuiltinByName("str"),
	"int":   object.GetBuiltinByName("int"),
	"type":  object.GetBuiltinByName("type"),
}
