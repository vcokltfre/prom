package cli

import (
	"os"
	"strings"

	cli "github.com/urfave/cli/v2"
	"github.com/vcokltfre/prom/prom/impl"
)

func closeProject(c *cli.Context) error {
	manager, err := impl.GetManager()
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}


	name := c.Args().Get(0)
	if name == "" {
		dirs := strings.Split(cwd, "/")
		name = dirs[len(dirs)-1]
	}

	return manager.CloseProject(name)
}
