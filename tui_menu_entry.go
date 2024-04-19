package menus

import (
	tea "github.com/charmbracelet/bubbletea"
)

type BubbleMenuEntry struct {
	Name        string
	Content     tea.Model
	description string
	menuIndex   int
}

func (bm BubbleMenuEntry) Title() string       { return bm.Name }
func (bm BubbleMenuEntry) Description() string { return bm.description }
func (bm BubbleMenuEntry) FilterValue() string { return bm.Name }

func NewBubbleMenuEntry(name string, content tea.Model, description string) BubbleMenuEntry {
	return BubbleMenuEntry{Name: name, Content: content, description: description}
}
