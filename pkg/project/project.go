package project

import (
	"fmt"
	"io"

	"github.com/dave/jennifer/jen"
)

//Project ...
type Project struct {
	Settings Settings
	Error    []Error
}

//Settings ...
type Settings struct {
	Name    string
	Service string
	Keys    map[string]string
}

//Gen ...
func (prj *Project) Gen(file io.Writer) error {
	pack := jen.NewFile(prj.Settings.Name)
	pack.PackageComment("Code generated by erygo")

	//consts := []jen.Code{}

	for _, errDecl := range prj.Error {
		fmt.Println(errDecl)
		pack.Line().Add(errDecl.GenerateSource(prj))
	}

	pack.Func().Id("renderTemplate").Params(jen.Id("templText").String()).String().Block(
		jen.Id("buf").Op(":=").Op("&").Qual("bytes", "Buffer").Values(),
		jen.Id("templ").Op(",").Id("err").Op(":=").Qual("text/template", "New").Call(jen.Lit("")).
			Dot("Parse").Call(jen.Id("templText")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err").Dot("Error").Call()),
		),
		jen.Id("err").Op("=").Id("templ").Dot("Execute").Call(jen.Id("buf"), jen.Map(jen.String()).String().Values(jen.DictFunc(func(d jen.Dict) {
			for k, v := range prj.Settings.Keys {
				d[jen.Lit(k)] = jen.Lit(v)
			}
		}))),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err").Dot("Error").Call()),
		),
		jen.Return(jen.Id("buf").Dot("String").Call()),
	)

	return pack.Render(file)
}
