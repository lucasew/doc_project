package lua_parser

import (
	"errors"
	"fmt"
	"log"

	// "github.com/davecgh/go-spew/spew"
	lua_stdlib "github.com/lucasew/doc_project/lua/stdlib"
	utils_lua "github.com/lucasew/doc_project/lua/utils"
	"github.com/redstorm-fyy/lpeg"
	lua "github.com/yuin/gopher-lua"
)

func init() {
    lua_stdlib.RegisterLuaLib(OpenParser, "lpeg")
}

func normalizePattern(value lua.LValue) (*lpeg.Pattern, error) {
    trivial := utils_lua.Trivialize(value, true)
    switch v := trivial.(type) {
        case LuaPattern:
            return v.Pattern, nil
        case *lua.LTable:
            ret, err := normalizeGrammar(value)
            if err != nil {
                return nil, err
            }
            return ret, nil
        case nil:
            return nil, errors.New("pattern can't be nil")
        default:
            parsed := lpeg.P(v)
            return parsed, nil
    }
}

func normalizeGrammar(value lua.LValue) (*lpeg.Pattern, error) {
    var initialSymbol string
    switch typedValue := utils_lua.Trivialize(value, true).(type) {
        case *lua.LTable:
            switch firstValue := utils_lua.Trivialize(typedValue.RawGetInt(1), true).(type) {
                case LuaPattern:
                    var g lpeg.Grammar
                    g.AddRule("1", firstValue.Pattern)
                    return lpeg.P(g), nil
                case string, int, int64, float64:
                    initialSymbol = fmt.Sprintf("%v", firstValue)
                default:
                    return nil, fmt.Errorf("invalid value for first table value: %T", firstValue)
            }
            var g lpeg.Grammar
            var err error
            g.AddRule("__init__", lpeg.V(initialSymbol))
            typedValue.ForEach(func (k lua.LValue, v lua.LValue) {
                valueAsPattern, err := normalizePattern(v)
                if err != nil {
                    err = fmt.Errorf("invalid key '%s' in table parameter: %s", k.String(), err.Error())
                }
                g.AddRule(k.String(), valueAsPattern)
            })
            if err != nil {
                return nil, err
            }
            return lpeg.P(g), nil
    }
    return nil, errors.New("Invalid input value")
}

var parserModule = map[string]lua.LGFunction{
    "version": lpegVersionFn,
    "P": lpegP,
    "S": lpegS,
    "R": lpegR,
    "V": lpegV,
    "B": lpegB,
    "C": lpegC,
    "Cs": lpegCs,
    "Ct": lpegCt,
    "Cf": utils_lua.LuaUnimplemented,
    "Cgn": utils_lua.LuaUnimplemented,
    "Cc": lpegCc,
    "Cmt": utils_lua.LuaUnimplemented,
    "Carg": utils_lua.LuaUnimplemented,
    "Cg": lpegCg,
    "Cp": lpegCp,
    "type": lpegType,
    "match": lpegMatch,
    "locale": lpegLocale,
}


func OpenParser(L *lua.LState) int {
    log.Printf("WARNING: the lpeg parser is still not ready")
    varname := L.CheckString(1)
    module := L.NewTable()
    L.SetFuncs(module, parserModule)
    L.SetGlobal(varname, module)
    L.Push(module)
    return 1
}




func lpegP(L *lua.LState) int {
    raw := L.CheckAny(1)
    pattern, err := normalizePattern(raw)
    utils_lua.RaiseErrorIfNotNull(L, err)
    L.Push(WrapPattern(L, pattern))
    return 1
}

func lpegS(L *lua.LState) int {
    str := L.CheckString(1)
    L.Push(WrapPattern(L, lpeg.S(str)))
    return 1
}

func lpegR(L *lua.LState) int {
    args := make([]string, 0, L.GetTop())
    for i := 1; i <= L.GetTop(); i++ {
        item := L.CheckString(i)
        args = append(args, item)
    }
    L.Push(WrapPattern(L, lpeg.R(args...)))
    return 1
}

func lpegV(L *lua.LState) int {
    value := L.CheckAny(1).String()
    L.Push(WrapPattern(L, lpeg.V(value)))
    return 1
}

func lpegB(L *lua.LState) int {
    pattern, err := normalizePattern(L.CheckAny(1))
    utils_lua.RaiseErrorIfNotNull(L, err)
    L.Push(WrapPattern(L, lpeg.B(pattern)))
    return 1
}

func lpegC(L *lua.LState) int {
    pattern, err := normalizePattern(L.CheckAny(1))
    utils_lua.RaiseErrorIfNotNull(L, err)
    L.Push(WrapPattern(L, lpeg.C(pattern)))
    return 1
}

func lpegCs(L *lua.LState) int {
    pattern, err := normalizePattern(L.CheckAny(1))
    utils_lua.RaiseErrorIfNotNull(L, err)
    L.Push(WrapPattern(L, lpeg.Cs(pattern)))
    return 1
}

func lpegCt(L *lua.LState) int {
    pattern, err := normalizePattern(L.CheckAny(1))
    utils_lua.RaiseErrorIfNotNull(L, err)
    L.Push(WrapPattern(L, lpeg.Ct(pattern)))
    return 1
}

func lpegCg(L *lua.LState) int {
    pattern, err := normalizePattern(L.CheckAny(1))
    utils_lua.RaiseErrorIfNotNull(L, err)
    L.Push(WrapPattern(L, lpeg.Cg(pattern)))
    return 1
}

func lpegCp(L *lua.LState) int {
    L.Push(WrapPattern(L, lpeg.Cp()))
    return 1
}

func lpegVersionFn(L *lua.LState) int {
    L.Push(lua.LString("1.0.2"))
    return 1
}

func lpegType(L *lua.LState) int {
    any_obj := L.CheckAny(1)
    if any_obj.Type() != lua.LTTable {
        L.Push(lua.LNil)
    } else {
        table := L.CheckTable(1)
        _, ok := utils_lua.UnwrapObject(table).(LuaPattern)
        if ok {
            L.Push(lua.LString("pattern"))
        } else {
            L.Push(lua.LNil)
        }
    }
    return 1
}

func lpegMatch(L *lua.LState) int {
    pattern, err := normalizePattern(L.CheckAny(1))
    utils_lua.RaiseErrorIfNotNull(L, err)
    subject := L.CheckString(2)
    init := L.OptNumber(3, 1)
    idx, ret := lpeg.Match(pattern, subject, init)
    if ret != nil {
        num := 0
        for ; num < len(ret); num++ {
            switch tv := ret[num].(type) {
                case string:
                    L.Push(lua.LString(tv))
                case int:
                    L.Push(lua.LNumber(float64(tv)))
                case nil:
                    L.Push(lua.LNil)
                case lpeg.CaptureTable:
                    L.Push(WrapCaptureTable(L, tv))
                default:
                    L.RaiseError("match: type %T still not supported for return", tv)
            }
        }
        return num
    }
    if (idx == -1) {
        L.Push(lua.LNil)
    } else {
        L.Push(lua.LNumber(float64(idx)))
    }
    return 1
}

// TODO: Implement adding what should be added
func lpegLocale(L *lua.LState) int {
    ret := L.OptTable(1, L.NewTable())
    // Add what should be added
    L.Push(ret)
    return 1
}

func lpegCc(L *lua.LState) int {
    values := make([]interface{}, 0, L.GetTop())
    for i := 1; i <= L.GetTop(); i++ {
        item := utils_lua.Trivialize(L.CheckAny(i), true)
        values = append(values, item)
    }
    L.Push(WrapPattern(L, lpeg.Cc(values...)))
    return 1
}
