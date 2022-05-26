package scan

import (
	"context"
	"github.com/logerror/findv/pkg/artifact"
	"github.com/logerror/findv/pkg/db"
	"github.com/logerror/findv/pkg/log"
	pkgReport "github.com/logerror/findv/pkg/report"
	"github.com/logerror/findv/pkg/scanner"
	"github.com/logerror/findv/pkg/types"
	"github.com/logerror/findv/pkg/utils"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	"os"
)

type ScanType string

const (
	pathScanType ScanType = "ps"
)

func Exec(cliCtx *cli.Context, scanType ScanType) error {
	opt, err := InitExecOption(cliCtx)
	if err != nil {
		return err
	}

	return exec(cliCtx.Context, opt, scanType)
}

func exec(ctx context.Context, opt Option, scanType ScanType) (err error) {
	ctx, cancel := context.WithTimeout(ctx, opt.Timeout)
	defer cancel()

	defer func() {
		if xerrors.Is(err, context.DeadlineExceeded) {
			log.Logger.Warn("Increase --timeout value")
		}
	}()

	executor, err := NewExecutor(opt)
	if err != nil {
		return xerrors.Errorf("init error: %w", err)
	}
	defer executor.Close()

	var report types.Report
	switch scanType {
	case pathScanType:
		if report, err = executor.ScanFilePath(ctx, opt); err != nil {
			return xerrors.Errorf("filesystem scan error: %w", err)
		}
	}

	if err = executor.Report(opt, report); err != nil {
		return xerrors.Errorf("report error: %w", err)
	}

	Exit(opt, report.Results.Failed())

	return nil
}

func (r *Executor) Report(opt Option, report types.Report) error {
	if err := pkgReport.Write(report, pkgReport.Option{
		AppVersion: opt.GlobalOption.AppVersion,
		Format:     opt.Format,
		Output:     opt.Output,
	}); err != nil {
		return xerrors.Errorf("unable to write results: %w", err)
	}

	return nil
}

type Executor struct {
	dbOpen bool
}

type executorOption func(*Executor)

func NewExecutor(cliOption Option, opts ...executorOption) (*Executor, error) {
	r := &Executor{}
	for _, opt := range opts {
		opt(r)
	}

	err := log.InitLogger(cliOption.Debug)
	if err != nil {
		return nil, xerrors.Errorf("logger error: %w", err)
	}

	if err = r.initCache(cliOption); err != nil {
		return nil, xerrors.Errorf("cache error: %w", err)
	}

	if err = r.initDB(cliOption); err != nil {
		return nil, xerrors.Errorf("DB error: %w", err)
	}

	return r, nil
}

func (r *Executor) initDB(c Option) error {
	// todo download the database file

	if err := db.InitDbWithPath(c.CacheDir); err != nil {
		return xerrors.Errorf("error in vulnerability DB initialize: %w", err)
	}
	r.dbOpen = true

	return nil
}

func (r *Executor) initCache(c Option) error {
	utils.SetCacheDir(c.CacheDir)
	return nil
}

func (r *Executor) ScanFilePath(ctx context.Context, opt Option) (types.Report, error) {

	return r.Scan(ctx, opt)
}

func (r *Executor) Scan(ctx context.Context, opt Option) (types.Report, error) {
	report, err := scan(ctx, opt)
	if err != nil {
		return types.Report{}, xerrors.Errorf("scan error: %w", err)
	}

	return report, nil
}

func scan(ctx context.Context, opt Option) (
	types.Report, error) {

	scannerConfig, scanOptions, err := initScannerConfig(opt)
	if err != nil {
		return types.Report{}, err
	}

	report, err := scanner.ScanArtifact(ctx, scannerConfig, scanOptions)
	if err != nil {
		return types.Report{}, xerrors.Errorf("image scan failed: %w", err)
	}
	return report, nil
}

func initScannerConfig(opt Option) (types.ScannerConfig, types.ScanOptions, error) {
	scanOptions := types.ScanOptions{}

	target := opt.Target
	if opt.Input != "" {
		target = opt.Input
	}

	return types.ScannerConfig{
		Target: target,
		ArtifactOption: artifact.Option{
			SkipFiles: opt.SkipFiles,
			SkipDirs:  opt.SkipDirs,
		},
	}, scanOptions, nil
}

func InitExecOption(ctx *cli.Context) (Option, error) {
	opt, err := NewOption(ctx)
	if err != nil {
		return Option{}, xerrors.Errorf("option error: %w", err)
	}

	// initialize options
	if err = opt.Init(); err != nil {
		return Option{}, xerrors.Errorf("option initialize error: %w", err)
	}

	return opt, nil
}

func Exit(c Option, failedResults bool) {
	if c.ExitCode != 0 && failedResults {
		os.Exit(c.ExitCode)
	}
}

func (r *Executor) Close() error {
	var errs error

	if r.dbOpen {
		if err := db.Close(); err != nil {
			return err
		}
	}
	return errs
}
