package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const Version = "1.0.0-alpha"

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			CommandDockerRun(), // alias to "run" for "docker run"
			CommandRun(),
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
