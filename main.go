package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	logrus.Info("reg")

	app := cli.App{
		Name:        "reg",
		Description: "docker registry cli",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	app.Run(os.Args)
}
