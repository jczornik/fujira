package common

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ViewHistory []Widget

var (
	history = ViewHistory{}
)

func (h *ViewHistory) Push(m Widget) {
	*h = append(*h, m)
}

func (h *ViewHistory) Pop() Widget {
	if len(*h) == 0 {
		return nil
	}

	last := len(*h) - 1
	ret := (*h)[last]
	*h = (*h)[:last]

	return ret
}

func GetPrevView() (Widget, tea.Cmd) {
	prev := history.Pop()
	if prev == nil {
		return prev, tea.Quit
	}

	return prev, nil
}

func AddViewToHistory(m Widget) {
	history.Push(m)
}
