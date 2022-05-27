package cmd

import (
	"flux/app/config"
	"flux/app/consts"
	"flux/app/db"

	"github.com/urfave/cli/v2"
)

func intFlag(name string, value int, usage string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func stringFlag(name, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func runCommon() error {
	consts.Config = &config.Configuration{}
	err := config.ReadConfig("app.toml", consts.Config)
	if err != nil {
		return err
	}

	consts.DB = db.OpenDatabase(consts.Config.Database.DB_DRIVER, consts.Config.Database.DB_URI)

	return nil
}
