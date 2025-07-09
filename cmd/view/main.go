package main

import (
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kieranbarty/mcp-tui/internal/app/view"
)

func main() {
	// list of tabs to display

	// setup default logger

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// slog.SetDefault(logger)

	slog.SetLogLoggerLevel(slog.LevelDebug)
	// TODO: get this from a flag being passed in but for now hardcode

	debug := true
	if debug {
		f, err := tea.LogToFile("debug.log", "")
		if err != nil {
			slog.Default().Error("error starting logger", slog.Any("error", err))
			os.Exit(1)
		}

		defer f.Close()
	}

	m := view.NewModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		slog.Default().Error("error running cli", slog.Any("error", err))
		os.Exit(1)
	}
}
