package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/erichberger/sw-helper/internal/app"
	"github.com/erichberger/sw-helper/tasks"
	"github.com/erichberger/sw-helper/ui/component"
	"github.com/erichberger/sw-helper/ui/fields"
	"github.com/erichberger/sw-helper/ui/form"
	"github.com/erichberger/sw-helper/ui/style"
)

type storefrontJsScreen struct {
	config *app.Config
	form   *form.Form
}

type formSubmitMsg struct {
}
type formBackMsg struct {
}

type successMsg struct {
	message string
}

type BackMsg struct {
}

type errorMsg struct {
	err error
}

func (m *storefrontJsScreen) handleFormSubmit(values map[string]string) tea.Cmd {
	return func() tea.Msg {
		// call the eventual 'handler' for this
		task, err := tasks.NewStorefrontJsTask(tasks.StorefrontJsOptions{
			BasePluginName:   values["basePluginName"],
			PascalPluginName: values["pluginName"],
			ImportMode:       values["importMode"],
			Hook:             values["hook"],
		})

		if err != nil {
			return errorMsg{err: err}
		}

		if err := task.Validate(m.config); err != nil {
			return errorMsg{err: err}
		}
		if err := tasks.HandleTask(task, m.config); err != nil {
			return errorMsg{err: err}
		}
		return successMsg{message: "Successfully made plugin. Enjoy!"}
	}

}

func newStorefrontJsScreen(config *app.Config) tea.Model {
	// Base plugin
	base := fields.NewTextField("Base Shopware Plugin", "basePluginName")
	base.Input.Placeholder = "ITWholeSalerPro"
	base.Input.Width = 40
	base.Input.CharLimit = 40
	// Name field
	name := fields.NewTextField("JS Plugin Name (Pascal Case)", "pluginName")

	name.Input.Placeholder = "TitanExampleName"
	name.Input.Width = 40
	name.Input.CharLimit = 40

	// Hook field
	hook := fields.NewTextField("JS Hook in HTML (css selector, not required)", "hook")

	hook.Input.Placeholder = "[data-titan-attribute]"
	hook.Input.Width = 40
	hook.Input.CharLimit = 40

	mode := fields.NewRadioField(
		"How to register plugin in main.js: ",
		[]fields.RadioOption{
			{
				Label: "Synchronously",
				Value: tasks.STOREFRONT_JS_IMPORT_MODE_SYNC,
			},
			{
				Label: "Asynchronously",
				Value: tasks.STOREFRONT_JS_IMPORT_MODE_ASYNC,
			},
		},
		"importMode",
	)
	formFooter := component.NewFormFooter(
		[]component.Button{
			{
				Label: "Back",
				Msg:   formBackMsg{},
			},
			{
				Label: "Submit",
				Msg:   formSubmitMsg{},
			},
		},
	)

	return storefrontJsScreen{
		config: config,
		form: &form.Form{
			Title: style.SuccessTitle.Render("Options for creating a Javascript Storefront Plugin"),
			Components: []component.Component{
				base,
				name,
				hook,
				mode,
				formFooter,
			},
		},
	}
}
func (m storefrontJsScreen) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m storefrontJsScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	case formBackMsg:
		return m, emit(navigationMsg{
			target: home,
		})
	case formSubmitMsg:
		return m, m.handleFormSubmit(m.form.Values())
	}
	var cmd tea.Cmd = nil
	m.form, cmd = m.form.Update(msg)
	return m, cmd
}

func (m storefrontJsScreen) View() string {
	return m.form.View()

}
