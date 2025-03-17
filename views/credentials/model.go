package credentials

import (
	"fmt"
	"strings"

	"github.com/jczornik/fujira/config"
	"github.com/jczornik/fujira/views/common"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle         = common.FocusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = common.BlurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	fSubmitButton = common.FocusedStyle.Render("[ Submit ]")
	bSubmitButton = fmt.Sprintf("[ %s ]", common.BlurredStyle.Render("Submit"))

	fCancelButton = common.FocusedStyle.Render("[ Cancel ]")
	bCancelButton = fmt.Sprintf("[ %s ]", common.BlurredStyle.Render("Cancel"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func InitialModel() common.Widget {
	m := model{
		inputs: make([]textinput.Model, 2),
	}

	c := config.GetConfig()

	email := ""
	token := ""
	auth, err := c.GetBasicAuth()

	if err == nil {
		email = auth.GetEmail()
		token = auth.GetToken()
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 500

		switch i {
		case 0:
			if email == "" {
				t.Placeholder = "email"
			} else {
				t.SetValue(email)
			}
			t.Focus()
			t.PromptStyle = common.FocusedStyle
			t.TextStyle = common.FocusedStyle
		case 1:
			if token == "" {
				t.Placeholder = "token"
			} else {
				t.SetValue(token)
			}
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m model) Update(msg tea.Msg) (common.Widget, tea.Cmd) {
	handled := false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, common.GlobalKeys.Confirm):
			if m.focusIndex == len(m.inputs) {
				c := config.GetConfig()

				c.SetBasicAuth(m.inputs[0].Value(), m.inputs[1].Value())
				err := c.Save()
				if err != nil {
					panic(err)
				}

				// TODO: run infinit loader and check
				return checkCredentialsView(m)
			} else if m.focusIndex == len(m.inputs)+1 {
				return common.GetPrevView()
			}

		case key.Matches(msg, common.GlobalKeys.Next):
			m.focusIndex = (m.focusIndex + 1) % (len(m.inputs) + 2)
			handled = true

		case key.Matches(msg, common.GlobalKeys.Prev):
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) + 1
			}
			handled = true
		}
	}

	if handled == true {
		cmds := make([]tea.Cmd, len(m.inputs))
		for i := 0; i <= len(m.inputs)-1; i++ {
			if i == m.focusIndex {
				cmds[i] = m.inputs[i].Focus()
				m.inputs[i].PromptStyle = common.FocusedStyle
				m.inputs[i].TextStyle = common.FocusedStyle
				continue
			}

			m.inputs[i].Blur()
			m.inputs[i].PromptStyle = noStyle
			m.inputs[i].TextStyle = noStyle
		}

		return m, tea.Batch(cmds...)
	} else {
		cmd := m.updateInputs(msg)
		return m, cmd
	}
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	confirm := &bSubmitButton
	if m.focusIndex == len(m.inputs) {
		confirm = &fSubmitButton
	}

	cancel := &bCancelButton
	if m.focusIndex == len(m.inputs)+1 {
		cancel = &fCancelButton
	}
	fmt.Fprintf(&b, "\n\n%s\t\t%s\n\n", *confirm, *cancel)

	return b.String()
}

func (m model) Help() help.KeyMap {
	return nil
}
