package project

import (
	"github.com/dave/jennifer/jen"
)

type Response struct {
	Name       string
	StatusHTTP int
	Message    string
	Kind       int
	Comment    string
	Details    []string
}

func (err *Response) GenerateSource(prj *Project) *jen.Statement {

	fn := jen.Func().Id(err.Name).
		Params(jen.Id("params").Op("...").Func().Params(jen.Op("*").Qual(BasePath, "Err"))).
		Op("*").Qual(BasePath, "Err").
		Block(jen.Id("err").Op(":=").Id("&erygo.Err").Values(
			jen.Dict{
				jen.Id("Message"):    jen.Lit(err.Message),
				jen.Id("StatusHTTP"): jen.Lit(err.StatusHTTP),
				jen.Id("Info"): jen.Id("erygo.Info").Values(
					jen.Dict{
						jen.Id("Service"): jen.Lit(prj.Settings.Service),
						jen.Id("Kind"):    jen.Lit(err.Kind),
					},
				),
				jen.Id("Details"): jen.Index().String().ValuesFunc(func(g *jen.Group) {
					for _, v := range err.Details {
						g.Lit(v)
					}
				}),
			},
		),
			jen.For(jen.Id("_").Op(",").Id("param").Op(":=").Range().Id("params")).Block(
				jen.Id("param").Call(jen.Id("err")),
			),
			jen.For(jen.Id("i").Op(",").Id("detail").Op(":=").Range().Id("err").Dot("Details").Block(
				jen.Id("det").Op(":=").Id("renderTemplate").Call(jen.Id("detail")),
				jen.Id("err").Dot("Details").Index(jen.Id("i")).Op("=").Id("det"),
			)),
			jen.Return(jen.Id("err")),
		)
	if err.Comment != "" {
		return jen.Comment(err.Name + " error ").Line().
			Add(buildComments(err.Comment).Add(fn))
	}

	return jen.Comment(err.Name + " error ").Line().Add(fn).Line()
}
