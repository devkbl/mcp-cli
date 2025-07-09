package main

import (
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kieranbarty/mcp-tui/internal/app/view"
)

func main() {
	// list of tabs to display
	clientsTab, err := view.NewTab("Clients", "content for clients tab")
	if err != nil {
		slog.Default().Error("error creating clients tab", slog.Any("error", err))
	}

	newClientTab, err := view.NewTab("New Client", "content for new client tab")
	if err != nil {
		slog.Default().Error("error creating new client tab", slog.Any("error", err))
	}
	tabs := []*view.Tab{clientsTab, newClientTab}

	m := view.Model{Tabs: tabs}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		slog.Default().Error("error running cli", slog.Any("error", err))
		os.Exit(1)
	}
}
