package consts

import (
	"golder/app/config"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var App *cli.App
var Config *config.Configuration
var DB *gorm.DB
