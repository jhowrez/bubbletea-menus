package menus

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// for list.Item compatibility

func (bm BubbleMenu) Title() string       { return bm.title }
func (bm BubbleMenu) Description() string { return bm.desc }
func (bm BubbleMenu) FilterValue() string { return bm.title }

type BubbleMenu struct {
	title    string
	desc     string
	children []tea.Model

	menuEntryList     list.Model
	selectedMenuEntry int

	// behaviour properties
	ResetOnBack             bool
	HandleGoBackForChildren bool
	IsFilteringEnabled      bool

	//
	keyMap MenuKeyMap

	//
	isShouldResetView bool

	menuEntries []BubbleMenuEntry
}

func NewBubbleMenu(title string, children ...BubbleMenuEntry) BubbleMenu {
	items := []list.Item{}
	childrenContent := []tea.Model{}
	for i, child := range children {
		child.menuIndex = i
		items = append(items, child)
		childrenContent = append(childrenContent, child.Content)
	}

	menuEntryList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	menuEntryList.Title = title
	menuEntryList.SetShowStatusBar(false)
	menuEntryList.SetShowHelp(true)
	menuEntryList.DisableQuitKeybindings()
	menuEntryList.SetFilteringEnabled(true)

	bm := BubbleMenu{
		title:                   title,
		desc:                    title,
		children:                childrenContent,
		menuEntryList:           menuEntryList,
		selectedMenuEntry:       -1,
		HandleGoBackForChildren: true,
		IsFilteringEnabled:      true,
		menuEntries:             children,
	}

	{
		// help keys
		bm.keyMap = DefaultMenuKeyMap()
		bm.menuEntryList.AdditionalFullHelpKeys = func() []key.Binding {
			return []key.Binding{
				bm.keyMap.EnterView,
				bm.keyMap.ExitView,
				bm.keyMap.ForceQuit,
			}
		}

		bm.menuEntryList.AdditionalShortHelpKeys = func() []key.Binding {
			return []key.Binding{
				bm.keyMap.EnterView,
				bm.keyMap.ExitView,
			}
		}
	}

	return bm
}

func (bm BubbleMenu) Init() tea.Cmd {
	return nil
}

func (bm BubbleMenu) View() string {
	var b strings.Builder

	if bm.selectedMenuEntry != -1 {
		b.WriteString(bm.children[bm.selectedMenuEntry].View())

	} else if len(bm.children) >= 1 {
		b.WriteString(bm.menuEntryList.View())
	}

	return b.String()
}

func (bm *BubbleMenu) ResetActiveView() {
	bm.isShouldResetView = false
	if bm.selectedMenuEntry == -1 {
		return
	}

	if bm.ResetOnBack {
		bm.menuEntryList.ResetFilter()
		bm.menuEntryList.ResetSelected()
	}
	bm.selectedMenuEntry = -1

	if bm.IsFilteringEnabled {
		bm.menuEntryList.SetFilteringEnabled(true)
	}

	bm.menuEntryList.Help.ShowAll = false
}

func (bm *BubbleMenu) SelectActiveView(i int) {
	bm.isShouldResetView = false
	bm.selectedMenuEntry = i

	if bm.IsFilteringEnabled {
		bm.menuEntryList.SetFilteringEnabled(false)
	}
}

func (bm BubbleMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

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
			var childCmd tea.Cmd
			var nm tea.Model
			nm, childCmd = bm.children[bm.selectedMenuEntry].Update(msg)
			bm.children[bm.selectedMenuEntry] = nm

			var bmHandles = bm.HandleGoBackForChildren
			var bmEntryHandles = bm.menuEntries[bm.selectedMenuEntry].IsSelfHandled()
			var viewContentHandles = false

			if h, ok := nm.(SelfHandledEntry); ok {
				viewContentHandles = h.IsSelfHandled()
			}

			if bmEntryHandles || viewContentHandles {
				switch msg.(type) {
				case BubbleGoBackMsg:
					bm.ResetActiveView()
					return bm, tea.Batch(cmds...)
				}
			}

			if childCmd != nil {
				cmds = append(cmds, childCmd)
			}

			switch mt := nm.(type) {
			case BubbleMenu:
				if mt.isShouldResetView {
					mt.isShouldResetView = false
					bm.children[bm.selectedMenuEntry] = mt
					bm.ResetActiveView()
					return bm, tea.Batch(cmds...)
				}
			default:
				// do nothing
				if bmHandles && !viewContentHandles && !bmEntryHandles {
					switch msg := msg.(type) {
					case tea.KeyMsg:
						if key.Matches(msg, bm.keyMap.ExitView) {
							bm.ResetActiveView()
							return bm, tea.Batch(cmds...)
						}
					}
				}
			}

			return bm, tea.Batch(cmds...)
		}
	}

	if bm.menuEntryList.FilterState() != list.Filtering {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if key.Matches(msg, bm.keyMap.EnterView) {
				if bm.menuEntryList.SelectedItem() != nil {
					entryItem := bm.menuEntryList.SelectedItem().(BubbleMenuEntry)
					bm.SelectActiveView(entryItem.menuIndex)
				}
				return bm, tea.ClearScreen
			}

			if key.Matches(msg, bm.keyMap.ExitView) {
				bm.ResetActiveView()
				bm.isShouldResetView = true
				return bm, nil
			}
			if key.Matches(msg, bm.keyMap.ForceQuit) {
				return bm, tea.Quit
			}
		}
	}

	{
		// update list
		bm.menuEntryList, cmd = bm.menuEntryList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return bm, tea.Batch(cmds...)
}

func (bm *BubbleMenu) SetListStyle(listStyles list.Styles, itemStyles list.DefaultItemStyles) {
	bm.menuEntryList.Styles = listStyles
	del := list.NewDefaultDelegate()
	del.Styles = itemStyles
	bm.menuEntryList.SetDelegate(del)
}

func (bm *BubbleMenu) SetHelpStyle(styles help.Styles) {
	bm.menuEntryList.Help.Styles = styles
}
