package app

import (
	"gopkg.in/urfave/cli.v1"
)

const ()

//ErygoCmdOptions ...
type ErygoCmdOptions struct {
	Path string
}

//NewErygoCmdOptions ...
func NewErygoCmdOptions() *ErygoCmdOptions {
	return &ErygoCmdOptions{}
}

//AddFlags ...
func (opts *ErygoCmdOptions) AddFlags(app *cli.App) {

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "path",
			Value:       "",
			Usage:       "path",
			Destination: &opts.Path,
		},
	}

	app.Flags = append(app.Flags, flags...)
}
