package view

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// define custom int type to let us refer to list slice elements by a descriptive name
type listname int

// list names
const (
	clients   listname = iota // 0
	methods                   // 1
	responses                 // 2
)

// initial main model
type Model struct {
	// slice of models - client | methods | responses
	lists []list.Model
	// focused refers to the currently selected list in the above array - TODO: define custom type?
	focused listname
}

func NewModel() Model {
	m := Model{}

	// initialise client list
	clientList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	clientList.Title = "mcp clients"
	clientList.SetItems([]list.Item{
		client{name: "test-client", description: "test, fake mcp client", url: "http://localhost:8080"},
		client{name: "test-client-2", description: "test-2, fake mcp client", url: "http://localhost:8080"},
	})

	// initialize methods list
	methodsList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	methodsList.Title = "methods"

	// TODO: remove this
	methodsList.SetItems([]list.Item{
		client{name: "test-method", description: "test fake method", url: "no url"},
	})
	// initialize responses list
	responsesList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	responsesList.Title = "responses"

	m.lists = []list.Model{
		clientList,
		methodsList,
		responsesList,
	}

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
		// set all list sizes
		for i := 0; i < len(m.lists); i++ {
			m.lists[i].SetHeight(msg.Height - 10)
			m.lists[i].SetWidth(msg.Width - 4/2)
		}
	}
	// pass update to inner list model
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

// update the view based on the current model state
func (m Model) View() string {
	clientsView := m.lists[clients].View()
	methodsView := m.lists[methods].View()
	responsesView := m.lists[responses].View()

	switch m.focused {
	case clients:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focusedColumnStyle.Render(clientsView),
			columnStyle.Render(methodsView),
			columnStyle.Render(responsesView),
		)
	}

	return "loading..."
}
