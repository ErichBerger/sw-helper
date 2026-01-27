package fields

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erichberger/sw-helper/ui/component"
)

type RadioField struct {
	Label            string
	Modes            []RadioOption
	SelectedSymbol   string
	UnselectedSymbol string
	cursor           int
	key              string
}

type RadioOption struct {
	Label string
	Value string
}

func NewRadioField(label string, modes []RadioOption, key string) *RadioField {
	return &RadioField{
		Label:            label,
		Modes:            modes,
		SelectedSymbol:   "●",
		UnselectedSymbol: " ",
		cursor:           0,
		key:              key,
	}
}

func (f *RadioField) Update(msg tea.Msg) (component.Component, tea.Cmd) {

	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return f, nil
	}
	switch key.String() {
	case "left", "h", "up", "k":
		if f.cursor > 0 {
			f.cursor--
		}
	case "right", "l", "down", "j":
		if f.cursor < len(f.Modes)-1 {
			f.cursor++
		}
	}

	return f, nil
}

func (f *RadioField) View(active bool) string {
	output := strings.Builder{}
	output.WriteString(f.Label + "\n")
	output.WriteString(":: ")

	marker := "x"
	if active {
		marker = "●"
	}
	for i, mode := range f.Modes {
		output.WriteString("[")
		if f.cursor == i {
			output.WriteString(marker)
		} else {
			output.WriteString(" ")
		}
		output.WriteString("] ")
		output.WriteString(mode.Label)
		if i != len(f.Modes)-1 {
			output.WriteString("\t")
		}

	}
	output.WriteString("\n\n")
	return output.String()
}

func (f *RadioField) Value() string {
	return f.Modes[f.cursor].Value
}

func (f *RadioField) Key() string {
	return f.key
}
