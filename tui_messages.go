package menus

import tea "github.com/charmbracelet/bubbletea"

type BubbleGoBackMsg struct{}

func BubbleGoBack() tea.Cmd {
	return func() tea.Msg {
		return BubbleGoBackMsg{}
	}
}
