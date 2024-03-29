package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SampleModel struct {
	Text string

	VP         viewport.Model
	ViewWidth  int
	ViewHeight int
}

func (m SampleModel) Init() tea.Cmd {
	return nil
}

func (bm SampleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		bm.ViewWidth, bm.ViewHeight = msg.Width, msg.Height
		bm.VP.Width = bm.ViewWidth
		bm.VP.Height = bm.ViewHeight
	}
	bm.VP, cmd = bm.VP.Update(msg)
	return bm, cmd
}
func (m SampleModel) View() string {
	m.VP.SetContent(m.Text)
	return m.VP.View()
}

func NewSampleModel() SampleModel {
	sampleModel := SampleModel{
		Text: "Sample",
		VP:   viewport.New(0, 0),
	}
	sampleStyle := sampleModel.VP.
		Style.
		Border(lipgloss.NormalBorder(), true, true, true)
	sampleModel.VP.Style = sampleStyle
	return sampleModel
}
