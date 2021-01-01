package oci

import (
	// "github.com/containers/image/copy"

	"context"
	"os"
	"strings"
	"sync"

	"github.com/apex/log"
	"github.com/containers/image/copy"
	"github.com/containers/image/signature"
	"github.com/containers/image/transports"
	"github.com/containers/image/types"
	"github.com/pkg/errors"
)

type knownTransports struct {
	transports map[string]types.ImageTransport
	mu         sync.Mutex
}

var kt *knownTransports

func PullImage(imagesDir string, imageName string, imageTag string) ([]byte, error) {
	os.Chdir(imagesDir)
	ctx := context.Background()
	policyContext, err := getPolicyContext()
	if err != nil {
		log.Fatal("Failed to get policy context")
	}
	srcRef, err := ParseImageName(imageName + ":" + imageTag)
	if err != nil {
		log.Fatal("Invalid image name")
	}
	destRef, err := ParseImageName("oci:test:latest" + ":" + imageTag)
	if err != nil {
		log.Fatal("Failed to set destination name")
	}
	return copy.Image(ctx, policyContext, destRef, srcRef, &copy.Options{})
}

// ParseImageName converts a URL-like image name to a types.ImageReference. This function is taken from Skopeo.
func ParseImageName(imgName string) (types.ImageReference, error) {
	// Keep this in sync with TransportFromImageName!
	parts := strings.SplitN(imgName, ":", 2)
	if len(parts) != 2 {
		return nil, errors.Errorf(`Invalid image name "%s", expected colon-separated transport:reference`, imgName)
	}
	transport := transports.Get(parts[0])
	if transport == nil {
		return nil, errors.Errorf(`Invalid image name "%s", unknown transport "%s"`, imgName, parts[0])
	}
	return transport.ParseReference(parts[1])
}

func getPolicyContext() (*signature.PolicyContext, error) {
	var policy *signature.Policy // This could be cached across calls in opts.
	var err error
	policy = &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	if err != nil {
		return nil, err
	}
	return signature.NewPolicyContext(policy)
}
