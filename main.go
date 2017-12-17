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

// load config file (if exist)
func appBefore(c *cli.Context) error {
	// set log-level
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// make Credential
	registryAddr := ""
	credential := types.Credential{}

	dc, err := utils.ParseDockerCondig()
	if err != nil {
		return err
	}

	for addr, auth := range dc.Auths {
		registryAddr = addr
		credential, err = utils.ParseAuth(auth.Auth)
		if err != nil {
			return err
		}
		// use the first credential
		break
	}

	// last, we use the config you give in cli
	if c.String("registry") != "" {
		registryAddr = c.String("registry")
		credential = types.Credential{
			UserName: c.String("username"),
			PassWord: c.String("password"),
		}
	}

	logrus.Debugf("Registry: [%s], Auth: [%s:%s]", registryAddr,
		credential.UserName, credential.PassWord)

	// check registry
	if registryAddr == "" {
		return fmt.Errorf("Error! No Registry Provided")
	}
	cmds.SetRegistry(registryAddr, credential)

	return nil
}

func main() {
	app := &cli.App{
		Name:                  "reg",
		Usage:                 "docker registry.v2 cli",
		Before:                appBefore,
		EnableShellCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "registry",
				Aliases: []string{"r"},
				EnvVars: []string{"REG_REGISTRY"},
				Usage:   "docker registry address",
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				EnvVars: []string{"REG_USER"},
				Usage:   "docker registry username",
			},
			&cli.StringFlag{
				Name:    "pass",
				Aliases: []string{"p"},
				EnvVars: []string{"REG_PASS"},
				Usage:   "docker registry password",
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
			cmds.LsTags,
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
