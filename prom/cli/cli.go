package cli

import (
	"os"

	cli "github.com/urfave/cli/v2"
)

func RunCLI() {
	cli := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialize a new project",
				Action:  initProject,
			},
			{
				Name:    "close",
				Aliases: []string{"c"},
				Usage:   "Close a project",
				Action:  closeProject,
			},
			{
				Name:    "config",
				Usage:   "Modify or get the configuration",
				Subcommands: []*cli.Command{
					{
						Name:    "get",
						Usage:   "Get the configuration",
						Action:  getConfig,
					},
					{
						Name:    "set",
						Usage:   "Set the configuration",
						Action:  modifyConfig,
					},
				},
			},
		},
	}

	cli.Run(os.Args)
}
