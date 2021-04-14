package build

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jjg-akers/csci556_project/registrar"
	"github.com/urfave/cli/v2"
)

type EthRegistrarConfig struct{
	Gateway string
}

func NewEthRegistrarConfig() *EthRegistrarConfig{
	return &EthRegistrarConfig{}
}

func LoadEthRegistrarConfig(config *EthRegistrarConfig) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "gateway",
			EnvVars:     []string{"GATEWAY"},
			Destination: &config.Gateway,
		},
		// &cli.StringFlag{
		// 	Name:        "testhost",
		// 	EnvVars:     []string{"HOST"},
		// 	Destination: &config.Gateway,
		// },
	}
}

func NewEthRegistrar(config *EthRegistrarConfig) (*registrar.EthRegistrar, error){
	
	fmt.Println("gate way from env: ", config.Gateway)
	
	client, err := ethclient.Dial(config.Gateway) //address of testnet
	if err != nil {
		return nil, fmt.Errorf("could not dial eth client: %s", err)
	}
	
	return &registrar.EthRegistrar{
		Client: client,
	}, nil
}