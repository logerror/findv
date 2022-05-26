package scanner

import (
	"context"
	"github.com/logerror/findv/pkg/artifact"
	"github.com/logerror/findv/pkg/types"
)

type Scanner struct {
	driver   Driver
	artifact artifact.Artifact
}

type Driver interface {
	Scan(target string, artifactKey string, options types.ScanOptions) (
		results types.Results, err error)
}

func NewScanner(driver Driver, ar artifact.Artifact) Scanner {
	return Scanner{driver: driver, artifact: ar}
}

func ScanArtifact(ctx context.Context, config types.ScannerConfig, options types.ScanOptions) (types.Report, error) {
	//artifactInfo, err := inspectArtifact(ctx)
	//if err != nil {
	//	return types.Report{}, xerrors.Errorf("failed analysis: %w", err)
	//}

	// direct scan
	//results, err := s.Scan(artifactInfo.Name, artifactInfo.ID, artifactInfo.BlobIDs, options)
	//if err != nil {
	//
	//	return types.Report{}, fmt.Errorf("scan failed: %w", err)
	//}

	return types.Report{
		//SchemaVersion: report.SchemaVersion,
		//ArtifactName:  artifactInfo.Name,
		//ArtifactType:  artifactInfo.Type,
		//Metadata: types.Metadata{
		//	ImageID:     artifactInfo.ImageMetadata.ID,
		//	DiffIDs:     artifactInfo.ImageMetadata.DiffIDs,
		//	RepoTags:    artifactInfo.ImageMetadata.RepoTags,
		//	RepoDigests: artifactInfo.ImageMetadata.RepoDigests,
		//	ImageConfig: artifactInfo.ImageMetadata.ConfigFile,
		//},
		//Results: results,
	}, nil
}
