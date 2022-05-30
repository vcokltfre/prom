package cli

import (
	"os"
	"path"
	"strings"

	cli "github.com/urfave/cli/v2"
	"github.com/vcokltfre/prom/prom/impl"
)

func initProject(c *cli.Context) error {
	manager, err := impl.GetManager()
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var dir string

	name := c.Args().Get(0)
	if name == "" {
		dirs := strings.Split(cwd, "/")
		name = dirs[len(dirs)-1]
		dir = cwd
	} else {
		dir = path.Join(cwd, name)
	}

	return manager.InitProject(name, dir)
}
