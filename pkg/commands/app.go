package commands

import (
	"github.com/logerror/findv/pkg/commands/scan"
	"github.com/logerror/findv/pkg/utils"
	"github.com/urfave/cli/v2"
	"io"
	"time"
)

var (
	debugFlag = cli.BoolFlag{
		Name:    "debug",
		Aliases: []string{"d"},
		Usage:   "debug mode",
		EnvVars: []string{"FINDV_DEBUG"},
	}

	cacheDirFlag = cli.StringFlag{
		Name:    "cache-dir",
		Value:   utils.DefaultCacheDir(),
		Usage:   "cache directory",
		EnvVars: []string{"FINDV_CACHE_DIR"},
	}

	formatFlag = cli.StringFlag{
		Name:    "format",
		Aliases: []string{"f"},
		Value:   "table",
		Usage:   "format (table, json)",
		EnvVars: []string{"FINDV_FORMAT"},
	}

	outputFlag = cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "output file name",
		EnvVars: []string{"FINDV_OUTPUT"},
	}

	timeoutFlag = cli.DurationFlag{
		Name:    "timeout",
		Value:   time.Second * 300,
		Usage:   "timeout",
		EnvVars: []string{"FINDV_TIMEOUT"},
	}

	skipFiles = cli.StringSliceFlag{
		Name:    "skip-files",
		Usage:   "specify the file paths to skip traversal",
		EnvVars: []string{"FINDV_SKIP_FILES"},
	}

	skipDirs = cli.StringSliceFlag{
		Name:    "skip-dirs",
		Usage:   "specify the directories where the traversal is skipped",
		EnvVars: []string{"FINDV_SKIP_DIRS"},
	}

	globalFlags = []cli.Flag{
		&debugFlag,
		&cacheDirFlag,
	}
)

func NewApp(version string) *cli.App {
	cli.VersionPrinter = func(c *cli.Context) {
		showVersion(c.String("cache-dir"), c.String("parse"), c.App.Version, c.App.Writer)

	}
	app := cli.NewApp()

	app.Name = "findv"
	app.Version = version

	app.ArgsUsage = "target"
	app.Usage = "Find Vulnerabilities"
	app.EnableBashCompletion = true
	app.Flags = globalFlags

	app.Commands = []*cli.Command{
		NewPathScanCommand(),
		NewSbomCommand(),
	}

	return app
}

func showVersion(cacheDir, outputFormat, version string, outputWriter io.Writer) {

}

// NewPathScanCommand : detected path or file
func NewPathScanCommand() *cli.Command {

	var subCli = cli.Command{
		Name:      "pathscan",
		Aliases:   []string{"ps"},
		ArgsUsage: "path",
		Usage:     "scan local filesystem for language-specific dependencies and config files",
		Action:    scan.PathScanRun,
		Flags: []cli.Flag{
			&formatFlag,
			&outputFlag,
			&timeoutFlag,
			stringSliceFlag(skipFiles),
			stringSliceFlag(skipDirs),
		},
	}
	return &subCli
}

// NewSbomCommand : TODO: binary parse and detected
func NewSbomCommand() *cli.Command {
	return &cli.Command{}
}

func stringSliceFlag(f cli.StringSliceFlag) *cli.StringSliceFlag {
	return &f
}
