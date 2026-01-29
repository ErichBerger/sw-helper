package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erichberger/sw-helper/internal/app"
	"github.com/erichberger/sw-helper/ui/style"
)

type homeScreen struct {
	config  *app.Config
	choices []choice
	cursor  int
}
type choice struct {
	label string
	key   string
}

var choices = []choice{
	{
		label: "Storefront Javascript Plugin",
		key:   storefrontJs,
	},
	{
		label: "CMS Element",
		key:   cmsElement,
	},
}

func newHomeScreen(config *app.Config) tea.Model {
	return homeScreen{
		// Our to-do list is a grocery list
		config:  config,
		choices: choices,
		cursor:  0,
	}
}
func (m homeScreen) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m homeScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "enter":

			return m, emit(navigationMsg{
				target: m.choices[m.cursor].key,
			})

		case "down", "j", "tab":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		case "up", "k", "shift+tab":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}
	return m, nil
}

func (m homeScreen) View() string {
	s := strings.Builder{}
	s.WriteString(style.SuccessTitle.Render("What would you like to do?") + "\n\n")
	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("[x] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(m.choices[i].label)
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()

}
