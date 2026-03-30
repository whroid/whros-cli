package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitConfig() error {
	dataDir := filepath.Join(os.Getenv("HOME"), ".whros")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	viper.SetDefault("data_dir", dataDir)
	viper.SetDefault("default_priority", "medium")
	viper.SetDefault("date_format", "2006-01-02 15:04")

	configPath := filepath.Join(dataDir, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		viper.SetConfigFile(configPath)
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}
		fmt.Printf("Config created at: %s\n", configPath)
	}
	return nil
}

func LoadConfig() {
	dataDir := filepath.Join(os.Getenv("HOME"), ".whros")
	viper.SetConfigFile(filepath.Join(dataDir, "config.yaml"))
	viper.ReadInConfig()
}