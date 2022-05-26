package types

import "github.com/logerror/findv/pkg/artifact"

// ScanOptions holds the attributes for scanning vulnerabilities
type ScanOptions struct {
	VulnType            []string
	SecurityChecks      []string
	ScanRemovedPackages bool
	ListAllPackages     bool
}

type ScannerConfig struct {
	Target string

	// Artifact options
	ArtifactOption artifact.Option
}
