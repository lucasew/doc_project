package lua_stdlib

import (
	"log"
	"time"

	lua "github.com/yuin/gopher-lua"
	lua_json "layeh.com/gopher-json"
)

var luaLibs = map[string]lua.LGFunction{}

func RegisterLuaLib( loader lua.LGFunction, name string) {
    log.Printf("Registering module: %s", name)
    luaLibs[name] = loader
}

func init() {
    RegisterLuaLib(lua.OpenBase, lua.BaseLibName)
    RegisterLuaLib(lua.OpenString, lua.StringLibName)
    RegisterLuaLib(lua.OpenIo, lua.IoLibName)
    RegisterLuaLib(lua.OpenMath, lua.MathLibName)
    RegisterLuaLib(lua_json.Loader, "json")
}

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
    for k, v := range luaLibs {
        SetupLibrary(L, v, k)
    }
    return 0
}
