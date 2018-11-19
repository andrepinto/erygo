package hcl

import (
	"fmt"
	"io/ioutil"

	"github.com/andrepinto/erygo/pkg/project"
	multierror "github.com/hashicorp/go-multierror"
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
		"message",
	}

	list, ok := data.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("error parsing: file doesn't contain a root object")
	}

	if err := checkHCLKeys(list, valid); err != nil {
		return nil, err
	}

	if o := list.Filter("message"); len(o.Items) > 0 {
		if err := parseMessage(project, o); err != nil {
			return nil, fmt.Errorf("error parsing 'error': %s", err)
		}
	}

	return project, nil
}

func parseMessage(result *project.Project, list *ast.ObjectList) error {

	for _, item := range list.Items {

		tzpe := item.Keys[0].Token.Value().(string)
		name := item.Keys[1].Token.Value().(string)

		if tzpe == project.ErrorType {
			errorData := project.Error{}
			if err := hcl.DecodeObject(&errorData, item); err != nil {
				return multierror.Prefix(err, fmt.Sprintf("erygo.%s:", tzpe))
			}
			errorData.Name = name

			result.Error = append(result.Error, errorData)
		}

		if tzpe == project.ResponseType {
			resposeData := project.Response{}
			if err := hcl.DecodeObject(&resposeData, item); err != nil {
				return multierror.Prefix(err, fmt.Sprintf("erygo.%s:", tzpe))
			}
			resposeData.Name = name
			result.Responses = append(result.Responses, resposeData)
		}

	}

	return nil
}
