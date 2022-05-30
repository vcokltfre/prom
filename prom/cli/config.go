package cli

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
	"github.com/vcokltfre/prom/prom/impl"
)

func modifyConfig(c *cli.Context) error {
	config, err := impl.LoadConfig()
	if err != nil {
		return err
	}

	key := c.Args().Get(0)
	value := c.Args().Get(1)

	if key == "stale_dir" {
		config.StaleDir = value
	} else {
		return cli.NewExitError("Unknown config key: " + key, 1)
	}

	return config.Save()
}

func getConfig(c *cli.Context) error {
	config, err := impl.LoadConfig()
	if err != nil {
		return err
	}

	key := c.Args().Get(0)

	if key == "stale_dir" {
		fmt.Println(config.StaleDir)
	} else {
		return cli.NewExitError("Unknown config key: " + key, 1)
	}

	return nil
}
