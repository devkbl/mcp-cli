package view

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

// client contains a given client to put in the list
type clientItem struct {
	name        string
	description string
	url         string
	client      *client.Client
	// clientItem contains a list of items (methodItems)
	methodItems []list.Item
}

// implement filter value to satisfy listitem interface
func (c clientItem) FilterValue() string {
	return c.name
}

// implement title func to satisfy DefaultItem interface
func (c clientItem) Title() string {
	return c.name
}

// implement description func to satisfy DefaultItem interface
func (c clientItem) Description() string {
	return c.description
}

// methodItem contains all the information needed to properly describe an MCP 'method' - tools, prompts, resources, etc etc
type methodItem struct {
	name        string
	description string
	feature     string
	request     string
	inputSchema mcp.ToolInputSchema
}

// implement filter value to satisfy listitem interface
func (c methodItem) FilterValue() string {
	return c.name
}

// implement title func to satisfy DefaultItem interface
func (c methodItem) Title() string {
	return c.name
}

// implement description func to satisfy DefaultItem interface
func (c methodItem) Description() string {
	return c.description
}

// responseItem contains all the information needed to properly describe an MCP 'method' - tools, prompts, resources, etc etc
type responseItem struct {
	title       string
	response    string
	contentType string
}

// implement filter value to satisfy listitem interface
func (c responseItem) FilterValue() string {
	return c.response
}

// implement title func to satisfy DefaultItem interface
func (c responseItem) Title() string {
	return c.title
}

// implement description func to satisfy DefaultItem interface
func (c responseItem) Description() string {
	return c.response
}
