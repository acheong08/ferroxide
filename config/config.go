package config

import (
	"os"
	"path/filepath"
)

var configHome string

func SetConfigHome(path string) {
	// Check if path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create path
		if err := os.MkdirAll(path, 0700); err != nil {
			panic(err)
		}
	}
	configHome = path
}

func Path(filename string) (string, error) {
	var err error
	if configHome == "" {
		configHome, err = os.UserConfigDir()
	}
	if err != nil {
		return "", err
	}

	p := filepath.Join(configHome, "hydroxide", filename)

	dirname, _ := filepath.Split(p)
	if err := os.MkdirAll(dirname, 0700); err != nil {
		return "", err
	}

	return p, nil
}
