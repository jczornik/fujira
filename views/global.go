package views

import tea "github.com/charmbracelet/bubbletea"

var (
	width  int
	height int
)

func GetDim() (int, int) {
	return width, height
}

func HandleGlobal(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch MapKeyToType(msg.String()) {
		case ExitInput:
			return tea.Quit
		}

	case tea.WindowSizeMsg:
		width = msg.Width
		height = msg.Height
	}

	return nil
}
