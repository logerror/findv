package main

import (
	"github.com/logerror/findv/pkg/commands"
	"log"
	"os"
)

var (
	version = "source-dev"
)

func main() {
	app := commands.NewApp(version)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
