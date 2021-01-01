package run

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ImageConfig struct {
	OciVersion    string        `json:"ociVersion"`
	ProcessConfig ProcessConfig `json:"process"`
	Hostname      string        `json:"hostname"`
}

type ProcessConfig struct {
	Terminal bool           `json:"terminal"`
	User     map[string]int `json:"user"`
	Args     []string       `json:"args"`
	Env      []string       `json:"env"`
	Cwd      string         `json:"cwd"`
}

func GetImageConfig(imageDir string) ImageConfig {
	configFile, err := os.Open(imageDir + "/config.json")
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(configFile)
	var config ImageConfig
	json.Unmarshal(bytes, &config)
	return config
}
