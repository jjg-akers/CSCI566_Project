package registrar

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthRegistrar struct{
	Client *ethclient.Client
}

func (er *EthRegistrar) RegisterVoter([]byte) error{
	fmt.Println("registering voter")
	return nil
}

func (er *EthRegistrar) CheckRegistration([]byte) (bool, error){
	fmt.Println("Checking voter registration")
	return true, nil
}


