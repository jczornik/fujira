package alert

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jczornik/fujira/views"
	"github.com/jczornik/fujira/views/messages"
)

var (
	style = lipgloss.
		NewStyle().
		Padding(1, 2).
		BorderStyle(
			lipgloss.RoundedBorder())
)

type model struct {
	msg      string
	nextView tea.Model
	nextCmd  tea.Cmd
	width    int
	height   int
}

func InitialModel(msg string, nextView tea.Model, nextCmd tea.Cmd) tea.Model {
	return model{msg: msg, nextView: nextView, nextCmd: nextCmd}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, views.GlobalKeys.Confirm) {
			return m.nextView, m.nextCmd
		}

	case messages.ResizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	content := style.Render(
		fmt.Sprintf("%s\n\n\n[ %s ]", m.msg, views.FocusedStyle.Render("OK")),
	)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
