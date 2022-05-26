package option

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"io"
	"os"
)

type ReportOption struct {
	Format   string
	Output   io.Writer
	output   string
	ExitCode int
}

func NewReportOption(c *cli.Context) ReportOption {
	return ReportOption{
		output: c.String("output"),
		Format: c.String("format"),
	}
}

func (c *ReportOption) Init(output io.Writer, logger *zap.SugaredLogger) error {
	// The output is os.Stdout by default
	if c.output != "" {
		var err error
		if output, err = os.Create(c.output); err != nil {
			return xerrors.Errorf("failed to create an output file: %w", err)
		}
	}

	c.Output = output

	return nil
}
