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
}

type ProcessConfig struct {
	Cwd string `json:"cwd"`
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
