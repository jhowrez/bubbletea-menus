package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type spinnerModel struct {
	sp       spinner.Model
	refInput string
}

func newSpinnerModel() spinnerModel {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	m := spinnerModel{
		sp:       spinner.New(),
		refInput: "NOT STARTED",
	}
	return m
}

func (model spinnerModel) Init() tea.Cmd {
	return tea.Batch(
		model.sp.Tick,
		initCmd(),
	)
}

func (model spinnerModel) IsInitOnEnter() bool {
	return true
}

func (model spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case initMsg:
		model.refInput = "STARTED"
	case spinner.TickMsg:
		model.sp, cmd = model.sp.Update(msg)
		cmds = append(cmds, cmd)
	default:
		_ = msg
	}

	return model, tea.Batch(cmds...)
}

func (model spinnerModel) View() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("Spinner  %s\n", model.sp.View()))
	str.WriteString(fmt.Sprintf("%s\n", model.refInput))
	str.WriteString("\n")
	return str.String()
}

type initMsg struct{}

func initCmd() tea.Cmd {
	return func() tea.Msg {
		return initMsg{}
	}
}
