package utils_lua

import (
	app_lua "github.com/lucasew/doc_project/lua"
	lua "github.com/yuin/gopher-lua"
)

func NewVoidLuaFunction(L *lua.LState, fn func ()) *lua.LFunction {
    return L.NewFunction(func (L *lua.LState) int {
        fn()
        return 0
    })
}

const userdataRefKey = lua.LString("_ref")

func WrapObject(L *lua.LState, object app_lua.CustomLuaType, methods map[string]lua.LGFunction) *lua.LTable {
    ud := L.NewUserData()
    ud.Value = object
    mt := L.NewTypeMetatable(object.LuaType())
    mt.RawSetString("__index", mt)
    L.SetFuncs(mt, methods)
    ret := L.NewTable()
    L.SetMetatable(ret, mt)
    ret.RawSetString("__index", mt)
    ret.RawSet(userdataRefKey, ud)
    return ret
}

func UnwrapObject(table *lua.LTable) interface{} {
    ref := table.RawGet(userdataRefKey)
    switch v := ref.(type) {
        case *lua.LUserData:
            return v.Value
        default:
            return nil
    }
}

func Trivialize(value lua.LValue, integerify bool) interface {} {
    switch value.Type() {
        case lua.LTTable:
            unwrapped := UnwrapObject(value.(*lua.LTable))
            if unwrapped != nil {
                return unwrapped
            }
            return value
        case lua.LTBool:
            return lua.LVAsBool(value)
        case lua.LTNil:
            return nil
        case lua.LTNumber:
            n := lua.LVAsNumber(value)
            if integerify {
                return int(n)
            } else {
                return float64(n)
            }
        case lua.LTString:
            n := lua.LVAsString(value)
            return n
    }
    return nil
}
