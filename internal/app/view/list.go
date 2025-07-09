package view

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// initial main model
type Model struct {
	list list.Model
}

func NewModel() Model {
	// initialize main models list

	m := Model{}

	// TODO: set height somehow ?
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)

	// set the title of the initial, main list
	m.list.Title = "MCP Clients"

	// initialize the list of items, using our client items.
	m.list.SetItems([]list.Item{
		client{name: "test-client", description: "test, fake mcp client", url: "http://localhost:8080"},
		client{name: "test-client-2", description: "test-2, fake mcp client", url: "http://localhost:8080"},
	})

	return m
}

// initialize the model
func (m Model) Init() tea.Cmd {
	return nil
}

// update our model based on cli input
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetHeight(msg.Height)
		m.list.SetWidth(msg.Width)
	}
	// pass update to inner list model
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// update the view based on the current model state
func (m Model) View() string {
	return m.list.View()
}
