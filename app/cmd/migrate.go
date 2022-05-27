package cmd

import (
	"flux/app/consts"
	"flux/app/db"

	"github.com/urfave/cli/v2"
)

var Migrate = cli.Command{
	Name:   "migrate",
	Usage:  "Updates database tables.",
	Action: runMigrate,
	Flags:  []cli.Flag{},
}

func runMigrate(c *cli.Context) error {
	runCommon()

	consts.DB.AutoMigrate(db.Mod{}, db.Modpack{}, db.Version{})

	return nil
}
