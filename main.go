package main

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erichberger/sw-helper/internal/app"
	"github.com/erichberger/sw-helper/ui"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("swh-config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			panic(fmt.Errorf("failure when parsing configuration\n%w\n", err))
		}
	}
	viper.SetDefault("BaseDirectory", "custom/static-plugins")
	baseDir := viper.GetString("BaseDirectory")
	if baseDir == "" {
		panic(fmt.Errorf("no base directory defined either in default or config file\n"))
	}

	config := &app.Config{
		BaseDir: baseDir,
	}
	p := tea.NewProgram(
		ui.NewApp(config),
		tea.WithAltScreen(),
	)
	_, err = p.Run()
	if err != nil {
		fmt.Fprintln(os.Stdout, "Oh no: %w", err)
		os.Exit(1)
	}

	os.Exit(0)

}
