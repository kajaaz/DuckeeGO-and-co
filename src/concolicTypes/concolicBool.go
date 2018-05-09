package concolicTypes

import "../symTypes"
import "github.com/aclements/go-z3/z3"

type ConcolicBool struct {
  Value bool
  Sym   symTypes.SymBool
}

func (self ConcolicBool) equals(o interface{}) ConcolicBool {
  switch o.(type) {
  case bool:
    res := self.Value == bool(o)
    sym := self.Sym.z3Expr.Eq(ctx.FromBool(bool(o)))
  case ConcolicBool:
    res := self.Value == ConcolicBool(o).Value
    sym := self.Sym.z3Expr.Eq(ConcolicBool(o).Sym.z3Expr)
  default:
    reportError("cannot compare with == : incompatible types", self, o)
    // do something?
  }
  return ConcolicBool{Value: res, Sym: SymBool{sym}}
}


