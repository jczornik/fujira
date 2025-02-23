package credentials

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jczornik/fujira/api"
	"github.com/jczornik/fujira/views"
	"github.com/jczornik/fujira/views/alert"
	"github.com/jczornik/fujira/views/infinitloader"
)

func checkCredentials() error {
	_, err := api.MyPreferences("jira.user.timezone")
	return err
}

func checkCredentialsView(caller tea.Model) (tea.Model, tea.Cmd) {
	logic := func() tea.Msg {
		if err := checkCredentials(); err != nil {
			log.Println("Error while checking credentials: ", err)
			m := alert.InitialModel(fmt.Sprint("Error: ", err), caller, nil)
			return infinitloader.Finished{Model: m}
		}

		prev, _ := views.GetPrevView()
		m := alert.InitialModel("Given credentials are correct!", prev, nil)
		return infinitloader.Finished{Model: m}
	}

	m := infinitloader.InitialModel("Checking credentials", logic)
	return m, m.Init()
}
