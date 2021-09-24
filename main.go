package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func StartServer(c *cli.Context) {
	fmt.Println("TODO: Start server.")
}

func Execute() {
	app := &cli.App{
		Name: "boom",
		Usage: "boom boom challange",
		Action: StartServer,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func main() {
	Execute()
}