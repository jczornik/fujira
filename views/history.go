package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ViewHistory []tea.Model

var (
	history = ViewHistory{}
)

func (h *ViewHistory) Push(m tea.Model) {
	*h = append(*h, m)
}

func (h *ViewHistory) Pop() tea.Model {
	if len(*h) == 0 {
		return nil
	}

	last := len(*h) - 1
	ret := (*h)[last]
	*h = (*h)[:last]

	return ret
}

func GetPrevView() (tea.Model, tea.Cmd) {
	prev := history.Pop()
	if prev == nil {
		return prev, tea.Quit
	}

	return prev, nil
}

func AddViewToHistory(m tea.Model) {
	history.Push(m)
}
