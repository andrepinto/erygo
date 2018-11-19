package app

import (
	"gopkg.in/urfave/cli.v1"
)

const ()

//ErygoCmdOptions ...
type ErygoCmdOptions struct {
	File string
}

//NewErygoCmdOptions ...
func NewErygoCmdOptions() *ErygoCmdOptions {
	return &ErygoCmdOptions{}
}

//AddFlags ...
func (opts *ErygoCmdOptions) AddFlags(app *cli.App) {

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "file",
			Value:       "",
			Usage:       "file",
			Destination: &opts.File,
		},
	}

	app.Flags = append(app.Flags, flags...)
}
