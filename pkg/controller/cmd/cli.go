package cmd

import (
	"github.com/m-mizutani/go-prj-template/pkg/controller/server"
	"github.com/m-mizutani/go-prj-template/pkg/infra"
	"github.com/m-mizutani/go-prj-template/pkg/infra/db"
	"github.com/m-mizutani/go-prj-template/pkg/usecase"
	"github.com/m-mizutani/go-prj-template/pkg/utils"

	"github.com/m-mizutani/zlog"
	"github.com/urfave/cli/v2"
)

type config struct {
	LogLevel string
	DBName   string
}

func Run(argv []string) error {
	var cfg config
	app := &cli.App{
		Name:  "go-prj-template",
		Usage: "Rename me",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				Aliases:     []string{"l"},
				EnvVars:     []string{"LOG_LEVEL"},
				Usage:       "Log level [trace|debug|info|warn|error]",
				Value:       "info",
				Destination: &cfg.LogLevel,
			},
			&cli.StringFlag{
				Name:        "db-name",
				Aliases:     []string{"d"},
				EnvVars:     []string{"DB_NAME"},
				Destination: &cfg.DBName,
				Required:    true,
			},
		},
		Before: func(ctx *cli.Context) error {
			utils.Logger = utils.Logger.Clone(zlog.WithLogLevel(cfg.LogLevel))
			return nil
		},

		Action: func(ctx *cli.Context) error {
			clients := infra.New(
				infra.WithDB(db.New(cfg.DBName)),
			)
			uc := usecase.New(clients)

			return server.Listen(uc)
		},
	}

	if err := app.Run(argv); err != nil {
		utils.Logger.Err(err).Error("exit with error")
		return err
	}
	return nil
}
