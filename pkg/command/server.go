package command

import (
	"context"
	"syscall"

	"github.com/micro/cli"
	"github.com/oklog/run"
	srv "github.com/refs/example-extension/pkg/server"
)

func ServerCommand() cli.Command {
	return cli.Command{
		Name: "server",
		Action: func(c *cli.Context) error {
			var (
				gr          = run.Group{}
				ctx, cancel = context.WithCancel(context.Background())
			)

			defer cancel()

			server := srv.NewServer(ctx)

			gr.Add(func() error {
				return server.Run()
			}, func(_ error) {
				cancel()
			})

			gr.Add(run.SignalHandler(ctx, syscall.SIGKILL))
			return gr.Run()
		},
	}
}
