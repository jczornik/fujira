package mainmenu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jczornik/fujira/views/common"

	"github.com/jczornik/fujira/views/credentials"
	"github.com/jczornik/fujira/views/issues"
)

var (
	options = map[int]string{
		0:  "Auth settings",
		1:  "My tasks",
		2:  "Exit",
		3:  "Auth settings",
		4:  "My tasks",
		5:  "Exit",
		6:  "Auth settings",
		7:  "My tasks",
		8:  "Exit",
		9:  "Auth settings",
		10: "My tasks",
		11: "Exit",
		12: "Auth settings",
		13: "My tasks",
		14: "Exit",
		15: "Auth settings",
		16: "My tasks",
		17: "Exit",
		18: "Auth settings",
		19: "My tasks",
		20: "Exit",
		21: "Auth settings",
		22: "My tasks",
		23: "Exit",
		24: "Auth settings",
		25: "My tasks",
		26: "Exit",
	}

	actions = map[int]func() common.Widget{
		0: credentials.InitialModel,
		1: issues.InitialModel,
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

func (m model) Update(msg tea.Msg) (common.Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, common.GlobalKeys.Next):
			m.focusIndex = (m.focusIndex + 1) % len(options)
			return m, nil

		case key.Matches(msg, common.GlobalKeys.Prev):
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(options) - 1
			}

			return m, nil

		case key.Matches(msg, common.GlobalKeys.Confirm):
			if m.focusIndex == len(options)-1 {
				return m, tea.Quit
			}

			common.AddViewToHistory(m)
			new := actions[m.focusIndex]()
			return new, new.Init()
		}
	}

	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	for key := 0; key < len(options); key++ {
		value := options[key]
		if m.focusIndex == key {
			b.WriteString(common.FocusedStyle.Render(fmt.Sprintf("[ %s ]", value)) + "\n")
		} else {
			b.WriteString(fmt.Sprintf("[ %s ]\n", common.BlurredStyle.Render(value)))
		}
	}

	return b.String()
}

func (m model) Help() help.KeyMap {
	return nil
}
