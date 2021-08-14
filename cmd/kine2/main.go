package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	app := &cli.App{
		Name:  "kine2",
		Usage: "Minimal etcd v2 API to support SQL backed storage engines",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Whether to enable debug logging",
				Value: false,
			},
		},
		Action: entrypoint,
		Before: func(cctx *cli.Context) (err error) {
			if err = initLogging(cctx); err != nil {
				return
			}

			return
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func initLogging(cctx *cli.Context) (err error) {
	var cfg zap.Config

	if cctx.Bool("debug") {
		cfg = zap.NewDevelopmentConfig()
		cfg.Level.SetLevel(zapcore.DebugLevel)
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.Development = false
	} else {
		cfg = zap.NewProductionConfig()
		cfg.Level.SetLevel(zapcore.InfoLevel)
	}

	cfg.OutputPaths = []string{
		"stdout",
	}

	logger, err := cfg.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)

	return
}

func entrypoint(cctx *cli.Context) (err error) {
	defer func() { _ = zap.L().Sync() }()

	zap.L().Info("hello world")

	return
}
