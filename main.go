package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func main() {
	app := &cli.App{
		Name: "nyed",
		Action: func(ctx *cli.Context) error {
			return NewEditor(ctx.Args().Slice()).Start()
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error starting nyed: %v", err)
		os.Exit(1)
	}
}
