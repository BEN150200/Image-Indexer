package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func InitCli() *cli.Command {
	cmd := &cli.Command{
		Name:  "indexer",
		Usage: "Indexes images",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					file := cmd.Args().First()
					fmt.Println("Adding file ", file)
					return AddImage(file)
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					file := cmd.Args().First()
					fmt.Println("Reading file ", file)
					return ReadImage(file)
				},
			},
		},
	}
	return cmd
}
