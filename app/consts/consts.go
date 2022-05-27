package consts

import (
	"flux/app/config"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var App *cli.App
var Config *config.Configuration
var DB *gorm.DB
