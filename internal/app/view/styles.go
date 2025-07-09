package view

import "github.com/charmbracelet/lipgloss"

var (
	// general column stylez
	columnStyle = lipgloss.NewStyle().Padding(1, 2)

	// stylez specific to the focused column/list
	focusedColumnStyle = lipgloss.NewStyle().
				Padding(1, 2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62"))

	// stylez specific to our help text
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
