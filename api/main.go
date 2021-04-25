package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jjg-akers/csci556_project/build"
	"github.com/jjg-akers/csci556_project/handlers"
	"github.com/jjg-akers/csci556_project/registrar"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

type votingApp struct {
	config *build.AppConfig
}

func newVotingApp() *votingApp {
	return &votingApp{
		config: build.NewAppConfig(),
	}
}

const envLoc = "./../.env"


func loadEnv() error{
	var err error
	if err = godotenv.Load(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
		return err
	}
	return nil
}

func main() {
	// time.Sleep(time.Second*10)

	// if err := loadEnv(); err != nil{
	// 	log.Fatalf("failed to load env: %s", err)
	// }

	api := newVotingApp()

	//build cli app
	var app = cli.NewApp()
	app.Usage = "Voting dApp"

	// load ENV
	app.Flags = build.LoadAppConfig(api.config)

	app.Action = api.startAPI

	if err := app.Run(os.Args); err != nil {
		log.Println("Error running app: ", err)
	}
}

func (api *votingApp) startAPI(cliCtx *cli.Context) error {
	fmt.Println("Starting Application")

	// set up handlers
	registrar, err := build.NewEthRegistrar(api.config.RegistrarConfig)
	if err != nil {
		return fmt.Errorf("failed to build registrar, err: %s", err)
	}

	defer registrar.Client.Close()

	// try to deploy contract
	proceed := false
	var proceedErr error
	for i:= 0; i< 3; i ++ {
		if proceed{
			break
		}

		time.Sleep(time.Second*time.Duration((i*i)))
		if proceedErr = initContract(cliCtx.Context, registrar.Client); proceedErr != nil{
			proceedErr = fmt.Errorf("failed to initialize base contract: %w", proceedErr)
			continue
		}

		proceed = true
	}

	if proceedErr != nil{
		return proceedErr
	}

	accountAddress := common.HexToAddress("2e4adcdc087cc47d7940f14009a56747d7244afb")
	// common: 0x7809169cf6772832f6c80e5334E6DB9e366fBb2d
    balance, _ := registrar.Client.BalanceAt(cliCtx.Context, accountAddress, nil)
    log.Printf("Application wallet Eth Balance: %d\n",balance)

	registrationHandler := &handlers.VoterRegistrationHandler{
		Registrar: registrar,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go startServer(cliCtx.Context, &wg, interrupt, registrationHandler)
	wg.Wait()

	return nil
}

func startServer(ctx context.Context, wg *sync.WaitGroup, interrupt chan os.Signal, registrar http.Handler) {

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.Handle("/registrar", registrar)

	// Start server -- listen at localhost, port 8080
	go func() {
		fmt.Println("starting server of 8040")
		log.Fatal(http.ListenAndServe(":8040", nil))
	}()

	<-interrupt
	wg.Done()
}

//CONTRACTADDR="0x56c43E8b7441Fc1833c750CE2430D8Ac4617aDac"

func initContract(ctx context.Context, client *ethclient.Client) ( error) {
	// create a new session
	// We havn't spedified a value for the contract field yet, we will do that after obtaining a contract instance
	// when we deploy a new contract or whaen we load an existing
	var err error
	
	session, err := registrar.NewSession(ctx)
	if err != nil{
		return fmt.Errorf("error getting new session: %w", err)
	}

	// deploy if not exist
		//load or deploy contract, and update session with contract instance
		
		if addr := os.Getenv("CONTRACTADDR"); addr == ""{
			_, err = registrar.NewContract(session, client)
			if err != nil{
				return err
			}
			log.Println("new contract deployed successfully")
			return nil
		} 

		log.Println("contract address already exists at: ", os.Getenv("CONTRACTADDR"))

		return nil
}