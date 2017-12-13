package cmds

import (
	"fmt"

	"github.com/wrfly/reg/registry"

	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

var LsRepo = &cli.Command{
	Name:    "list",
	Usage:   "list repos",
	Aliases: []string{"ls"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "filter",
			Aliases: []string{"f"},
			Usage:   "filter: tag|namespace|image",
		},
		&cli.IntFlag{
			Name:    "num",
			Aliases: []string{"n"},
			Usage:   "show `n` repos(default is all)",
		},
		&cli.BoolFlag{
			Name:  "tags",
			Usage: "show tags",
		},
	},
	Action: func(c *cli.Context) error {
		logrus.Debug("ls repos")

		num := c.Int("num")
		last := ""
		showTags := c.Bool("tags")

		r := registry.Registry{
			RegistryAddr: registryAddr,
			Credential:   credential,
		}

		repos, err := r.ListRepos(num, last)
		if err != nil {
			logrus.Fatal(err)
		}
		if !showTags {
			for _, repo := range repos {
				fmt.Println(repo)
			}
		}

		return nil
	},
}
