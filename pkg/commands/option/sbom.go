package option

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var supportedSbomFormats = []string{""}

type SbomOption struct {
	ArtifactType string
	SbomFormat   string
}

func (c *SbomOption) Init(ctx *cli.Context, logger *zap.SugaredLogger) error {
	if ctx.Command.Name != "sbom" {
		return nil
	}

	return nil
}

func NewSbomOption(c *cli.Context) SbomOption {
	return SbomOption{
		ArtifactType: c.String("artifact-type"),
		SbomFormat:   c.String("sbom-format"),
	}
}
