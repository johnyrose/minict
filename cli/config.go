package cli

import (
	"fmt"
	"os"
)

type AppConfig struct {
	ContainersDir string
	ImagesDir     string
}

func GetAppConfig() AppConfig {
	config := AppConfig{
		ContainersDir: fmt.Sprintf("%s/containers", getenv("MINICT_DIR", "/var/lib/minict")),
		ImagesDir:     fmt.Sprintf("%s/images", getenv("MINICT_DIR", "/var/lib/minict")),
	}
	return config
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
