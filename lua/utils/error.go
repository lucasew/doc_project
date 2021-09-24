package utils_lua

import lua "github.com/yuin/gopher-lua"

func RaiseErrorIfNotNull(L *lua.LState, err error) {
    if err != nil {
        L.RaiseError(err.Error())
    }
}

func LuaUnimplemented(L *lua.LState) int {
    L.RaiseError("unimplemented")
    return 0
}
