package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/jjg-akers/csci556_project/build"
	"github.com/jjg-akers/csci556_project/handlers"
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

func main() {

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
