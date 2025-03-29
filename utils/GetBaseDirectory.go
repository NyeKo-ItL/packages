package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetBaseDirectory determines the base directory based on the execution context
func GetBaseDirectory() (dir string, state string) {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error finding executable path: %v", err)
	}
	execDir := filepath.Dir(execPath)

	// Check if the executable path is within the Go build cache by calculating a relative path
	if err == nil && strings.Contains(execDir, "var") || err == nil && strings.Contains(execDir, "Temp") || err == nil && strings.Contains(execDir, "tmp") {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current working directory: %v", err)
		}
		return cwd, "Debug"
	}

	// Otherwise, return the directory of the executable
	return execDir, "Prod"
}
