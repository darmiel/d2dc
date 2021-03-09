package main

import "github.com/urfave/cli/v2"

func CommandDockerRun() *cli.Command {
	return &cli.Command{
		Name: "docker",
		Subcommands: []*cli.Command{
			CommandRun(),
		},
		Usage: "Alias to ./",
	}
}
