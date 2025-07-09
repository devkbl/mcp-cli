package view

// client contains a given client to put in the list
type client struct {
	name        string
	description string
	url         string
}

// implement filter value to satisfy listitem interface
func (c client) FilterValue() string {
	return c.name
}

// implement title func to satisfy DefaultItem interface
func (c client) Title() string {
	return c.name
}

// implement description func to satisfy DefaultItem interface
func (c client) Description() string {
	return c.description
}
