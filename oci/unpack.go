package oci

import (
	"fmt"
	"os"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/opencontainers/umoci"
	"github.com/opencontainers/umoci/oci/cas/dir"
	"github.com/opencontainers/umoci/oci/casext"
	"github.com/opencontainers/umoci/oci/layer"
	"github.com/simonz05/util/log"
)

func UnpackImage(imagesDir string, containersDir string, containerName string, imageName string, imageTag string) error {
	os.Chdir(imagesDir)
	engine, err := dir.Open(imageName)
	if err != nil {
		log.Fatal(err)
	}
	engineExt := casext.NewEngine(engine)
	var mapOptions layer.MapOptions
	var meta umoci.Meta
	meta.Version = umoci.MetaVersion
	meta.MapOptions.Rootless = true
	// TODO: Check why adding rootless doesn't work.
	mapOptions = meta.MapOptions
	fullContainerPath := fmt.Sprintf("%s/%s", containersDir, containerName)
	return umoci.Unpack(engineExt, imageTag, fullContainerPath, mapOptions, nil, v1.Descriptor{})
}
