package artifact

import (
	"context"
	"github.com/aquasecurity/fanal/types"
)

type Option struct {
	SkipFiles []string
	SkipDirs  []string
}

type Artifact interface {
	Inspect(ctx context.Context) (reference types.ArtifactReference, err error)
	Clean(reference types.ArtifactReference) error
}
