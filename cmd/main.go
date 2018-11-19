package main

import (
	"fmt"
	"os"

	client "github.com/andrepinto/erygo/cmd/app"
)

func main() {
	app := client.NewCliApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
