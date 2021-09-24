package lua_new

import (
	"github.com/lucasew/doc_project/lua/stdlib"
	lua "github.com/yuin/gopher-lua"
)

func NewCommonState(opts lua.Options) *lua.LState {
    opts.SkipOpenLibs = true
    L := lua.NewState(opts)
    L.Push(L.NewFunction(lua_stdlib.LoadStdlib))
    L.Call(0, 0)
    return L
}
