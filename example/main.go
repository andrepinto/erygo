package main

import (
	"fmt"
	"os"
	"path"

	"github.com/andrepinto/erygo/pkg/hcl"
)

func main() {
	project, err := hcl.Parse("./demo.hcl")

	fmt.Println(project, err)

	source, err := createOutputFile("./data", project.Settings.Name)
	project.Gen(source)

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
