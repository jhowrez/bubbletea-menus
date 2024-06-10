package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type slowTextInput struct {
	input textinput.Model
}

func newSlowTextInput() slowTextInput {
	model := slowTextInput{}
	model.input = textinput.New()
	model.input.Focus()
	return model
}

func (model slowTextInput) Init() tea.Cmd {
	return nil
}

func (model slowTextInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	model.input, cmd = model.input.Update(msg)
	return model, cmd
}

func (model slowTextInput) View() string {
	return model.input.View()
}
