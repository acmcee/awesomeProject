package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

func main() {

	app := cli.NewApp()

	app.Name = os.Args[0]
	app.Version = "0.0.1"
	app.Usage = "Hello World"

	var language string

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "p,port",
			Usage: "listening port",
			Value: 8000,
		}, cli.StringFlag{
			Name:        "l,lang",
			Usage:       "read language",
			Value:       "english",
			Destination: &language,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "calc 1+1",
			Category:  "arithmetic",
			Action: func(c *cli.Context) error {
				sum := 0
				n := 0
				var err error
				for _, ns := range c.Args() {
					n, err = strconv.Atoi(ns)
					if err != nil {
						return err
					} else {
						sum += n
					}
				}
				fmt.Printf("%s = %d \n", strings.Join(c.Args(), " + "), sum)
				return nil
			},
		},
		{
			Name:      "sub",
			ShortName: "s",
			Usage:     "calc 2-1",
			Category:  "arithmetic",
			Action: func(c *cli.Context) error {
				ret := 0
				n := 0
				var err error
				for _, ns := range c.Args() {
					n, err = strconv.Atoi(ns)
					if err != nil {
						return err
					} else {
						ret -= n
					}
				}
				fmt.Printf("%s = %d \n", strings.Join(c.Args(), " - "), ret)
				return nil
			},
		},
		{
			Name:      "plus",
			ShortName: "p",
			Usage:     "cal 2*2",
			Category:  "arithmetic",
			Action: func(c *cli.Context) error {
				ret := 1
				n := 0
				var err error
				for _, ns := range c.Args() {
					n, err = strconv.Atoi(ns)
					if err != nil {
						return err
					} else {
						ret *= n
					}
				}
				fmt.Printf("%s = %d \n", strings.Join(c.Args(), " * "), ret)
				return nil
			},
		},
		{
			Name:     "db",
			Usage:    "database operations",
			Category: "database",
			Subcommands: []cli.Command{
				{
					Name:  "insert",
					Usage: "insert data",
					Action: func(c *cli.Context) error {
						fmt.Println("insert subcommand")
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "delete data",
					Action: func(c *cli.Context) error {
						fmt.Println("delete subcommand")
						return nil
					},
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("boom")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}

}
