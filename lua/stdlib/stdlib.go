package lua_stdlib

import (
	"time"

    "github.com/lucasew/doc_project/lua/base/time"
	lua "github.com/yuin/gopher-lua"
    luajson "layeh.com/gopher-json"
)

func LoadStdlib(L *lua.LState) int {
    L.SetGlobal("sleep", L.NewFunction(func (L *lua.LState) int {
        ms := L.CheckInt(1)
        time.Sleep(time.Millisecond*time.Duration(ms))
        return 0
    }))
    L.Push(L.NewFunction(lua.OpenBase))
    L.Push(lua.LString(lua.BaseLibName))
    L.Call(1, 0)
    L.Push(L.NewFunction(lua.OpenString))
    L.Push(lua.LString(lua.StringLibName))
    L.Call(1, 0)
    L.Push(L.NewFunction(lua.OpenIo))
    L.Push(lua.LString(lua.IoLibName))
    L.Call(1, 0)
    L.Push(L.NewFunction(lua.OpenMath))
    L.Push(lua.LString(lua.MathLibName))
    L.Call(1, 0)
    L.Push(L.NewFunction(lua_time.OpenTime))
    L.Push(lua.LString("time"))
    L.Call(1, 0)
    luajson.Loader(L)
    return 0
}
