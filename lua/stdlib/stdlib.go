package lua_stdlib

import (
	"time"

	lua "github.com/yuin/gopher-lua"
    lua_json "layeh.com/gopher-json"
    lua_parser "github.com/lucasew/doc_project/lua/stdlib/parser"
    lua_time "github.com/lucasew/doc_project/lua/stdlib/base/time"
)

func SetupLibrary(L *lua.LState, lib func(*lua.LState) int, name string) {
    L.Push(L.NewFunction(lib))
    L.Push(lua.LString(name))
    L.Call(1, 0)
}

func LoadStdlib(L *lua.LState) int {
    L.SetGlobal("sleep", L.NewFunction(func (L *lua.LState) int {
        ms := L.CheckInt(1)
        time.Sleep(time.Millisecond*time.Duration(ms))
        return 0
    }))
    SetupLibrary(L, lua.OpenBase, lua.BaseLibName)
    SetupLibrary(L, lua.OpenString, lua.StringLibName)
    SetupLibrary(L, lua.OpenIo, lua.IoLibName)
    SetupLibrary(L, lua.OpenMath, lua.MathLibName)
    SetupLibrary(L, lua_time.OpenTime, "time")
    SetupLibrary(L, lua_parser.OpenParser, "lpeg")
    SetupLibrary(L, lua_json.Loader, "json")
    return 0
}
