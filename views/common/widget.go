package common

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Widget interface {
	Init() tea.Cmd
	Update(tea.Msg) (Widget, tea.Cmd)
	View() string
	Help() help.KeyMap
}
