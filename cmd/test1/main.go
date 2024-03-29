package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	menus "github.com/jhowrez/bubbletea-menus"
)

func main() {
	sampleModel := NewSampleModel()

	subMenu1 := menus.NewBubbleMenu(
		"Sub Menu 1",
		menus.NewBubbleMenuEntry("Tool A", sampleModel),
		menus.NewBubbleMenuEntry("Tool B", sampleModel),
	)

	subMenu2 := menus.NewBubbleMenu(
		"Sub Menu 2",
		menus.NewBubbleMenuEntry("Tool C", sampleModel),
		menus.NewBubbleMenuEntry("Tool D", sampleModel),
	)

	mainMenu := menus.NewBubbleMenu(
		"Main Menu",
		menus.NewBubbleMenuEntry("Sub Menu 1", subMenu1),
		menus.NewBubbleMenuEntry("Sub Menu 2", subMenu2),
	)

	p := tea.NewProgram(ModelResetOnResize{Content: mainMenu},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
