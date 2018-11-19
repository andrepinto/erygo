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

func (resp *Response) GenerateSource(prj *Project) *jen.Statement {

	fn := jen.Func().Id(resp.Name).
		Params(jen.Id("params").Op("...").Func().Params(jen.Op("*").Qual(BasePath, "Response"))).
		Op("*").Qual(BasePath, "Response").
		Block(jen.Id("resp").Op(":=").Id("&erygo.Response").Values(
			jen.Dict{
				jen.Id("Message"):    jen.Lit(resp.Message),
				jen.Id("Success"):    jen.Id("checkStatus").Call(jen.Lit(resp.StatusHTTP)),
				jen.Id("StatusHTTP"): jen.Lit(resp.StatusHTTP),
				jen.Id("Info"): jen.Id("erygo.Info").Values(
					jen.Dict{
						jen.Id("Service"): jen.Lit(prj.Settings.Service),
						jen.Id("Kind"):    jen.Lit(resp.Kind),
					},
				),
				jen.Id("Details"): jen.Index().String().ValuesFunc(func(g *jen.Group) {
					for _, v := range resp.Details {
						g.Lit(v)
					}
				}),
			},
		),
			jen.For(jen.Id("_").Op(",").Id("param").Op(":=").Range().Id("params")).Block(
				jen.Id("param").Call(jen.Id("resp")),
			),
			jen.For(jen.Id("i").Op(",").Id("detail").Op(":=").Range().Id("resp").Dot("Details").Block(
				jen.Id("det").Op(":=").Id("renderTemplate").Call(jen.Id("detail")),
				jen.Id("resp").Dot("Details").Index(jen.Id("i")).Op("=").Id("det"),
			)),
			jen.Return(jen.Id("resp")),
		)
	if resp.Comment != "" {
		return jen.Comment(resp.Name + " response ").Line().
			Add(buildComments(resp.Comment).Add(fn))
	}

	return jen.Comment(resp.Name + " response ").Line().Add(fn).Line()
}
