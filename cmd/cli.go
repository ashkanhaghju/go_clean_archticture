package cmd

import (
	"github.com/urfave/cli/v2"
	"go_web_boilerplate/internal/app"
	"os"
)

func Run() error {
	app := cli.App{
		Commands: []*cli.Command{app.ServeCMD, app.ServeGrpcCMD},
	}

	return app.Run(os.Args)
}
