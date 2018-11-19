package app

import (
	"sort"

	"gopkg.in/urfave/cli.v1"
)

//NewCliApp ...
func NewCliApp() *cli.App {

	app := cli.NewApp()

	app.Name = "erygo"
	app.Version = "0.1.0"

	opts := NewErygoCmdOptions()
	opts.AddFlags(app)

	app.Action = func(c *cli.Context) error {

		proc := NewErygoApp()
		return proc.Run(opts)
	}

	// sort flags by name
	sort.Sort(cli.FlagsByName(app.Flags))

	return app
}
