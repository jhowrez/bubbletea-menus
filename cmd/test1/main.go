package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	menus "github.com/jhowrez/bubbletea-menus"
)

func main() {
	handledModel := NewHandlerModel()
	sampleModel := NewSampleModel()

	subMenu1 := menus.NewBubbleMenu(
		"Sub Menu 1",
		menus.NewBubbleMenuEntry(
			"Tool Other",
			handledModel,
			"Self handling model",
		),
	)

	subMenu2 := menus.NewBubbleMenu(
		"Sub Menu 2",
		menus.NewBubbleMenuEntry("Tool C", sampleModel, "Tooling C"),
		menus.NewBubbleMenuEntry("Tool D", sampleModel, "Tooling D"),
	)

	subMenu4 := menus.NewBubbleMenu(
		"Sub Sub Menu 1",
		menus.NewBubbleMenuEntry("Self handling model", handledModel, "Is slow for some reason"),
	)
	subMenu3 := menus.NewBubbleMenu(
		"Sub Menu 3",
		menus.NewBubbleMenuEntry("Sub Menu 4", subMenu4, "Nested self handling model"),
		menus.NewBubbleMenuEntry("Text Input", newSlowTextInput(), "Is slow for some reason"),
	)

	menusList := []menus.BubbleMenuEntry{
		menus.NewBubbleMenuEntry("Sub Menu 1", subMenu1, "Tools for A and B"),
		menus.NewBubbleMenuEntry("Sub Menu 2", subMenu2, "Tools for C and D"),
		menus.NewBubbleMenuEntry("Sub Menu 3", subMenu3, "Text Input Model"),
	}

	for range 0 {
		menusList = append(menusList, menusList[0])

	}

	mainMenu := menus.NewBubbleMenu(
		"Main Menu",
		menusList...,
	)

	//	mainMenu.SelectActiveView(2)

	p := tea.NewProgram(ModelResetOnResize{Content: mainMenu},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
