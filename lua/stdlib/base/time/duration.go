package lua_time

import (
	"time"

	"github.com/lucasew/doc_project/lua/utils"
	lua "github.com/yuin/gopher-lua"
)

type LuaDuration struct {
    time.Duration
}

func (LuaDuration) LuaType() string {
    return "Duration"
}

func NewLuaDuration(d time.Duration) LuaDuration {
    return LuaDuration{d}
}

var DurationMetatable = map[string]lua.LGFunction{}
func init() {
    DurationMetatable["__tostring"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        L.Push(lua.LString(d.String()))
        return 1
    }
    DurationMetatable["hours"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        L.Push(lua.LNumber(d.Hours()))
        return 1
    }
    DurationMetatable["minutes"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        L.Push(lua.LNumber(d.Minutes()))
        return 1
    }
    DurationMetatable["seconds"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        L.Push(lua.LNumber(d.Seconds()))
        return 1
    }
    DurationMetatable["milliseconds"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        L.Push(lua.LNumber(d.Milliseconds()))
        return 1
    }
    DurationMetatable["microseconds"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        L.Push(lua.LNumber(d.Microseconds()))
        return 1
    }
    DurationMetatable["nanoseconds"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        L.Push(lua.LNumber(d.Nanoseconds()))
        return 1
    }
    DurationMetatable["round"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        d2 := UnwrapDuration(L, L.CheckTable(2))
        L.Push(lua.LNumber(d1.Round(d2)))
        return 1
    }
    DurationMetatable["truncate"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        d2 := UnwrapDuration(L, L.CheckTable(2))
        L.Push(lua.LNumber(d1.Truncate(d2)))
        return 1
    }
    DurationMetatable["__unm"] = func(L *lua.LState) int {
        d := UnwrapDuration(L, L.CheckTable(1))
        mn := -d
        L.Push(WrapDuration(L, mn))
        return 1
    }
    DurationMetatable["__add"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        d2 := UnwrapDuration(L, L.CheckTable(2))
        df := d1 + d2
        L.Push(WrapDuration(L, df))
        return 1
    }
    DurationMetatable["__sub"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        d2 := UnwrapDuration(L, L.CheckTable(2))
        df := d1 - d2
        L.Push(WrapDuration(L, df))
        return 1
    }
    DurationMetatable["__mul"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        by := L.CheckInt(2)
        df := d1 * time.Duration(by)
        L.Push(WrapDuration(L, df))
        return 1
    }
    DurationMetatable["__div"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        by := L.CheckInt(2)
        df := d1 / time.Duration(by)
        L.Push(WrapDuration(L, df))
        return 1
    }
    DurationMetatable["__eq"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        d2 := UnwrapDuration(L, L.CheckTable(2))
        L.Push(lua.LBool(d1 == d2))
        return 1
    }
    DurationMetatable["__lt"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        d2 := UnwrapDuration(L, L.CheckTable(2))
        L.Push(lua.LBool(d1 < d2))
        return 1
    }
    DurationMetatable["__le"] = func(L *lua.LState) int {
        d1 := UnwrapDuration(L, L.CheckTable(1))
        d2 := UnwrapDuration(L, L.CheckTable(2))
        L.Push(lua.LBool(d1 <= d2))
        return 1
    }
}

func WrapDuration(L *lua.LState, d time.Duration) *lua.LTable {
    return utils_lua.WrapObject(L, NewLuaDuration(d), DurationMetatable)
}

func UnwrapDuration(L *lua.LState, tbl *lua.LTable) time.Duration {
    val, ok := utils_lua.UnwrapObject(tbl).(LuaDuration)
    if !ok {
        L.RaiseError("not a Duration")
    }
    return val.Duration
}
