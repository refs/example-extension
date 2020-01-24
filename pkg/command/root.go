package command

import (
	"os"

	"github.com/micro/cli"
)

func RootCommand() error {
	app := cli.App{
		Commands: []cli.Command{
			ServerCommand(),
		},
	}

	return app.Run(os.Args)
}
