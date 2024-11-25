package menus

import (
	tea "github.com/charmbracelet/bubbletea"
)

type BubbleMenuEntry struct {
	Name          string
	Content       tea.Model
	description   string
	menuIndex     int
	selfHandled   bool
	isInitOnEnter bool
}

type BubbleMenuEntryOption = func(*BubbleMenuEntry)

func (bm BubbleMenuEntry) Title() string       { return bm.Name }
func (bm BubbleMenuEntry) Description() string { return bm.description }
func (bm BubbleMenuEntry) FilterValue() string { return bm.Name }
func (bm *BubbleMenuEntry) SetSelfHandled(sh bool) {
	bm.selfHandled = sh
}

func (bm BubbleMenuEntry) IsSelfHandled() bool {
	return bm.selfHandled
}

func (bm BubbleMenuEntry) IsInitOnEnter() bool {
	return bm.isInitOnEnter
}

func NewBubbleMenuEntry(name string, content tea.Model, description string, options ...BubbleMenuEntryOption) BubbleMenuEntry {
	bm := BubbleMenuEntry{Name: name, Content: content, description: description}
	for _, opt := range options {
		opt(&bm)
	}
	return bm
}

func WithSelfHandled() BubbleMenuEntryOption {
	return func(bm *BubbleMenuEntry) {
		bm.selfHandled = true
	}
}

func WithInitOnEnter() BubbleMenuEntryOption {
	return func(bm *BubbleMenuEntry) {
		bm.isInitOnEnter = true
	}
}

type SelfHandledEntry interface {
	IsSelfHandled() bool
}

type InitOnEnterEntry interface {
	IsInitOnEnter() bool
}
