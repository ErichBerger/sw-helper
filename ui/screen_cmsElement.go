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

type cmsElementScreen struct {
	config *app.Config
	form   *form.Form
}

func (m *cmsElementScreen) handleFormSubmit(values map[string]string) tea.Cmd {
	return func() tea.Msg {
		// call the eventual 'handler' for this
		task, err := tasks.NewCmsElementTask(tasks.CmsElementOptions{
			BasePluginName:       values["basePluginName"],
			PascalCmsElementName: values["cmsElementName"],
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
		return successMsg{message: "Successfully made CMS Element. Enjoy!"}
	}

}

func newCmsElementScreen(config *app.Config) tea.Model {

	base := fields.NewTextField("Base Shopware Plugin", "basePluginName")
	base.Input.Placeholder = "ITWholeSalerPro"
	base.Input.Width = 40
	base.Input.CharLimit = 40

	name := fields.NewTextField("CMS Element Name (Pascal Case)", "cmsElementName")
	name.Input.Placeholder = "TitanExampleName"
	name.Input.Width = 40
	name.Input.CharLimit = 40

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
	return cmsElementScreen{
		// Base plugin
		config: config,
		form: &form.Form{
			Title: style.SuccessTitle.Render("Options for creating a CMS Element"),
			Components: []component.Component{
				base,
				name,
				formFooter,
			},
		},
	}
}
func (m cmsElementScreen) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m cmsElementScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m cmsElementScreen) View() string {
	return m.form.View()

}
