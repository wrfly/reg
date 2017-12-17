package cmds

import (
	"fmt"
	"strings"
	"sync"

	"github.com/wrfly/reg/types"

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

		repos, err := r.ListRepos(num, last)
		if err != nil {
			logrus.Fatal(err)
		}

		var wg sync.WaitGroup
		var errors []error
		if showTags {
			wg.Add(len(repos))
			for _, repo := range repos {
				go func(repo string) {
					defer wg.Done()
					filter := types.TagsFilter{}
					tags, err := r.ListTags(repo, filter)
					if err != nil {
						errors = append(errors, err)
					}
					fmt.Printf("%s:\t%s\n", repo, strings.Join(tags, ","))
				}(repo)
			}
		} else {
			// not show tags
			for _, repo := range repos {
				fmt.Println(repo)
			}
		}

		for _, err := range errors {
			logrus.Error(err)
		}

		wg.Wait()

		return nil
	},
}
