package lua_new

import (
    lua "github.com/yuin/gopher-lua"
    "github.com/lucasew/doc_project/lua/stdlib"
)

func NewCommonState(opts lua.Options) *lua.LState {
    opts.SkipOpenLibs = true
    L := lua.NewState(opts)
    L.Push(L.NewFunction(lua_stdlib.LoadStdlib))
    L.Call(0, 0)
    return L
}
