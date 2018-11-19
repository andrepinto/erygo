package app

import (
	"fmt"
	"os"
	"path"

	"github.com/andrepinto/erygo/pkg/hcl"
	log "github.com/sirupsen/logrus"
)

// ErygoApp ...
type ErygoApp struct {
}

// NewErygoApp ...
func NewErygoApp() *ErygoApp {
	return &ErygoApp{}
}

//Run ...
func (cli *ErygoApp) Run(options *ErygoCmdOptions) error {

	log.SetLevel(log.DebugLevel)

	log.Infof("Options: \n %v", options)

	project, err := hcl.Parse("./example/demo.hcl")

	log.Info(project, err)

	source, err := createOutputFile("./example/data", project.Settings.Name)
	project.Gen(source)

	return nil
}

func createOutputFile(dir string, name string) (*os.File, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	outputFilePath := path.Join(dir, fmt.Sprintf("%s.go", name))
	outputFile, err := os.Create(outputFilePath)
	return outputFile, err
}
