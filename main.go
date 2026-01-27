package main

import (
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
	viper.SetDefault("BaseDirectory", "custom/static-plugins")
	baseDir := viper.GetString("BaseDirectory")
	if baseDir == "" {
		fmt.Fprintf(os.Stderr, "No base directory found")
	}

	if err != nil {
		panic(fmt.Errorf("fatal error with config file: %w", err))
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
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	os.Exit(0)

}
