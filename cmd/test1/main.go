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
		menus.NewBubbleMenuEntry("Tool A", sampleModel, "Tooling for Subsystem A"),
		menus.NewBubbleMenuEntry("Tool B", sampleModel, "Tooling for Subsystem B"),
	)

	subMenu2 := menus.NewBubbleMenu(
		"Sub Menu 2",
		menus.NewBubbleMenuEntry("Tool C", sampleModel, "Tooling C"),
		menus.NewBubbleMenuEntry("Tool D", sampleModel, "Tooling D"),
	)

	subMenu3 := menus.NewBubbleMenu(
		"Sub Menu 3",
		menus.NewBubbleMenuEntry("Text Input", menus.NewBubbleMenu(
			"Sub Sub Menu 1",
			menus.NewBubbleMenuEntry("Text Input", newSlowTextInput(), "Is slow for some reason"),
		), "Is slow for some reason"))

	mainMenu := menus.NewBubbleMenu(
		"Main Menu",
		menus.NewBubbleMenuEntry("Sub Menu 1", subMenu1, "Tools for A and B"),
		menus.NewBubbleMenuEntry("Sub Menu 2", subMenu2, "Tools for C and D"),
		menus.NewBubbleMenuEntry("Sub Menu 3", subMenu3, "Text Input Model"),
	)
	mainMenu.SelectActiveView(2)

	p := tea.NewProgram(ModelResetOnResize{Content: mainMenu},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
