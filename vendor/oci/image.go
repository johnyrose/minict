package oci

import (
	"encoding/json"
	"io/ioutil"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Image struct {
	Layers      []string
	ImageConfig v1.Image
}

type Manifest struct {
	ConfigFile string   `json:"Config,omitempty"`
	Layers     []string `json:"Layers,omitempty"`
}

func GetImageObject(imageDir string) Image {
	manifestFile, _ := ioutil.ReadFile(imageDir + "/manifest.json")
	manifest := Manifest{}
	_ = json.Unmarshall([]byte(manifestFile), &manifest)
	image := v1.Image{}
	_ = json.Unmarshall([]byte(manifest.ConfigFile), &image)
	return Image{
		Layers:      manifest.Layers,
		ImageConfig: image,
	}
}
