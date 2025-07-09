package view

import "github.com/charmbracelet/lipgloss"

var (
	// general column stylez
	columnStyle = lipgloss.NewStyle().Padding(1, 2)

	// stylez specific to the focused column/list
	focusedColumnStyle = lipgloss.NewStyle().
				Padding(1, 2).
				Height(10).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("63"))

	// stylez specific to our help text
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
