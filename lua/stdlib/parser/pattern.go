package lua_parser

import (
	utils_lua "github.com/lucasew/doc_project/lua/utils"
	"github.com/redstorm-fyy/lpeg"
	lua "github.com/yuin/gopher-lua"
)


var PatternMetatable = map[string]lua.LGFunction{}
func init() {
    PatternMetatable["__pow"] = func (L *lua.LState) int {
        pattern, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        npow := L.CheckInt(2)
        L.Push(WrapPattern(L, pattern.Pow(npow)))
        return 1
    }
    PatternMetatable["__mul"] = func (L *lua.LState) int {
        a, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        b, err := normalizePattern(L.CheckAny(2))
        utils_lua.RaiseErrorIfNotNull(L, err)
        L.Push(WrapPattern(L, a.Seq(b)))
        return 1
    }
    PatternMetatable["__add"] = func (L *lua.LState) int {
        a, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        b, err := normalizePattern(L.CheckAny(2))
        utils_lua.RaiseErrorIfNotNull(L, err)
        L.Push(WrapPattern(L, a.Or(b)))
        return 1
    }
    PatternMetatable["__sub"] = func (L *lua.LState) int {
        a, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        b, err := normalizePattern(L.CheckAny(2))
        utils_lua.RaiseErrorIfNotNull(L, err)
        L.Push(WrapPattern(L, a.Sub(b)))
        return 1
    }
    PatternMetatable["__div"] = func (L *lua.LState) int {
        a, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        br := L.CheckAny(2)
        switch trivialized := utils_lua.Trivialize(br, true).(type) {
            case string:
                L.Push(WrapPattern(L, a.Sc(trivialized)))
            case int:
                L.Push(WrapPattern(L, a.Nc(trivialized)))
            case *lua.LTable:
                qt := lpeg.CaptureTable{}
                trivialized.ForEach(func (k lua.LValue, v lua.LValue) {
                    qt.Set(k.String(), utils_lua.Trivialize(v, true))
                })
                L.Push(WrapPattern(L, a.Qc(qt)))
            default:
                L.RaiseError("FATAL: dividing a pattern with a function or table is still not supported")
        }
        return 1
    }
    PatternMetatable["__unm"] = func (L *lua.LState) int {
        a, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        L.Push(WrapPattern(L, a.Not()))
        return 1
    }
    PatternMetatable["__len"] = func (L *lua.LState) int {
        a, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        L.Push(WrapPattern(L, a.And()))
        return 1
    }
    PatternMetatable["Sc"] = func (L *lua.LState) int {
        pattern, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        str := L.CheckString(2)
        L.Push(WrapPattern(L, pattern.Sc(str)))
        return 1
    }
    PatternMetatable["Nc"] = func (L *lua.LState) int {
        pattern, err := normalizePattern(L.CheckAny(1))
        utils_lua.RaiseErrorIfNotNull(L, err)
        num := L.CheckInt(2)
        L.Push(WrapPattern(L, pattern.Nc(num)))
        return 1
    }
    for k, v := range parserModule {
        PatternMetatable[k] = v
    }
    // TODO: Qc Fc
}

type LuaPattern struct {
    *lpeg.Pattern
}

func (LuaPattern) LuaType() string {
    return "Pattern"
}

func NewLuaPattern(pattern *lpeg.Pattern) LuaPattern {
    return LuaPattern{pattern}
}

func WrapPattern(L *lua.LState, pattern *lpeg.Pattern) lua.LValue {
    return utils_lua.WrapObject(L, NewLuaPattern(pattern), PatternMetatable)
}

func UnwrapPattern(table *lua.LTable) *lpeg.Pattern {
    return utils_lua.UnwrapObject(table).(LuaPattern).Pattern
}
