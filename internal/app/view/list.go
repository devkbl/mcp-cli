package view

import (
	"context"
	"log/slog"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	mcpClientLib "github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	mcp2 "github.com/mark3labs/mcp-go/mcp"
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

	// set up item delegate styles
	delegate := list.NewDefaultDelegate()

	// initialise client list
	clientList := list.New([]list.Item{}, delegate, 0, 0)
	clientList.Title = "mcp clients"
	clientList.SetItems([]list.Item{
		clientItem{name: "test-client", description: "test, fake mcp client", url: "http://localhost:8080/mcp"},
	})

	// TODO: intialize methods from server ?

	for _, client := range clientList.Items() {
		c := client.(clientItem)
		slog.Debug("client item", slog.Any("client name", c.name), slog.Any("description", c.description))
		ctx := context.Background()
		url := "http://localhost:8080/mcp"
		transport, err := transport.NewStreamableHTTP(url)
		if err != nil {
			slog.Default().ErrorContext(ctx, "error creating StreamableHTTP transport", slog.Any("error", err))
		}

		// create new client based on transport
		mcpClient := mcpClientLib.NewClient(transport)

		// close client connection when function ends
		defer func() {
			if err := mcpClient.Close(); err != nil {
				slog.Default().ErrorContext(ctx, "error closing client", slog.Any("error", err))
			}
		}()

		// initialize client

		initReq := mcp2.InitializeRequest{}

		_, err = mcpClient.Initialize(ctx, initReq)
		if err != nil {
			slog.Default().ErrorContext(ctx, "error initializing client session", slog.Any("error", err))
		}

		tools, err := mcpClient.ListTools(context.TODO(), mcp2.ListToolsRequest{})
		if err != nil {
			slog.Error("error listing MCP tools", slog.Any("error", err))
		}

		for _, tool := range tools.Tools {
			slog.Debug("tool", slog.String("tool_name", tool.Name))
		}
	}

	// initialize methods list
	methodsList := list.New([]list.Item{}, delegate, 0, 0)
	methodsList.Title = "methods"

	// TODO: remove this
	methodsList.SetItems([]list.Item{
		clientItem{name: "test-method", description: "test fake method", url: "no url"},
		clientItem{name: "test-method-2", description: "asdsadasdasd", url: "no url"},
	})
	// initialize responses list
	responsesList := list.New([]list.Item{}, delegate, 0, 0)
	responsesList.Title = "responses"

	m.lists = []list.Model{
		clientList,
		methodsList,
		responsesList,
	}

	return m
}

// TODO: update Next() and Prev() to just check for currently on first/last elem

// update focused column to next
func (m *Model) Next() {
	// if currently on the last element, set next to first element
	if m.focused == responses {
		m.focused = clients
		return
	}

	m.focused++
}

// update focused column to previous
func (m *Model) Prev() {
	if m.focused == clients {
		m.focused = responses
		return
	}

	m.focused--
}

// initialize the model
func (m Model) Init() tea.Cmd {
	return nil
}

// update our model based on cli input
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// set style sizes
		// set all list sizes
		for i := 0; i < len(m.lists); i++ {
			m.lists[i].SetHeight(msg.Height - 10)
			m.lists[i].SetWidth(msg.Width - 4/2)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			m.Prev()
			return m, nil
		case "right", "l":
			m.Next()
			return m, nil
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

	// TODO: has to be a better way of doing this...
	switch m.focused {
	case clients:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focusedColumnStyle.Render(clientsView),
			columnStyle.Render(methodsView),
			columnStyle.Render(responsesView),
		)
	case methods:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(clientsView),
			focusedColumnStyle.Render(methodsView),
			columnStyle.Render(responsesView),
		)
	case responses:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(clientsView),
			columnStyle.Render(methodsView),
			focusedColumnStyle.Render(responsesView),
		)
	}

	return "loading..."
}
