package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Paths struct {
		ObsidianVault string `mapstructure:"obsidianvault"`
	} `mapstructure:"paths"`
	Device Device `mapstructure:"device"`
}

var C Config

var (
	ErrConfigPathUnavailable = errors.New("Config directory could not be determined")
	ErrConfigCreateDeclined  = errors.New("Config file creation was declined by the user")
	ErrConfigWriteFailed     = errors.New("Failed to write default config file")
)

func ConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrConfigPathUnavailable, err)
	}
	return filepath.Join(configDir, "core", "config.yaml"), nil
}

func EnsureConfigFileExists() (string, error) {
	configPath, err := ConfigPath()
	if err != nil {
		return "", err
	}

	if _, statErr := os.Stat(configPath); statErr == nil {
		return configPath, nil // config file exists
	} else if !errors.Is(statErr, os.ErrNotExist) {
		return "", fmt.Errorf("accessing config file failed: %v", statErr)
	}

	fmt.Printf("Config file not found at %s\nCreate default config file? [y/n]: ", configPath)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.ToLower(strings.TrimSpace(input))
	if input != "y" && input != "yes" {
		return "", ErrConfigCreateDeclined
	}

	defaultConfig := generateDefaultConfig()
	if writeErr := writeConfigToFile(defaultConfig, configPath); writeErr != nil {
		return "", writeErr
	}

	return configPath, nil
}

func generateDefaultConfig() Config {
	var defaultConfig Config
	homeDir, _ := os.UserHomeDir()
	defaultConfig.Paths.ObsidianVault = filepath.Join(homeDir, "documents", "Obsidian Vault")
	defaultConfig.Device = Desktop
	return defaultConfig
}

func writeConfigToFile(config Config, path string) error {
	v := viper.New()
	v.Set("paths.obsidianvault", config.Paths.ObsidianVault)
	v.Set("device", config.Device)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("%w: mkdir: %v", ErrConfigWriteFailed, err)
	}

	v.SetConfigFile(path)
	if err := v.WriteConfigAs(path); err != nil {
		return fmt.Errorf("%w: write: %v", ErrConfigWriteFailed, err)
	}
	return nil
}

func Load() error {
	configPath, err := ConfigPath()
	if err != nil {
		return err
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	defautConfig := generateDefaultConfig()
	v.SetDefault("paths.obsidianvault", defautConfig.Paths.ObsidianVault)
	v.SetDefault("device", defautConfig.Device)

	// overwrite with actual values from file
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("reading config failed: %v", err)
	}
	if err := v.Unmarshal(&C); err != nil {
		return fmt.Errorf("unmarshaling config failed: %v", err)
	}
	return nil
}
