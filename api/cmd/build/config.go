package build

import (
	"github.com/urfave/cli/v2"
)

type AppConfig struct {
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func LoadAppConfig(confg *AppConfig) []cli.Flag {

	var flags []cli.Flag

	return flags

}
