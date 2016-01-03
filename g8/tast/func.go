package tast

import (
	"e8vm.io/e8vm/g8/types"
	"e8vm.io/e8vm/sym8"
)

// Func is a function.
type Func struct {
	Sym *sym8.Symbol // function symbol

	This     *types.Pointer // implicit receiver
	Receiver *sym8.Symbol   // explicit receiver

	Args      []*sym8.Symbol
	NamedRets []*sym8.Symbol

	Body []Stmt
}

// FuncAlias is a function alias.
type FuncAlias struct {
	Sym *sym8.Symbol
	Of  *sym8.Symbol
}
