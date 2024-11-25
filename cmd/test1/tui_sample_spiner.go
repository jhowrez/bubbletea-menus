package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type spinnerModel struct {
	sp spinner.Model
}

func newSpinnerModel() spinnerModel {
	return spinnerModel{
		sp: spinner.New(),
	}
}

func (model spinnerModel) Init() tea.Cmd {
	return model.sp.Tick
}

func (model spinnerModel) IsInitOnEnter() bool {
	return true
}

func (model spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case spinner.TickMsg:
		model.sp, cmd = model.sp.Update(msg)
	default:
		_ = msg
	}

	return model, cmd
}

func (model spinnerModel) View() string {
	return fmt.Sprintf("Spinner  %s", model.sp.View())
}
