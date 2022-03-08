package model

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"time"
)

func Clix() {
	app := &cli.App{
		Name:     "Akseach",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "https://github.com/NonAbsolute",
				Email: "fjd@geekzwzs.cn",
			},
		},
		Copyright: "(c) GPL-3.0 License",
		Usage:     "A catalog collection scanning tool",
		UsageText: "./xxx [command] [global options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "url",
				Aliases: []string{"u"},
				Usage:   "Url target",
			},
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "DIY dictionaries file",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "auto",
				Aliases: []string{"a"},
				Usage:   "From File Read url",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "noAuto",
				Aliases: []string{"na"},
				Usage:   "Command input url",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
