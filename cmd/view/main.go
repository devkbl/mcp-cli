package main

import (
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kieranbarty/mcp-tui/internal/app/view"
)

func main() {
	// list of tabs to display

	m := view.NewModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		slog.Default().Error("error running cli", slog.Any("error", err))
		os.Exit(1)
	}
}
