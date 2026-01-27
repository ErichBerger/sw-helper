package ui

import tea "github.com/charmbracelet/bubbletea"

// commands must return the tea.Cmd type: func() tea.Msg
// eg. return func() tea.Msg { foo(); return nil;}

func storefrontJsCmd() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
