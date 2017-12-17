package cmds

import (
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

var RmImage = &cli.Command{
	Name:  "rm",
	Usage: "rm repo",
	// Flags: []cli.Flag{
	// 	&cli.StringFlag{
	// 		Name:    "filter",
	// 		Aliases: []string{"f"},
	// 		Usage:   "filter: name",
	// 	},
	// },
	Action: func(c *cli.Context) error {
		logrus.Debug("rm repo")

		repo := c.Args().First()
		logrus.Debugf("rm repo [%s]", repo)
		if repo == "" {
			logrus.Error("empty repo name")
			cli.ShowCommandHelp(c, "rm")
			return nil
		}

		s := strings.Split(repo, ":")
		repoName := s[0]
		tagName := "latest"
		if len(s) > 1 {
			tagName = s[1]
		}

		return r.DeleteImages(repoName, tagName)
	},
}
