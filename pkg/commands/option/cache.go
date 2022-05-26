package option

import "github.com/urfave/cli/v2"

type CacheOption struct {
}

func (c *CacheOption) Init() error {

	return nil
}

func NewCacheOption(c *cli.Context) CacheOption {
	return CacheOption{}
}
