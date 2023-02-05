package app

import (
	startcmd "github.com/ea3hsp/mclog/cmd/server/commands/start"
	vercmd "github.com/ea3hsp/mclog/cmd/server/commands/version"
	"github.com/ea3hsp/mclog/version"
	"github.com/urfave/cli/v2"
)

var extraCmds = []*cli.Command{}

// New returns a *cli.App instance.
func New() *cli.App {
	app := cli.NewApp()
	app.Name = "server"
	app.Version = version.Version
	app.Usage = "mclog server"
	app.Description = "mclog server"

	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "api-addr",
			Usage:   "API IP bind address",
			Value:   "0.0.0.0",
			Aliases: []string{"aa"},
			EnvVars: []string{"API_ADDR"},
		},
		&cli.IntFlag{
			Name:    "api-port",
			Usage:   "API TCP IP port",
			Value:   3000,
			Aliases: []string{"ap"},
			EnvVars: []string{"API_PORT"},
		},
		&cli.StringFlag{
			Name:    "log-level",
			Usage:   "log level",
			Value:   "debug",
			Aliases: []string{"ll"},
			EnvVars: []string{"LOG_LEVEL"},
		},
	}
	app.Commands = append([]*cli.Command{
		// version command
		vercmd.Command,
		// start command
		startcmd.Command,
	}, extraCmds...)
	app.Before = func(ctx *cli.Context) error {
		return nil
	}
	app.Authors = []*cli.Author{
		{
			Name:  "Albert Espín",
			Email: "aespin@gcelsa.com",
		},
	}
	return app
}
