package menus

import tea "github.com/charmbracelet/bubbletea"

type BubbleMenuBackMsg struct{}

func BubbleMenuBack() tea.Msg {
	return BubbleMenuBackMsg{}
}
