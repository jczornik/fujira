package issues

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jczornik/fujira/api"
	"github.com/jczornik/fujira/views/common"
)

var (
	cols = []table.Column{
		{Title: "id", Width: 10},
		{Title: "code", Width: 15},
		{Title: "description", Width: 200},
	}

	rows = []table.Row{
		{"123", "JIRA-123", "dasdsadsadsjlk sadlsa jdlksa dlksjd kld salkd alk dalk"},
		{"123", "JIRA-223", "dasdsadsadsjlk sadlsa jdlksa dlksjd kld salkd alk dalk"},
		{"123", "JIRA-323", "dasdsadsadsjlk sadlsa jdlksa dlksjd kld salkd alk dalk"},
		{"123", "JIRA-423", "dasdsadsadsjlk sadlsa jdlksa dlksjd kld salkd alk dalk"},
		{"123", "JIRA-523", "dasdsadsadsjlk sadlsa jdlksa dlksjd kld salkd alk dalk"},
	}
)

type model struct {
	msg   string
	table table.Model
}

type result struct {
	response string
	err      error
}

func getIssues() tea.Msg {
	res, err := api.SearchIssues([]string{"id", "description"}, "assignee = currentUser()")
	if err != nil {
		return result{err: err}
	}

	return result{response: res}
}

func InitialModel() common.Widget {
	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return model{table: t}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (common.Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

	case result:
		if msg.err != nil {
			m.msg = msg.err.Error()
		} else {
			m.msg = msg.response
		}

		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	// width, height := views.GetDim()
	// m.table.SetWidth(width)
	// m.table.SetHeight(height)
	return m.table.View()
}

func (m model) Help() help.KeyMap {
	return nil
}
