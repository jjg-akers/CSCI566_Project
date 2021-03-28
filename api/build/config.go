package build

import "github.com/urfave/cli/v2"

type AppConfig struct {
	RegistrarConfig *EthRegistrarConfig
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		RegistrarConfig: NewEthRegistrarConfig(),
	}
}

func LoadAppConfig(confg *AppConfig) []cli.Flag {

	var flags []cli.Flag

	flags = append(flags, LoadEthRegistrarConfig(confg.RegistrarConfig)...)
	
	return flags
}
