package cmds

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

var LsRepo = &cli.Command{
	Name:     "list",
	Aliases:  []string{"ls"},
	HelpName: "list repos",
	Action: func(c *cli.Context) error {
		logrus.Debug("ls repos")
		if c.Args().Len() == 0 {
			cli.ShowCommandHelp(c, "ls")
		}
		return nil
	},
}
