package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erichberger/sw-helper/internal/app"
)

type Config struct {
	BaseDir string
}

type AppModel struct {
	Config *app.Config
	screen tea.Model
}

type screenFactory func(*app.Config) tea.Model

type navigationMsg struct {
	target string
}

const (
	home         = "home"
	storefrontJs = "storefront-js"
)

var config = Config{}

var routes = map[string]screenFactory{
	home:         newHomeScreen,
	storefrontJs: newStorefrontJsScreen,
}

func logError(err string) {
	fmt.Fprintf(os.Stderr, "%s", err)
}
func emit(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg

	}
}

func NewApp(config *app.Config) tea.Model {
	return AppModel{
		Config: config,
		screen: newHomeScreen(config),
	}
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Upate is run after the tea.Cmd resolves, in our case giving us the new target
	switch payload := msg.(type) {
	case navigationMsg:
		if screenFactory, ok := routes[payload.target]; ok {
			m.screen = screenFactory(m.Config)
			return m, nil
		}
	case errorMsg:
		m.screen = endScreen{message: payload.err.Error()}
		return m, nil
	case successMsg:
		m.screen = endScreen{message: payload.message}
		return m, nil
	}

	var cmd tea.Cmd

	if m.screen != nil {
		m.screen, cmd = m.screen.Update(msg)
	}

	return m, cmd

}

func (m AppModel) View() string {
	return m.screen.View()
}

func (m AppModel) Init() tea.Cmd {
	return m.screen.Init()
}
