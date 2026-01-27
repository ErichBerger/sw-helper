package tasks

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/erichberger/sw-helper/internal/app"
)

// Task is an interface because Steps is a behavior. The steps of a task differ depending on its options
type Task interface {
	Steps(*app.Config) ([]Step, error)
}
type Step interface {
	Run() error
}

type appendTemplateStep struct {
	name     string
	path     string
	template *template.Template
	data     any
}

func (s *appendTemplateStep) Run() error {
	dir := filepath.Dir(s.path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}

	f, err := os.OpenFile(s.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	if err := ensureTrailingNewline(s.path, w); err != nil {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}
	err = s.template.Execute(w, s.data)
	if err != nil {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}
	return nil
}

type ensureLineExistsStep struct {
	name string
	line string
	path string
}

func (s *ensureLineExistsStep) Run() error {
	dir := filepath.Dir(s.path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}
	// checking to see if it exists
	contents, err := os.ReadFile(s.path)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}

	lines := strings.Split(string(contents), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == s.line {
			return nil
		}
	}
	file, err := os.OpenFile(s.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	if err := ensureTrailingNewline(s.path, w); err != nil {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}
	if _, err := w.WriteString(s.line + "\n"); err != nil {
		return fmt.Errorf("error in step %s: \n%s", s.name, err.Error())
	}

	return nil
}

type WriteMode int

const (
	Overwrite WriteMode = iota
	Append
)

func HandleTask(task Task, config *app.Config) error {
	steps, err := task.Steps(config)
	if err != nil {
		return err
	}
	for _, step := range steps {
		err := step.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
