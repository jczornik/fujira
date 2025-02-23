package alert

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jczornik/fujira/views"
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
		switch views.MapKeyToType(msg.String()) {
		case views.ConfirmInput:
			return m.nextView, m.nextCmd
		}
	}

	return m, views.HandleGlobal(msg)
}

func (m model) View() string {
	content := style.Render(
		fmt.Sprintf("%s\n\n\n[ %s ]", m.msg, views.FocusedStyle.Render("OK")),
	)
	width, height := views.GetDim()
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}
