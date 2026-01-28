package style

import "github.com/charmbracelet/lipgloss"

var TitanGreen = lipgloss.Color("#85d038")
var Gray = lipgloss.Color("#939393")
var FormActiveLabel = lipgloss.NewStyle().Bold(true)
var FormActiveColor = lipgloss.NewStyle().Foreground(TitanGreen).Bold(true)
var FormInactiveInput = lipgloss.NewStyle().Foreground(Gray)
var InputMarker = ":: "
var Red = lipgloss.Color("#C21807")
var ErrorTitle = lipgloss.NewStyle().Bold(true).Foreground(Red)
var SuccessTitle = lipgloss.NewStyle().Bold(true).Foreground(TitanGreen)
