package scan

import "github.com/urfave/cli/v2"

func PathScanRun(ctx *cli.Context) error {
	return Exec(ctx, pathScanType)
}
