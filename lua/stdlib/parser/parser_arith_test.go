package lua_parser

import (
	"testing"

	lua_new "github.com/lucasew/doc_project/lua/new"
	lua "github.com/yuin/gopher-lua"
)

const arithTestCode = `
-- lpeg = require 'lpeg'

-- From: http://www.inf.puc-rio.br/~roberto/lpeg/
-- Lexical Elements
local Space = lpeg.S(" \n\t")^0
local Number = lpeg.C(lpeg.P"-"^-1 * lpeg.R("09")^1) * Space
local TermOp = lpeg.C(lpeg.S("+-")) * Space
local FactorOp = lpeg.C(lpeg.S("*/")) * Space
local Open = "(" * Space
local Close = ")" * Space

-- Grammar
local Exp, Term, Factor = lpeg.V"Exp", lpeg.V"Term", lpeg.V"Factor"
G = lpeg.P{ Exp,
  Exp = lpeg.Ct(Term * (TermOp * Term)^0);
  Term = lpeg.Ct(Factor * (FactorOp * Factor)^0);
  Factor = Number + Open * Exp * Close;
}

G = Space * G * -1

-- Evaluator
function eval (x)
  if type(x) == "string" then
    return tonumber(x)
  else
    local op1 = eval(x[1])
    for i = 2, #x, 2 do
      local op = x[i]
      local op2 = eval(x[i + 1])
      if (op == "+") then op1 = op1 + op2
      elseif (op == "-") then op1 = op1 - op2
      elseif (op == "*") then op1 = op1 * op2
      elseif (op == "/") then op1 = op1 / op2
      end
    end
    return op1
  end
end

-- Parser/Evaluator
function evalExp (s)
  local t = lpeg.match(G, s)
  if not t then error("syntax error", 2) end
  return eval(t)
end

for i = 0, 10 do
    x = math.random(0, 1000)
    y = math.random(1, 1000)
    op = math.random(1, 4)
    local ops = ({
        x + y,
        x - y,
        x * y,
        x / y
    })[op]
    local opsym = ({
        "+",
        "-",
        "*",
        "/"
    })[op]
    expr = tostring(x) .. " " .. opsym .. " " .. tostring(y)
    expected = ops
    got = evalExp(expr)
    if expected == got then
        print("OK")
    else
        error("parse failed: expected: ".. expected .. " got: " .. got)
    end
end
`

func TestParserArith(t *testing.T) {
    t.Skipf("parser test skipped because it's not implemented at all")
    L := lua_new.NewCommonState(lua.Options{})
    err := L.DoString(arithTestCode)
    if err != nil {
        t.Errorf(err.Error())
    }
}
