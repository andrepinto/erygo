package hcl

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl/hcl/ast"
)

func checkHCLKeys(node ast.Node, valid []string) error {
	var list *ast.ObjectList
	switch n := node.(type) {
	case *ast.ObjectList:
		list = n
	case *ast.ObjectType:
		list = n.List
	default:
		return fmt.Errorf("cannot check HCL keys of type %T", n)
	}

	validMap := make(map[string]struct{}, len(valid))
	for _, v := range valid {
		validMap[v] = struct{}{}
	}

	var result error
	for _, item := range list.Items {
		key := item.Keys[0].Token.Value().(string)
		if _, ok := validMap[key]; !ok {
			result = multierror.Append(result, fmt.Errorf(
				"invalid key '%s' on line %d", key, item.Assign.Line))
		}
	}

	return result
}

var (
	exp = regexp.MustCompile(`\$[a-zA-Z0-9_\-\.]+`)
)

func InterpolateVar(arg string, environments ...map[string]string) string {
	return exp.ReplaceAllStringFunc(arg, func(arg string) string {
		stripped := arg[9 : len(arg)-1]
		for _, env := range environments {
			if value, ok := env[stripped]; ok {
				return value
			}
		}

		return arg
	})
}

func InterpolateArray(args []string, environments map[string]string) []string {
	for v, k := range args {
		args[v] = InterpolateVar(k, environments)
	}
	return args
}

func InterpolateMapString(args map[string]string, environments map[string]string) map[string]string {
	for v, k := range args {
		args[v] = InterpolateVar(k, environments)
	}
	return args
}
