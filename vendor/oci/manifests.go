package oci

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/apex/log"
)

type ManifestConfig struct {
	schemaVersion int        `json:"schemaVersion"`
	manifests     []Manifest `json:"manifests"`
}

type Manifest struct {
	MediaType   string              `json:"mediaType"`
	Digest      string              `json:"digest"`
	Size        int                 `json:"size"`
	Annotations ManifestAnnotations `json:"annotations"`
}

type ManifestAnnotations struct {
	Tag string `json:"org.opencontainers.image.ref.name"`
}

func ListManifests(imagesDir string) []Manifest {
	var manifests []Manifest
	files, err := ioutil.ReadDir(imagesDir)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, file := range files {
		if file.IsDir() {
			manifestConfig := readManifest(fmt.Sprintf("%s/%s", imagesDir, file.Name()))
			for _, manifest := range manifestConfig.manifests {
				manifests = append(manifests, manifest)
			}
		}
	}
	return manifests
}

func readManifest(imageDir string) ManifestConfig {
	manifestFile, err := os.Open(fmt.Sprintf("%s/index.json", imageDir))
	if err != nil {
		log.Fatal(err.Error())
	}
	bytes, _ := ioutil.ReadAll(manifestFile)
	var manifestConfig ManifestConfig
	json.Unmarshal(bytes, &manifestConfig)
	return manifestConfig
}
