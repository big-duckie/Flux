package main

import (
	"golder/app/cmd"
	"golder/app/consts"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	consts.App = cli.NewApp()
	consts.App.Name = "Golder"
	consts.App.Usage = ""
	consts.App.Authors = []*cli.Author{
		{
			Name:  "Big Duckie",
			Email: "bigduckie@outlook.com",
		},
	}
	consts.App.Version = "0.0.1+dev"
	consts.App.Commands = []*cli.Command{
		&cmd.Web,
		&cmd.Migrate,
	}

	if err := consts.App.Run(os.Args); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
