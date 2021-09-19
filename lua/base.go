package app_lua

import (
	"github.com/yuin/gopher-lua"
)

func NewVoidLuaFunction(L *lua.LState, fn func ()) *lua.LFunction {
    return L.NewFunction(func (L *lua.LState) int {
        fn()
        return 0
    })
}

const userdataRefKey = lua.LString("_ref")

type CustomLuaType interface {
    LuaType() string
}

func WrapObject(L *lua.LState, object CustomLuaType, methods map[string]lua.LGFunction) *lua.LTable {
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
    ud := table.RawGet(userdataRefKey).(*lua.LUserData)
    return ud.Value
}
