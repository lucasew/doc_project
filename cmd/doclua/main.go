package main

import (
	app_lua "github.com/lucasew/doc_project/lua"
	lua_new "github.com/lucasew/doc_project/lua/new"
	lua "github.com/yuin/gopher-lua"
)

func main() {
    println("Welcome to the DocLUA REPL!")
    L := lua_new.NewCommonState(lua.Options{})
    app_lua.RunREPL(L)
}
