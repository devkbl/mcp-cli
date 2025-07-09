package view

import "errors"

type Tab struct {
	Title   string
	Content string
}

func NewTab(title, content string) (*Tab, error) {
	if title == "" {
		return nil, errors.New("title must not be empty")
	}

	return &Tab{
		Title:   title,
		Content: content,
	}, nil
}
