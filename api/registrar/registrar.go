package registrar

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jjg-akers/csci556_project/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Registrar defines the methods necessary for a registrar
type Registrar interface{
	RegisterVoter([]byte) error
	CheckRegistration([]byte) (bool, error)
}

//Sessions are wrappers that allow us to make contract calls without having to pass aorund auth, creds,
//  and call parameters constantly
//A session wraps:
//  a contract instance
//	a bind.CallOpts struct that contains opts for making contract calls
//	a bind.TransactOpts struct that contains auth creds and params for created a valid Ethereum transaction
func NewSession(ctx context.Context) (*contract.ContractSession, error) {

	// read data from keystore file
	keystore, err := os.Open(os.Getenv("KEYSTORE"))
	// keystore, err := os.Open(keystorePath)
	if err != nil {
		return nil, fmt.Errorf("could not load keystore: %w", err)
	}

	defer keystore.Close()

	keystorepass := os.Getenv("KEYSTORE_PASS")
	if keystorepass == ""{
		log.Println("key store pass is empty")
	}

	// get transact opts
	auth, err := bind.NewTransactor(keystore, keystorepass)
	if err != nil {
		return nil, fmt.Errorf("error getting auth: %w", err)
	}

	callOpts := bind.CallOpts{
		From:    auth.From,
		Context: ctx,
	}

	//return session without contract instance - we'll do that after we've obtained a contract instance when we deploy a new contract
	//  or when we load an existing contract
	return &contract.ContractSession{
		TransactOpts: *auth,
		CallOpts:     callOpts,
	}, nil
}

//NewContract deploys a contract if no existing contract exists and assigns to session
//A new contract takes the following params
//session: the client object, which we initialezed in main()
//question: a string containing the question we want the user to answer
//answer: the answer to the question
func NewContract(session *contract.ContractSession, client *ethclient.Client) (*contract.ContractSession, error) {
	log.Println("Creating new contract")

	// deploy quizt to get the contract addr, transaction object, adnd instance
	contractAddress, tx, instance, err := contract.DeployContract(&session.TransactOpts, client)
	if err != nil {
		return nil, fmt.Errorf("could not deploy contract: %w", err)
	}

	//print the adress of the transaction - this can be used to up on ethresan the progress of the transaction
	fmt.Println("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	session.Contract = instance

	//save the address of the deployed contract
	// updateEnvFile("CONTRACTADDR", contractAddress.Hex())
	if err := os.Setenv("CONTRACTADDR", contractAddress.Hex()); err != nil{
		return nil, err
	}

	return session, nil
}

//LoadContract loads a contract if one exists, and assigns to the session
//attempts to load an existing contract by looking for a entry in the env file
func LoadContract(session contract.ContractSession, client *ethclient.Client, contractAddr string) (*contract.ContractSession, error) {
	log.Println("Loading contract")

	// check for existing contract
	// addr := common.HexToAddress(os.Getenv("CONTRACTADDR"))
	addr := common.HexToAddress(contractAddr)


	// create new contract instance
	instance, err := contract.NewContract(addr, client) //creates a new instance of Quiz, bound to a specific deployed contract
	if err != nil {
		// if addr doesn't exists we don't know where to locate our contract on the blockchain
		// log.Fatalf("could not load contract: %v\n", err)
		//log.Println(ErrTransactionWait)
		return nil, fmt.Errorf("could not load contract: %w", err)
	}

	session.Contract = instance

	return &session, nil
}
