package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(2, 2)
)

type SampleM struct {
	Text string

	ViewWidth  int
	ViewHeight int
}

func (m SampleM) Init() tea.Cmd {
	return nil
}

func (bm SampleM) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		bm.ViewWidth, bm.ViewHeight = msg.Width, msg.Height
		bm.Text = fmt.Sprintf("%dx%d", bm.ViewWidth, bm.ViewHeight)
	case tea.KeyMsg:
		switch msg.String() {
		case "right":

		}
	}
	return bm, nil
}
func (m SampleM) View() string { return m.Text }

func main() {

	sampleModel := SampleM{Text: "Device A Description"}

	aquaMenu := NewBubbleMenu("Aqua Menu", NewBubbleMenuEntry("Device A", sampleModel))
	mainMenu := NewBubbleMenu("Main Menu", NewBubbleMenuEntry("Aqua Menu", aquaMenu))

	p := tea.NewProgram(ModelResetOnResize{Content: mainMenu},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
