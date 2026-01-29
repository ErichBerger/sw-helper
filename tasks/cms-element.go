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
	FILEPATH_ADMIN          = "src/Resources/app/administration/src"
	FILEPATH_ADMIN_ELEMENTS = "src/Resources/app/administration/src/module/sw-cms/elements"
)

type cmsElementTask struct {
	options CmsElementOptions
}

type CmsElementOptions struct {
	BasePluginName   string
	PascalPluginName string
}

type cmsElementTemplateData struct {
	PascalName string
	KebabName  string
	TwigName   string
}

func (t *cmsElementTask) generateTemplateMap() (map[string]*template.Template, error) {
	templates := make(map[string]*template.Template)
	files, err := fs.Glob(cmsElementTemplates, "templates/cms-element/*")
	if err != nil {
		return nil, fmt.Errorf("error parsing templates: %w", err)
	}

	for _, file := range files {
		fileName := filepath.Base(file)
		templ, err := template.ParseFS(cmsElementTemplates, file)
		if err != nil {
			return nil, err
		}
		templates[fileName] = templ
	}

	return templates, nil
}
func NewCmsElementTask(options CmsElementOptions) (*cmsElementTask, error) {
	return &cmsElementTask{
		options: options,
	}, nil
}

func (t *cmsElementTask) Steps(config *app.Config) ([]Step, error) {
	templates, err := t.generateTemplateMap()
	if err != nil {
		return nil, err
	}

	kebabName := pascalCaseToKebab(t.options.PascalPluginName)
	twigName := strings.ReplaceAll(kebabName, "-", "_")
	elementRootDir := filepath.Join(config.BaseDir, t.options.BasePluginName, FILEPATH_ADMIN_ELEMENTS, kebabName)
	data := cmsElementTemplateData{
		PascalName: t.options.PascalPluginName,
		KebabName:  kebabName,
		TwigName:   twigName,
	}

	steps := []Step{}
	steps = append(steps, &appendTemplateStep{
		name:     "Import element in main.js",
		path:     filepath.Join(config.BaseDir, t.options.BasePluginName, FILEPATH_ADMIN, "main.js"),
		template: templates["main-js.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element index.js",
		path:     filepath.Join(elementRootDir, "index.js"),
		template: templates["element-index.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element component index.js",
		path:     filepath.Join(elementRootDir, "component", "index.js"),
		template: templates["component-index.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element component template",
		path:     filepath.Join(elementRootDir, "component", "sw-cms-el-"+kebabName+".html.twig"),
		template: templates["component-twig.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element component style",
		path:     filepath.Join(elementRootDir, "component", "sw-cms-el-"+kebabName+".scss"),
		template: templates["component-style.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element preview index.js",
		path:     filepath.Join(elementRootDir, "preview", "index.js"),
		template: templates["preview-index.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element preview template",
		path:     filepath.Join(elementRootDir, "preview", "sw-cms-el-preview-"+kebabName+".html.twig"),
		template: templates["preview-twig.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element preview style",
		path:     filepath.Join(elementRootDir, "preview", "sw-cms-el-preview-"+kebabName+".scss"),
		template: templates["preview-style.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element config index.js",
		path:     filepath.Join(elementRootDir, "config", "index.js"),
		template: templates["config-index.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element config template",
		path:     filepath.Join(elementRootDir, "config", "sw-cms-el-config-"+kebabName+".html.twig"),
		template: templates["config-twig.tmpl"],
		data:     data,
	})
	steps = append(steps, &appendTemplateStep{
		name:     "Create element storefront view",
		path:     filepath.Join(config.BaseDir, t.options.BasePluginName, "src/Resources/views/storefront/element", "cms-element-"+kebabName+".html.twig"),
		template: templates["storefront-view.tmpl"],
		data:     data,
	})

	return steps, nil
}

func (t *cmsElementTask) Validate(config *app.Config) error {
	if t.options.BasePluginName == "" {
		return fmt.Errorf("base plugin not selected")
	}
	if t.options.PascalPluginName == "" {
		return fmt.Errorf("plugin name not selected")
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
	elementRootDir := filepath.Join(config.BaseDir, t.options.BasePluginName, FILEPATH_ADMIN_ELEMENTS, kebabName)

	info, err = os.Stat(elementRootDir)

	if err == nil {
		if info.IsDir() {
			return fmt.Errorf("%s is already a plugin name", elementRootDir)
		}
		return fmt.Errorf("%s already exists and is not a directory", elementRootDir)
	}
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}
