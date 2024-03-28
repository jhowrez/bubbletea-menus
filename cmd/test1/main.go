package main

import (
	"fmt"
	"os"

	menus "github.com/beowulf20/bubbletea-menus"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(2, 2)
)

type SampleM struct {
	Text string

	VP         viewport.Model
	ViewWidth  int
	ViewHeight int
}

func (m SampleM) Init() tea.Cmd {
	return nil
}

func (bm SampleM) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		bm.ViewWidth, bm.ViewHeight = msg.Width, msg.Height
		bm.Text = fmt.Sprintf("%dx%d", bm.ViewWidth, bm.ViewHeight)
		bm.VP.Width = bm.ViewWidth
		bm.VP.Height = bm.ViewHeight
	case tea.KeyMsg:
		bm.Text = msg.String()
	}
	bm.VP, cmd = bm.VP.Update(msg)
	return bm, cmd
}
func (m SampleM) View() string {
	m.VP.SetContent(m.Text)
	return m.VP.View()
}

func main() {

	sampleModel := SampleM{
		Text: "Device A Description",
		VP:   viewport.New(0, 0),
	}
	sampleStyle := sampleModel.VP.
		Style.
		Border(lipgloss.NormalBorder(), true, true, true)
	sampleModel.VP.Style = sampleStyle

	aquaMenu := menus.NewBubbleMenu(
		"Aqua Menu",
		menus.NewBubbleMenuEntry("Device A", sampleModel),
	)
	mainMenu := menus.NewBubbleMenu(
		"Main Menu",
		menus.NewBubbleMenuEntry("Aqua Menu", aquaMenu),
	)

	p := tea.NewProgram(ModelResetOnResize{Content: mainMenu},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
