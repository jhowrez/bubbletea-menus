package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type BubbleMenuEntry struct {
	Name    string
	Content tea.Model
}

func (bm BubbleMenuEntry) Title() string       { return bm.Name }
func (bm BubbleMenuEntry) Description() string { return bm.Name }
func (bm BubbleMenuEntry) FilterValue() string { return bm.Name }

type BubbleMenu struct {
	title    string
	desc     string
	children []tea.Model

	menuEntryList     list.Model
	selectedMenuEntry int

	// behaviour properties
	ResetOnBack             bool
	HandleGoBackForChildren bool
}

func (bm BubbleMenu) Title() string       { return bm.title }
func (bm BubbleMenu) Description() string { return bm.desc }
func (bm BubbleMenu) FilterValue() string { return bm.title }

func NewBubbleMenuEntry(name string, content tea.Model) BubbleMenuEntry {
	return BubbleMenuEntry{Name: name, Content: content}
}

func NewBubbleMenu(title string, children ...BubbleMenuEntry) BubbleMenu {
	items := []list.Item{}
	childrenContent := []tea.Model{}
	for _, child := range children {
		items = append(items, child)
		childrenContent = append(childrenContent, child.Content)
	}

	menuEntryList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	menuEntryList.DisableQuitKeybindings()
	menuEntryList.Title = title

	return BubbleMenu{
		title:             title,
		desc:              title,
		children:          childrenContent,
		menuEntryList:     menuEntryList,
		selectedMenuEntry: -1,

		HandleGoBackForChildren: true,
	}
}

func (bm BubbleMenu) Init() tea.Cmd {
	return nil
}

func (bm BubbleMenu) View() string {
	var b strings.Builder

	b.WriteString(
		fmt.Sprintf("\n(%s) WxH = %d,%d\n",
			bm.title,
			bm.menuEntryList.Width(), bm.menuEntryList.Height(),
		))

	if bm.selectedMenuEntry != -1 {
		b.WriteString(bm.children[bm.selectedMenuEntry].View())

	} else if len(bm.children) >= 1 {
		b.WriteString(bm.menuEntryList.View())
	}

	return b.String()
}

func (bm *BubbleMenu) ResetActiveView() {
	if bm.ResetOnBack {
		bm.menuEntryList.ResetSelected()
	}
	bm.selectedMenuEntry = -1
}

func (bm BubbleMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	bm.menuEntryList, cmd = bm.menuEntryList.Update(msg)
	cmds = append(cmds, cmd)

	{
		// update size for all children
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			bm.menuEntryList.SetSize(msg.Width, msg.Height)
			for i := range bm.children {
				var m tea.Model
				m, cmd = bm.children[i].Update(msg)
				cmds = append(cmds, cmd)
				bm.children[i] = m
			}
		}
	}

	{
		// update active view
		if bm.selectedMenuEntry != -1 {
			nm, cmd := bm.children[bm.selectedMenuEntry].Update(msg)
			bm.children[bm.selectedMenuEntry] = nm

			if cmd != nil {
				if batchCmds, ok := cmd().(tea.BatchMsg); ok {
					for _, cmd := range batchCmds {
						if _, ok := cmd().(BubbleMenuBackMsg); ok {
							bm.ResetActiveView()
						}
					}
				} else if _, ok := cmd().(BubbleMenuBackMsg); ok {
					bm.ResetActiveView()
				} else {
					cmds = append(cmds, cmd)
				}
			}

			if _, ok := nm.(BubbleMenu); !ok {
				if bm.HandleGoBackForChildren {
					// handle specific events for children
					switch msg := msg.(type) {
					case tea.KeyMsg:
						switch msg.String() {
						case "left":
							bm.ResetActiveView()
							// return bm, BubbleMenruBack
							// return bm, nil
						}
					}
				}
			}

			return bm, tea.Batch(cmds...)
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return bm, tea.Quit
		case "right":
			if bm.menuEntryList.SelectedItem() != nil {
				bm.selectedMenuEntry = bm.menuEntryList.Index()
			}
			return bm, tea.ClearScreen
		case "left":
			return bm, BubbleMenuBack
		}
	}

	return bm, tea.Batch(cmds...)
}
