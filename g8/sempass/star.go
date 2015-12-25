package sempass

import (
	"e8vm.io/e8vm/g8/ast"
	"e8vm.io/e8vm/g8/tast"
	"e8vm.io/e8vm/g8/types"
)

func buildStarExpr(b *Builder, expr *ast.StarExpr) tast.Expr {
	opPos := expr.Star.Pos

	addr := b.BuildExpr(expr.Expr)
	if addr == nil {
		return nil
	}

	addrRef := addr.R()
	if addrRef.List != nil {
		b.Errorf(opPos, "* on expression list")
		return nil
	}
	if t, ok := addrRef.T.(*types.Type); ok {
		r := tast.NewRef(&types.Type{&types.Pointer{t}})
		return &tast.Type{r}
	}

	t, ok := addrRef.T.(*types.Pointer)
	if !ok {
		b.Errorf(opPos, "* on non-pointer")
		return nil
	}

	r := tast.NewAddressableRef(t)
	return &tast.StarExpr{addr, r}
}
