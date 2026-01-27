package fields

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/erichberger/sw-helper/ui/component"
	"github.com/erichberger/sw-helper/ui/style"
)

// Implements Field interface for rendering in use with other forms

type TextField struct {
	Input textinput.Model
	Label string
	key   string
}

func NewTextField(label string, key string) *TextField {
	input := textinput.New()
	input.Prompt = ":: "
	return &TextField{
		Label: label,
		Input: input,
		key:   key,
	}
}

func (f *TextField) Update(msg tea.Msg) (component.Component, tea.Cmd) {
	var cmd tea.Cmd
	f.Input, cmd = f.Input.Update(msg)
	return f, cmd
}

func (f *TextField) View(active bool) string {
	if active {
		f.Input.Focus()
	} else {
		f.Input.Blur()
	}
	out := strings.Builder{}
	if active {
		out.WriteString(style.FormActiveLabel.Render(f.Label))
	} else {
		out.WriteString(f.Label)
	}
	out.WriteString("\n")
	out.WriteString(f.Input.View())
	out.WriteString("\n")
	return out.String()
}

func (f *TextField) Value() string {
	return f.Input.Value()
}

func (f *TextField) Key() string {
	return f.key
}
