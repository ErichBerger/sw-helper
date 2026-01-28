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
	f.Input.Blur()
	f.Input.Prompt = style.InputMarker

	out := strings.Builder{}
	label := f.Label

	if active {
		f.Input.Focus()
		label = style.FormActiveLabel.Render(f.Label)
		f.Input.Prompt = style.FormActiveColor.Render(style.InputMarker)
	}
	out.WriteString(label)

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
