package common

import (
	"slices"

	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Exit    key.Binding
	Next    key.Binding
	Prev    key.Binding
	Confirm key.Binding
	Help    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Exit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.Prev, k.Confirm},
		{k.Help, k.Exit},
	}
}

var GlobalKeys = keyMap{
	Prev: key.NewBinding(
		key.WithKeys("ctrl+p", "up", "shift+tab"),
		key.WithHelp("↑/ctrl+p/shift+tab", "move up"),
	),
	Next: key.NewBinding(
		key.WithKeys("ctrl+n", "down", "tab"),
		key.WithHelp("↓/ctrl+n/tab", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "toggle help"),
	),
	Exit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),
}

var (
	exitSeq    = []string{"esc", "ctrl+c"}
	nextSeq    = []string{"ctrl+n", "down", "tab"}
	prevSeq    = []string{"ctrl+p", "up", "shift+tab"}
	confirmSeq = []string{"enter"}

	AllSeq = slices.Concat(exitSeq, nextSeq, prevSeq, confirmSeq)
)
