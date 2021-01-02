package run

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ImageConfig struct {
	OciVersion    string         `json:"ociVersion"`
	ProcessConfig ProcessConfig  `json:"process"`
	Hostname      string         `json:"hostname"`
	MountsConfig  []MountsConfig `json:"mounts"`
}

type ProcessConfig struct {
	Terminal bool           `json:"terminal"`
	User     map[string]int `json:"user"`
	Args     []string       `json:"args"`
	Env      []string       `json:"env"`
	Cwd      string         `json:"cwd"`
}

type MountsConfig struct {
	Destination string   `json:"destination"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	Options     []string `json:"options,omitempty"`
}

func GetImageConfig(imageDir string) ImageConfig {
	configFile, err := os.Open(fmt.Sprintf("%s/config.json", imageDir))
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(configFile)
	var config ImageConfig
	json.Unmarshal(bytes, &config)
	return config
}
