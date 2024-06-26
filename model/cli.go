package model

import (
	"github.com/urfave/cli/v2"
	"os"
	"sort"
	"time"
)

type ReturnInfo struct {
	Url   string
	Dir   string
	Proxy string
	Type  string
}

func Clix() (*ReturnInfo, error) {
	var returnInfo ReturnInfo
	//var flags =

	app := &cli.App{
		Name:     "Akseach",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "https://github.com/NonAbsolute",
				Email: "fjd@geekzwzs.cn",
			},
		},
		Copyright: "(c) GPL-3.0 License",
		Usage:     "A catalog collection scanning tool",
		UsageText: "./xxx [command] [global options]",

		Commands: []*cli.Command{
			{
				Name:    "auto",
				Aliases: []string{"a"},
				Usage:   "From File Read url",
				Action: func(c *cli.Context) error {
					returnInfo.Type = "auto"
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "url",
						Aliases:     []string{"u"},
						Usage:       "Url target",
						Required:    true,
						Destination: &returnInfo.Url,
					},
					&cli.StringFlag{
						Name:        "dir",
						Aliases:     []string{"d"},
						Usage:       "DIY dictionaries file",
						Value:       "Stillness Speaks",
						Destination: &returnInfo.Dir,
					},
					&cli.StringFlag{
						Name:        "proxy",
						Aliases:     []string{"p"},
						Usage:       "Proxy IP",
						Value:       "",
						Destination: &returnInfo.Proxy,
					},
				},
			},
		},
		//Flags: flags,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		return nil, err
	}
	return &returnInfo, nil
}
