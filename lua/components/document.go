package lua_components

import (
	"github.com/lucasew/doc_project/components"
	app_lua "github.com/lucasew/doc_project/lua"
	lua "github.com/yuin/gopher-lua"
)

type LuaDocument struct {
    components.Document
}

func (LuaDocument) LuaType() string {
    return "Document"
}

func NewLuaDocument(doc components.Document) LuaDocument {
    return LuaDocument{doc}
}

var DocumentMetatable = map[string]lua.LGFunction{
    "title": func(L *lua.LState) int {
        doc := UnwrapDocument(L, L.CheckTable(1))
        L.Push(lua.LString(doc.DocumentTitle()))
        return 1
    },
}

func WrapDocument(L *lua.LState, document components.Document) *lua.LTable {
    return app_lua.WrapObject(L, NewLuaDocument(document), DocumentMetatable)
}

func UnwrapDocument(L *lua.LState, tbl *lua.LTable) components.Document {
    val, ok := app_lua.UnwrapObject(tbl).(components.Document)
    if !ok {
        L.RaiseError("not a Document")
    }
    return val
}
