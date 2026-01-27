package form

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erichberger/sw-helper/ui/component"
)

type Form struct {
	Title      string
	Components []component.Component
	focus      int
}

func (f *Form) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (f *Form) Update(msg tea.Msg) (*Form, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			f.focus++
			if f.focus >= len(f.Components) {
				f.focus = 0
			}
		case "shift+tab":
			f.focus--
			if f.focus < 0 {
				f.focus = len(f.Components) - 1
			}
		}

	}
	var cmd tea.Cmd
	f.Components[f.focus], cmd = f.Components[f.focus].Update(msg)
	return f, cmd
}

func (f *Form) View() string {
	s := strings.Builder{}
	s.WriteString(f.Title)
	s.WriteString("\n\n")

	for i := 0; i < len(f.Components); i++ {
		s.WriteString(f.Components[i].View(f.focus == i))
		s.WriteString("\n")
	}

	s.WriteString("\nHint: press tab to cycle through fields, arrows/hjkl to switch between options")
	return s.String()

}
func (f *Form) Values() map[string]string {
	values := make(map[string]string)

	for _, c := range f.Components {
		if field, ok := c.(component.Field); ok {
			values[field.Key()] = field.Value()
		}
	}

	return values
}
