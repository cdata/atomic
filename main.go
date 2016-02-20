package main

import (
	"fmt"
	cli "github.com/codegangsta/cli"
	//	cbor "github.com/jbenet/go-multicodec/cbor"
	//	json "github.com/jbenet/go-multicodec/json"
	"github.com/cdata/atomic/command"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Atomic"
	app.Usage = "Decentralized pub/sub thoughtcasting."
	//app.Action = runAtomicFromCli

	app.Commands = []cli.Command{
		{
			Name:   "post",
			Usage:  "Post a message.",
			Action: command.Post,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "message, m",
					Usage: "Message content of the post",
				},
			},
		},
	}

	defer gracefulExit()

	err := app.Run(os.Args)

	if err != nil {
		panic(err)
	}
}

func gracefulExit() {
	if r := recover(); r != nil {
		fmt.Println(r)
		os.Exit(1)
	}
}
