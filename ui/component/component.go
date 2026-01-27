package component

import tea "github.com/charmbracelet/bubbletea"

type Component interface {
	Update(msg tea.Msg) (Component, tea.Cmd)
	View(active bool) string
}

type Field interface {
	Key() string
	Value() string
}

func Emit(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
