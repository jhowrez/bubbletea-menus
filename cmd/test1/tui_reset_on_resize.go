package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ModelResetOnResize struct {
	Content tea.Model
}

func (m ModelResetOnResize) Init() tea.Cmd {
	return m.Content.Init()
}
func (m ModelResetOnResize) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if wMsg, ok := msg.(tea.WindowSizeMsg); ok {
		h, v := docStyle.GetFrameSize()
		wMsg.Width -= h
		wMsg.Height -= v
		cmds = append(cmds, tea.ClearScreen)
		msg = wMsg
	}

	m.Content, cmd = m.Content.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
func (m ModelResetOnResize) View() string {
	return docStyle.Render(m.Content.View())
}
