package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/jczornik/fujira/views"
	"github.com/jczornik/fujira/views/mainmenu"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "")
		if err != nil {
			panic(err)
		}

		log.Println("Logging is set up")
		defer f.Close()
	}

	log.Println("Starting program")
	if _, err := tea.NewProgram(views.InitialModel(mainmenu.InitialModel()), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
