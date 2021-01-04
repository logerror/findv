package internal

import (
	"e.welights.net/devsecops/findv/internal/artifacts"
	"e.welights.net/devsecops/findv/pkg/utils"
	"github.com/urfave/cli/v2"
	"io"
	"time"
)

var (
	quietFlag = cli.BoolFlag{
		Name:    "quiet",
		Aliases: []string{"q"},
		Usage:   "suppress progress bar and log output",
		EnvVars: []string{"FINDV_QUIET"},
	}

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

	timeoutFlag = cli.DurationFlag{
		Name:    "timeout",
		Value:   time.Second * 120,
		Usage:   "docker timeout",
		EnvVars: []string{"FINDV_TIMEOUT"},
	}

	formatFlag = cli.StringFlag{
		Name:    "format",
		Aliases: []string{"f"},
		Value:   "table",
		Usage:   "format (table, json, template)",
		EnvVars: []string{"FINDV_FORMAT"},
	}

	globalFlags = []cli.Flag{
		&quietFlag,
		&debugFlag,
		&cacheDirFlag,
	}

	imageFlags = []cli.Flag{}

	// deprecated options
	deprecatedFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "only-update",
			Usage:   "deprecated",
			EnvVars: []string{"FINDV_ONLY_UPDATE"},
		},
	}
)

func NewApplication(version string) *cli.App {
	cli.VersionPrinter = func(c *cli.Context) {
		showVersion(c.App.Version, c.App.Writer)
	}

	app := cli.NewApp()
	app.Name = "findv"
	app.Version = version
	app.ArgsUsage = "target"
	app.Usage = "A simple and comprehensive vulnerability scanner for containers"
	app.EnableBashCompletion = true

	flags := append(globalFlags, setHidden(deprecatedFlags, true)...)
	flags = append(flags, setHidden(imageFlags, true)...)

	app.Flags = flags
	app.Commands = []*cli.Command{
		NewFilesystemCommand(),
	}
	app.Action = artifacts.ImageRun
	return app
}

func showVersion(version string, outputWriter io.Writer) {

}

func setHidden(flags []cli.Flag, hidden bool) []cli.Flag {
	var newFlags []cli.Flag
	for _, flag := range flags {
		var f cli.Flag
		switch pf := flag.(type) {
		case *cli.StringFlag:
			stringFlag := *pf
			stringFlag.Hidden = hidden
			f = &stringFlag
		case *cli.BoolFlag:
			boolFlag := *pf
			boolFlag.Hidden = hidden
			f = &boolFlag
		case *cli.IntFlag:
			intFlag := *pf
			intFlag.Hidden = hidden
			f = &intFlag
		case *cli.DurationFlag:
			durationFlag := *pf
			durationFlag.Hidden = hidden
			f = &durationFlag
		}
		newFlags = append(newFlags, f)
	}
	return newFlags
}

func NewFilesystemCommand() *cli.Command {
	return &cli.Command{
		Name:      "filesystem",
		Aliases:   []string{"fs"},
		ArgsUsage: "dir",
		Usage:     "scan local filesystem",
		Action:    artifacts.FilesystemRun,
		Flags: []cli.Flag{
			&formatFlag,
			&quietFlag,
			&debugFlag,
			&cacheDirFlag,
			&timeoutFlag,
		},
	}
}
