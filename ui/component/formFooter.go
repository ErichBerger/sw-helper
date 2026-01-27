package component

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type FormFooter struct {
	cursor  int
	Buttons []Button
}

type Button struct {
	Label string
	Msg   tea.Msg
}

func NewFormFooter(buttons []Button) *FormFooter {
	return &FormFooter{
		Buttons: buttons,
		cursor:  0,
	}
}

func (f *FormFooter) Update(msg tea.Msg) (Component, tea.Cmd) {

	if len(f.Buttons) == 0 {
		return f, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h", "k", "up", "left":
			if f.cursor > 0 {
				f.cursor--
			}
		case "j", "l", "down", "right":
			f.cursor = (f.cursor + 1) % len(f.Buttons)
		case "enter":
			return f, Emit(f.Buttons[f.cursor].Msg)

		}

	}
	return f, nil
}

func (f *FormFooter) View(active bool) string {
	output := strings.Builder{}
	if active {
		output.WriteString("")
	} else {
		output.WriteString("")
	}
	for i, button := range f.Buttons {
		if f.cursor == i {
			output.WriteString("[")
		} else {
			output.WriteString(" ")
		}
		output.WriteString(button.Label)
		if f.cursor == i {
			output.WriteString("]")
		} else {
			output.WriteString(" ")
		}
		if i != len(f.Buttons)-1 {
			output.WriteString("\t\t")
		}
	}

	// output.WriteString("\t\t")
	return output.String()

}
