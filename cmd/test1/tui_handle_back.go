package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	menus "github.com/jhowrez/bubbletea-menus"
)

type HandlerModel struct {
	Text string

	VP         viewport.Model
	ViewWidth  int
	ViewHeight int

	list list.Model
}

func (m HandlerModel) Init() tea.Cmd {
	return nil
}

func (m HandlerModel) IsSelfHandled() bool {
	return true
}

func (bm HandlerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		bm.ViewWidth, bm.ViewHeight = msg.Width, msg.Height
		bm.VP.Width = bm.ViewWidth
		bm.VP.Height = bm.ViewHeight
	case tea.KeyMsg:
		if bm.list.FilterState() == list.FilterApplied {
			if msg.String() == "esc" {
				return bm, menus.BubbleGoBack()
			}
		}
	}
	bm.VP, cmd = bm.VP.Update(msg)
	cmds = append(cmds, cmd)

	bm.list, cmd = bm.list.Update(msg)
	cmds = append(cmds, cmd)

	return bm, tea.Batch(cmds...)
}
func (m HandlerModel) View() string {
	m.VP.SetContent(m.Text + m.list.View())
	return m.VP.View()
}

func NewHandlerModel() HandlerModel {
	sampleModel := HandlerModel{
		Text: "Write anything to list fileter and press 'ESC' to go back",
		VP:   viewport.New(0, 0),
		list: newSimpleList(),
	}
	sampleStyle := sampleModel.VP.
		Style.
		Border(lipgloss.NormalBorder(), true, true, true)
	sampleModel.VP.Style = sampleStyle
	return sampleModel
}
