package option

import (
	"github.com/logerror/findv/pkg/log"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

type DBOption struct {
	Reset          bool
	DownloadDBOnly bool
	SkipDBUpdate   bool
	Light          bool
	NoProgress     bool
	DBRepository   string
}

func (c *DBOption) Init() (err error) {
	if c.SkipDBUpdate && c.DownloadDBOnly {
		return xerrors.New("--skip-db-update and --download-db-only options can not be specified both")
	}
	if c.Light {
		log.Logger.Warn("'--light' option is deprecated and will be removed. See also: https://github.com/aquasecurity/trivy/discussions/1649")
	}
	return nil
}

func NewDBOption(c *cli.Context) DBOption {
	return DBOption{
		Reset:          c.Bool("reset"),
		DownloadDBOnly: c.Bool("download-db-only"),
		SkipDBUpdate:   c.Bool("skip-db-update"),
		Light:          c.Bool("light"),
		NoProgress:     c.Bool("no-progress"),
		DBRepository:   c.String("db-repository"),
	}
}
