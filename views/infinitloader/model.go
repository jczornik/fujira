package infinitloader

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jczornik/fujira/views"
)

type model struct {
	msg     string
	spinner spinner.Model
	action  tea.Cmd
}

type Finished struct {
	Model tea.Model
}

func InitialModel(msg string, action tea.Cmd) tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{msg: msg, spinner: s, action: action}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.action)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case Finished:
		return msg.Model, nil
	}

	gcmd := views.HandleGlobal(msg)

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, tea.Batch(cmd, gcmd)
}

func (m model) View() string {
	return fmt.Sprint(m.spinner.View(), m.msg)
}
