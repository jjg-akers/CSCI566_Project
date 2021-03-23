package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jjg-akers/csci566_project/cmd/build"
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
	return nil
}
