package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(2, 2)
)

func defaultListItemStyles() list.DefaultItemStyles {
	s := list.NewDefaultItemStyles()
	s.SelectedTitle.Foreground(lipgloss.Color("196"))
	s.SelectedDesc.Foreground(lipgloss.Color("124"))
	s.SelectedTitle.BorderForeground(lipgloss.Color("88"))
	s.SelectedDesc.BorderForeground(lipgloss.Color("88"))

	return s
}
