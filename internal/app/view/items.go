package view

import "github.com/mark3labs/mcp-go/client"

// client contains a given client to put in the list
type clientItem struct {
	name        string
	description string
	url         string
	client      *client.Client
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
