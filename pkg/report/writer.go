package report

import (
	"io"
	"time"

	"golang.org/x/xerrors"

	"github.com/logerror/findv/pkg/types"
)

const (
	SchemaVersion = 2
)

// Now returns the current time
var Now = time.Now

type Option struct {
	Format         string
	Output         io.Writer
	Severities     []int
	OutputTemplate string
	AppVersion     string

	// For misconfigurations
	IncludeNonFailures bool
	Trace              bool
}

// Write writes the result to output, format as passed in argument
func Write(report types.Report, option Option) error {
	var writer Writer
	switch option.Format {
	case "json":
		writer = &JSONWriter{Output: option.Output}
	default:
		return xerrors.Errorf("unknown format: %v", option.Format)
	}

	if err := writer.Write(report); err != nil {
		return xerrors.Errorf("failed to write results: %w", err)
	}
	return nil
}

// Writer defines the result write operation
type Writer interface {
	Write(types.Report) error
}
