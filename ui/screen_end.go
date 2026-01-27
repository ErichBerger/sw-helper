package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type endScreen struct {
	message string
}

func (m endScreen) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func (m endScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	}

	var cmd tea.Cmd = nil
	return m, cmd
}

func (m endScreen) View() string {
	b := strings.Builder{}
	b.WriteString(m.message + "\n\n")
	b.WriteString("Press any key to exit\n")
	return b.String()
}
