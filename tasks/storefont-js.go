package tasks

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/erichberger/sw-helper/internal/app"
)

const (
	STOREFRONT_JS_IMPORT_MODE_SYNC  = "sync"
	STOREFRONT_JS_IMPORT_MODE_ASYNC = "async"
	FILEPATH_APP_STOREFRONT         = "src/Resources/app/storefront/src/"
)

type storefrontJsTask struct {
	options      StorefrontJsOptions
	fileBasePath string
}

type StorefrontJsOptions struct {
	BasePluginName   string
	PascalPluginName string
	ImportMode       string
	Hook             string
}

type storefrontJsTemplateData struct {
	PascalName string
	KebabName  string
	Hook       string
}

func (t *storefrontJsTask) generateTemplateMap() (map[string]*template.Template, error) {
	templates := make(map[string]*template.Template)
	files, err := fs.Glob(storefrontJSTemplates, "templates/storefront-js/*")
	if err != nil {
		return nil, fmt.Errorf("error parsing templates: %w", err)
	}

	for _, file := range files {
		fileName := filepath.Base(file)
		templ, err := template.ParseFS(storefrontJSTemplates, file)
		if err != nil {
			return nil, err
		}
		templates[fileName] = templ
	}

	return templates, nil
}

func NewStorefrontJsTask(options StorefrontJsOptions) (*storefrontJsTask, error) {
	return &storefrontJsTask{
		options:      options,
		fileBasePath: "custom/static-plugins",
	}, nil
}

func (t *storefrontJsTask) Steps(config *app.Config) ([]Step, error) {
	templates, err := t.generateTemplateMap()
	if err != nil {
		return nil, err
	}

	kebabName := pascalCaseToKebab(t.options.PascalPluginName)
	data := storefrontJsTemplateData{
		PascalName: t.options.PascalPluginName,
		KebabName:  kebabName,
		Hook:       t.options.Hook,
	}

	steps := []Step{}
	steps = append(steps, &ensureLineExistsStep{
		name: "Ensure window manager is declared in main.js",
		line: "const PluginManager = window.PluginManager;",
		path: filepath.Join(config.BaseDir, t.options.BasePluginName, FILEPATH_APP_STOREFRONT, "main.js"),
	})
	switch t.options.ImportMode {
	case STOREFRONT_JS_IMPORT_MODE_SYNC:
		steps = append(steps, &appendTemplateStep{
			name:     "Add synchronous import to main.js",
			path:     filepath.Join(config.BaseDir, t.options.BasePluginName, FILEPATH_APP_STOREFRONT, "main.js"),
			template: templates["main-js-sync.tmpl"],
			data:     data,
		})
	case STOREFRONT_JS_IMPORT_MODE_ASYNC:
		steps = append(steps, &appendTemplateStep{
			name:     "Add asynchronous import to main.js",
			path:     filepath.Join(config.BaseDir, t.options.BasePluginName, FILEPATH_APP_STOREFRONT, "main.js"),
			template: templates["main-js-async.tmpl"],
			data:     data,
		})
	}

	pluginFileName := strings.Builder{}
	pluginFileName.WriteString(kebabName)
	pluginFileName.WriteString(".plugin.js")
	steps = append(steps, &appendTemplateStep{
		name:     "Write plugin file",
		path:     filepath.Join(config.BaseDir, t.options.BasePluginName, FILEPATH_APP_STOREFRONT, kebabName, pluginFileName.String()),
		template: templates["plugin-js.tmpl"],
		data:     data,
	})
	return steps, nil

}

func (t *storefrontJsTask) Validate(config *app.Config) error {
	if t.options.BasePluginName == "" {
		return fmt.Errorf("base plugin not selected")
	}
	if t.options.PascalPluginName == "" {
		return fmt.Errorf("plugin name not selected")
	}
	if t.options.ImportMode == "" {
		return fmt.Errorf("no input mode selected")
	}
	basePluginPath := filepath.Join(config.BaseDir, t.options.BasePluginName)

	info, err := os.Stat(basePluginPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// directory does not exist
			return fmt.Errorf("%s does not exist", basePluginPath)
		}
		return err // permission or other error
	}

	if !info.IsDir() {
		return fmt.Errorf("%s exists but is not a directory", basePluginPath)
	}

	kebabName := pascalCaseToKebab(t.options.PascalPluginName)
	newJsPluginPath := filepath.Join(basePluginPath, FILEPATH_APP_STOREFRONT, kebabName)

	info, err = os.Stat(newJsPluginPath)

	if err == nil {
		if info.IsDir() {
			return fmt.Errorf("%s is already a plugin name", newJsPluginPath)
		}
		return fmt.Errorf("%s already exists and is not a directory", newJsPluginPath)
	}
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}
