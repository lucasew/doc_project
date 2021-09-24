package lua_time

import (
	"time"

	lua "github.com/yuin/gopher-lua"
    "github.com/lucasew/doc_project/lua/stdlib"
)

func init() {
    lua_stdlib.RegisterLuaLib(OpenTime, "time")
}

func OpenTime(L *lua.LState) int {
    varname := L.CheckString(1)
    module := L.NewTable()
    module.RawSetString("microsecond", WrapDuration(L, time.Microsecond))
    module.RawSetString("millisecond", WrapDuration(L, time.Millisecond))
    module.RawSetString("second", WrapDuration(L, time.Second))
    module.RawSetString("minute", WrapDuration(L, time.Minute))
    module.RawSetString("hour", WrapDuration(L, time.Hour))

    module.RawSetString("now", L.NewFunction(LuaNow))

    module.RawSetString("duration_from_string", L.NewFunction(LuaDurationFromString))
    module.RawSetString("time_from_string", L.NewFunction(LuaTimeFromString))

    L.SetGlobal(varname, module)
    return 0
}

func LuaNow(L *lua.LState) int {
    t := time.Now()
    L.Push(WrapTime(L, t))
    return 1
}

func LuaDurationFromString(L *lua.LState) int {
    stmt := L.CheckString(1)
    duration, err := time.ParseDuration(stmt)
    if err != nil {
        L.RaiseError(err.Error())
    }
    L.Push(WrapDuration(L, duration))
    return 1
}

func LuaTimeFromString(L *lua.LState) int {
    layout := L.CheckString(1)
    stmt := L.CheckString(2)
    t, err := time.Parse(layout, stmt)
    if err != nil {
        L.RaiseError(err.Error())
    }
    L.Push(WrapTime(L, t))
    return 1
}
