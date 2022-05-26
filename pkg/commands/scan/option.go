package scan

import (
	"github.com/logerror/findv/pkg/commands/option"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

type Option struct {
	option.ArtifactOption
	option.DBOption
	option.CacheOption
	option.ReportOption
	option.SbomOption
	option.GlobalOption
}

func NewOption(c *cli.Context) (Option, error) {
	gc, err := option.NewGlobalOption(c)
	if err != nil {
		return Option{}, xerrors.Errorf("failed to initialize global options: %w", err)
	}

	return Option{
		GlobalOption:   gc,
		ArtifactOption: option.NewArtifactOption(c),
		DBOption:       option.NewDBOption(c),
		ReportOption:   option.NewReportOption(c),
		CacheOption:    option.NewCacheOption(c),
		SbomOption:     option.NewSbomOption(c),
	}, nil
}

func (c *Option) Init() error {
	if err := c.initPreScanOptions(); err != nil {
		return err
	}

	// --clear-cache, --download-db-only and --reset don't conduct the scan
	if c.skipScan() {
		return nil
	}

	if err := c.ArtifactOption.Init(c.Context, c.Logger); err != nil {
		return err
	}
	return nil
}

func (c *Option) initPreScanOptions() error {
	if err := c.ReportOption.Init(c.Context.App.Writer, c.Logger); err != nil {
		return err
	}
	if err := c.DBOption.Init(); err != nil {
		return err
	}
	if err := c.CacheOption.Init(); err != nil {
		return err
	}
	if err := c.SbomOption.Init(c.Context, c.Logger); err != nil {
		return err
	}
	return nil
}

func (c *Option) skipScan() bool {
	if c.ClearCache || c.DownloadDBOnly || c.Reset {
		return true
	}
	return false
}
