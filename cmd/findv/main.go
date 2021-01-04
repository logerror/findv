package main

import (
	"e.welights.net/devsecops/findv/internal"
	"e.welights.net/devsecops/findv/pkg/log"
	"go.uber.org/zap"
	"os"
)

func main() {
	app := internal.NewApplication("dev")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("cmd run error", zap.Error(err))
	}
}
