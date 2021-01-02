package oci

import (
	// "github.com/containers/image/copy"

	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/apex/log"
	"github.com/containers/image/copy"
	"github.com/containers/image/signature"
	"github.com/containers/image/transports"
	"github.com/containers/image/types"
)

type knownTransports struct {
	transports map[string]types.ImageTransport
	mu         sync.Mutex
}

var kt *knownTransports

func PullImage(imagesDir string, image string) ([]byte, error) {
	os.Chdir(imagesDir)
	ctx := context.Background()
	policyContext, err := getPolicyContext()
	if err != nil {
		log.Fatal("Failed to get policy context")
	}
	srcRef, err := ParseImageName(image, "docker")
	if err != nil {
		log.Fatal("Invalid image name")
	}
	destRef, err := ParseImageName(image, "oci")
	if err != nil {
		log.Fatal("Failed to set destination name")
	}
	return copy.Image(ctx, policyContext, destRef, srcRef, &copy.Options{})
}

// ParseImageName converts a URL-like image name to a types.ImageReference.
func ParseImageName(imgName string, transportType string) (types.ImageReference, error) {
	transport := transports.Get(transportType)
	if transport == nil {
		log.Fatal("Failed to get image transport type.")
	}
	imgNameSplit := strings.Split(imgName, "/")
	imgName = imgNameSplit[len(imgNameSplit)-1]
	if transportType == "docker" {
		imgName = fmt.Sprintf("//%s", imgName)
	}
	return transport.ParseReference(imgName)
}

func getPolicyContext() (*signature.PolicyContext, error) {
	var policy *signature.Policy
	var err error
	policy = &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	if err != nil {
		return nil, err
	}
	return signature.NewPolicyContext(policy)
}
