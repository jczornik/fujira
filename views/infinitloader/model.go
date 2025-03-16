package infinitloader

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jczornik/fujira/views/common"
)

type model struct {
	msg     string
	spinner spinner.Model
	action  tea.Cmd
}

type Finished struct {
	Model common.Widget
}

func InitialModel(msg string, action tea.Cmd) common.Widget {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{msg: msg, spinner: s, action: action}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.action)
}

func (m model) Update(msg tea.Msg) (common.Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case Finished:
		return msg.Model, msg.Model.Init()
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprint(m.spinner.View(), m.msg)
}

func (m model) Help() help.KeyMap {
	return nil
}
