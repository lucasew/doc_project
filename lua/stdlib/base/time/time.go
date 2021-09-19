package lua_time

import (
	"time"

	app_lua "github.com/lucasew/doc_project/lua"
	lua "github.com/yuin/gopher-lua"
)

type LuaTime struct {
    time.Time
}

func (LuaTime) LuaType() string {
    return "Time"
}

func NewLuaTime(t time.Time) LuaTime {
    return LuaTime{t}
}

var TimeMetatable = map[string]lua.LGFunction{}

func init() {
    TimeMetatable["unix"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Unix())))
        return 1
    }
    TimeMetatable["day"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Day())))
        return 1
    }
    TimeMetatable["month"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Month())))
        return 1
    }
    TimeMetatable["year"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Year())))
        return 1
    }
    TimeMetatable["second"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Second())))
        return 1
    }
    TimeMetatable["minute"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Minute())))
        return 1
    }
    TimeMetatable["hour"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Minute())))
        return 1
    }
    TimeMetatable["weekday"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LNumber(float64(t.Weekday())))
        return 1
    }
    TimeMetatable["__tostring"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        L.Push(lua.LString(t.String()))
        return 1
    }
    TimeMetatable["__add"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        d := UnwrapDuration(L, L.CheckTable(2))
        L.Push(WrapTime(L, t.Add(d)))
        return 1
    }
    TimeMetatable["__sub"] = func(L *lua.LState) int {
        t := UnwrapTime(L, L.CheckTable(1))
        secondArgument := app_lua.UnwrapObject(L.CheckTable(2))
        switch v := secondArgument.(type) {
            case time.Duration:
                L.Push(WrapTime(L, t.Add(-v)))
            case time.Time:
                L.Push(WrapDuration(L, t.Sub(v)))
            default:
                L.RaiseError("second argument is neither time nor duration")
        }
        return 1
    }
}
func WrapTime(L *lua.LState, t time.Time) *lua.LTable {
    return app_lua.WrapObject(L, NewLuaTime(t), TimeMetatable)
}

func UnwrapTime(L *lua.LState, tbl *lua.LTable) time.Time {
    val, ok := app_lua.UnwrapObject(tbl).(LuaTime)
    if !ok {
        L.RaiseError("not a Time")
    }
    return val.Time
}
