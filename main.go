package main

import (
	"os"

	"github.com/gatsu420/mary/cmd"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "mary",
		Commands: []*cli.Command{
			cmd.ServeCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
