package lua_new

import (
	"testing"

	utils_lua "github.com/lucasew/doc_project/lua/utils"
	lua "github.com/yuin/gopher-lua"
)

// FUN FACT: the test will just skip it if you do not run the binary created with `go test -c`
func TestRepl(t *testing.T) {
    vm := NewCommonState(lua.Options{})
    utils_lua.RunREPL(vm)
}
