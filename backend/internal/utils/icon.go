package utils

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed assets/URMC.ico
var APPICON []byte

func GetIconPath() string {
	tempDir := os.TempDir()
	iconPath := filepath.Join(tempDir, "urmc-hub-icon.png")

	if _, err := os.Stat(iconPath); os.IsNotExist(err) {
		os.WriteFile(iconPath, APPICON, 0644)
	}

	return iconPath
}
