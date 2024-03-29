package menus

import "github.com/charmbracelet/bubbles/key"

type MenuKeyMap struct {
	// Enter/Exit menu entry
	EnterView key.Binding
	ExitView  key.Binding

	//
	ForceQuit key.Binding
}

func DefaultMenuKeyMap() MenuKeyMap {
	return MenuKeyMap{
		EnterView: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("↵", "enter view"),
		),
		ExitView: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "exit view"),
		),
		ForceQuit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "force quit"),
		),
	}
}
