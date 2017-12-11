package main

import (
	"fmt"
	"os"

	"github.com/wrfly/reg/cmds"
	"github.com/wrfly/reg/types"
	"github.com/wrfly/reg/utils"
	v "github.com/wrfly/reg/version"

	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

var (
	RegistryAddr string
	Credential   types.Credential
)

// load config file (if exist)
func appBefore(c *cli.Context) error {
	// set log-level
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	dc, err := utils.ParseDockerCondig()
	if err != nil {
		return err
	}

	for addr, auth := range dc.Auths {
		RegistryAddr = addr
		Credential, err = utils.ParseAuth(auth.Auth)
		if err != nil {
			return err
		}
		// use the first credential
		break
	}

	logrus.Debugf("Registry: [%s], Auth: [%s:%s]",
		RegistryAddr, Credential.UserName, Credential.PassWord)

	// no registry provided
	if RegistryAddr == "" {
		return fmt.Errorf("Error! No Registry Provided")
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:                  "reg",
		Description:           "docker registry cli",
		Before:                appBefore,
		EnableShellCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "registry",
				Aliases:     []string{"r"},
				Usage:       "docker registry address",
				Destination: &RegistryAddr,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "debug switch",
			},
		},
		// commands
		Commands: []*cli.Command{
			cmds.LsRepo,
		},
		CommandNotFound: func(c *cli.Context, command string) {
			logrus.Errorf("No matching command '%s'", command)
			cli.ShowAppHelp(c)
		},
		// balabala
		Version: fmt.Sprintf("v%s-%s @%s",
			v.Version, v.CommitID, v.BuildAt),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "wrfly",
				Email: "mr.wrfly@gmail.com",
			},
		},
	}

	app.Run(os.Args)
}
