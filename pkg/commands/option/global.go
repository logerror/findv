package option

import (
	"github.com/logerror/findv/pkg/log"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type GlobalOption struct {
	Context *cli.Context
	Logger  *zap.SugaredLogger

	AppVersion string
	Debug      bool
	CacheDir   string
}

func NewGlobalOption(c *cli.Context) (GlobalOption, error) {
	debug := c.Bool("debug")
	logger, err := log.NewLogger(debug)
	if err != nil {
		return GlobalOption{}, xerrors.New("failed to create a logger")
	}

	return GlobalOption{
		Context: c,
		Logger:  logger,

		AppVersion: c.App.Version,
		Debug:      debug,
		CacheDir:   c.String("cache-dir"),
	}, nil
}
