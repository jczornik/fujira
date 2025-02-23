package views

import (
	"slices"
)

const (
	ExitInput = iota
	NextInput
	PrevInput
	ConfirmInput
	UnknownInput
)

var (
	exitSeq    = []string{"esc", "ctrl+c"}
	nextSeq    = []string{"ctrl+n", "down", "tab"}
	prevSeq    = []string{"ctrl+p", "up", "shift+tab"}
	confirmSeq = []string{"enter"}

	AllSeq = slices.Concat(exitSeq, nextSeq, prevSeq, confirmSeq)
)

func MapKeyToType(key string) uint8 {
	isIn := func(seq []string) bool {
		return slices.Contains(seq, key)
	}

	switch {
	case isIn(exitSeq):
		return ExitInput

	case isIn(nextSeq):
		return NextInput

	case isIn(prevSeq):
		return PrevInput

	case isIn(confirmSeq):
		return ConfirmInput

	default:
		return UnknownInput
	}
}
