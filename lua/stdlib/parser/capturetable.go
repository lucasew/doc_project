package lua_parser

import (
	"github.com/redstorm-fyy/lpeg"
	lua "github.com/yuin/gopher-lua"
)


func wrapCaptureTableValue(L *lua.LState, item interface {}) lua.LValue {
    switch v := item.(type) {
        case string:
            return lua.LString(v)
        case int:
            return lua.LNumber(float64(v))
        case nil:
            return lua.LNil
        case lpeg.CaptureTable:
            return WrapCaptureTable(L, v)
        default:
            L.RaiseError("can't wrap capture table element of type %T: not implemented", v)
    }
    return lua.LNil
}

func WrapCaptureTable(L *lua.LState, ct lpeg.CaptureTable) lua.LValue {
    tbl := L.NewTable()
    lenCt := len(ct)
    isList := true
    for i := 1; i <= lenCt; i++ {
        _, ok := ct.At(i)
        if !ok {
            isList = false
        }
    }
    if isList {
        for i := 0; i <= lenCt; i++ {
            item, _ := ct.At(i)
            tbl.Append(wrapCaptureTableValue(L, item))
        }
    } else {
        for k, v := range ct {
            tbl.RawSet(wrapCaptureTableValue(L, k), wrapCaptureTableValue(L, v))
        }
    }
    return tbl
}

