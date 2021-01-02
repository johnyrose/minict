package cli

import "os"

type AppConfig struct {
	ContainersDir string
	ImagesDir     string
}

func GetAppConfig() AppConfig {
	config := AppConfig{
		ContainersDir: getenv("MINICT_CONTAINERS_DIR", "/var/lib/minict/containers"),
		ImagesDir:     getenv("MINICT_IMAGES_DIR", "/var/lib/minict/images"),
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
