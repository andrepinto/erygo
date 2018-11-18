package hcl

import (
	"fmt"
	"io/ioutil"

	"github.com/andrepinto/erygo/pkg/project"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

//Parse ...
func Parse(path string) (*project.Project, error) {

	project := new(project.Project)

	f, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	data, err := hcl.Parse(string(f))

	if err != nil {
		return nil, err
	}

	if err := hcl.DecodeObject(&project, data); err != nil {
		return nil, err
	}

	valid := []string{
		"settings",
		"error",
	}

	list, ok := data.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("error parsing: file doesn't contain a root object")
	}

	if err := checkHCLKeys(list, valid); err != nil {
		return nil, err
	}

	if o := list.Filter("error"); len(o.Items) > 0 {
		if err := parseError(project, o); err != nil {
			return nil, fmt.Errorf("error parsing 'error': %s", err)
		}
	}

	return project, nil
}

func parseError(result *project.Project, list *ast.ObjectList) error {

	for v, item := range list.Items {

		tag := item.Keys[0].Token.Value().(string)

		result.Error[v].Name = tag

	}

	return nil
}
