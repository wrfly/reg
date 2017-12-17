package cmds

import (
	"fmt"
	"strings"

	"github.com/wrfly/reg/types"

	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

var LsTags = &cli.Command{
	Name:    "tags",
	Usage:   "list repo tags",
	Aliases: []string{"t"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "filter",
			Aliases: []string{"f"},
			Usage:   "filter: name",
		},
	},
	Action: func(c *cli.Context) error {
		logrus.Debug("ls repo tags")

		repo := c.Args().First()
		logrus.Debugf("list [%s] tags", repo)
		if repo == "" {
			logrus.Error("empty repo name")
			cli.ShowCommandHelp(c, "tags")
			return nil
		}

		filter := types.TagsFilter{}
		tags, err := r.ListTags(repo, filter)
		if err != nil {
			return err
		}

		fmt.Printf("%s:\t%s\n", repo, strings.Join(tags, ","))

		return nil
	},
}
