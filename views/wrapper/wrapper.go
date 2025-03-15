package views

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jczornik/fujira/views/common"
)

const (
	marginTop    = 3
	marginBottom = 1
	marginRight  = 1
	marginLeft   = 1

	paddingTop    = 0
	paddingBottom = 0
	paddingRight  = 0
	paddingLeft   = 0

	helpMarginTop = 1
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	internal common.Widget
	viewport viewport.Model
	help     help.Model
	width    int
	height   int
}

func calcHelpHeight(h help.Model) int {
	return strings.Count(h.View(common.GlobalKeys), "\n") + 1
}

func (m *model) calcInternalSize() (int, int) {
	width := m.width - marginLeft - marginRight - paddingLeft - paddingRight
	height := m.height - marginTop - marginBottom - paddingTop - paddingBottom

	return width, height
}

func (m *model) updateViewportDim() {
	width, height := m.calcInternalSize()
	height -= calcHelpHeight(m.help)

	m.viewport.Width = width
	m.viewport.Height = height
}

func InitialModel(initial common.Widget) tea.Model {
	return &model{viewport: viewport.New(0, 0), internal: initial, help: help.New()}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, common.GlobalKeys.Exit):
			return m, tea.Quit

		case key.Matches(msg, common.GlobalKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
			m.updateViewportDim()
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.help.Width = msg.Width
		m.updateViewportDim()
	}

	var vcmd tea.Cmd
	m.internal, cmd = m.internal.Update(msg)
	m.viewport.SetContent(m.internal.View())
	m.viewport, vcmd = m.viewport.Update(msg)
	return m, tea.Batch(vcmd, cmd)
}

func (m *model) View() string {
	return baseStyle.Render(m.viewport.View() + "\n" + m.help.View(common.GlobalKeys))
}
