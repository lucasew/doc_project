package main

import (
	lua_new "github.com/lucasew/doc_project/lua/new"
	utils_lua "github.com/lucasew/doc_project/lua/utils"
	_ "github.com/lucasew/doc_project/lua/stdlib/base/time"
	lua "github.com/yuin/gopher-lua"
)

func main() {
    println("Welcome to the DocLUA REPL!")
    L := lua_new.NewCommonState(lua.Options{})
    utils_lua.RunREPL(L)
}
