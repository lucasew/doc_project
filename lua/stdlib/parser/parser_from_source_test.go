package lua_parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	lua_new "github.com/lucasew/doc_project/lua/new"
	lua "github.com/yuin/gopher-lua"
)

func TestFromLPEGTestcases(t *testing.T) {
    L := lua_new.NewCommonState(lua.Options{IncludeGoStackTrace: true})
    lua.OpenTable(L)
    L.SetGlobal("spew", L.NewFunction(func (L *lua.LState) int {
        spew.Dump(L.CheckAny(1))
        return 0
    }))
    err := L.DoFile("./test.lua")
    if err != nil {
        t.Error(err)
    }
}
