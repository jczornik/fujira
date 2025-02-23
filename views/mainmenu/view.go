package mainmenu

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jczornik/fujira/views"

	"github.com/jczornik/fujira/views/credentials"
)

var (
	options = map[int]string{
		0: "Auth settings",
		1: "My tasks",
		2: "Exit",
	}

	actions = map[int]func() tea.Model{
		0: credentials.InitialModel,
	}
)

type model struct {
	focusIndex int
}

func InitialModel() model {
	return model{focusIndex: 0}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		inputType := views.MapKeyToType(msg.String())
		switch inputType {
		case views.NextInput:
			m.focusIndex = (m.focusIndex + 1) % len(options)
			return m, nil

		case views.PrevInput:
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(options) - 1
			}

			return m, nil

		case views.ConfirmInput:
			if m.focusIndex == len(options)-1 {
				return m, tea.Quit
			}

			views.AddViewToHistory(m)
			new := actions[m.focusIndex]()
			return new, new.Init()
		}
	}

	cmd := views.HandleGlobal(msg)

	return m, cmd
}

func (m model) View() string {
	var b strings.Builder
	for key := 0; key < len(options); key++ {
		value := options[key]
		if m.focusIndex == key {
			b.WriteString(views.FocusedStyle.Render(fmt.Sprintf("[ %s ]", value)) + "\n")
		} else {
			b.WriteString(fmt.Sprintf("[ %s ]\n", views.BlurredStyle.Render(value)))
		}
	}

	return b.String()
}
