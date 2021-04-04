package main

import (
	"github.com/darmiel/d2dc/internal/d2dc/cmds"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:                   "🐳 d2dc",
		Usage:                  "Docker to Docker-Compose",
		UsageText:              "d2dc <docker command>",
		Version:                "1.0.0-alpha",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			cmds.CommandDockerRun(), // alias to "run" for "docker run"
			cmds.CommandRun(),
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "quiet",
			},
		},
		EnableBashCompletion: true,
		OnUsageError: func(ctx *cli.Context, err error, _ bool) error {
			log.Println("🤬 Error:", err)
			return err
		},
		Authors: []*cli.Author{
			{Name: "Daniel", Email: "hi@d2a.io"},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Error running:", err)
	}
}
