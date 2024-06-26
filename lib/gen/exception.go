package gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/weimil/krpc-go/lib/utils"
	"github.com/weimil/krpc-go/types"
	"github.com/ztrue/tracerr"
)

// GenerateException generates an error for a given exception definition.
func GenerateException(f *jen.File, exception *types.Exception) error {
	// Names are given in the format XYZException. We want the more go-like
	// ErrXYZ.
	exceptionName := "Err" + exception.Name[:len(exception.Name)-len("exception")]
	docs, err := utils.ParseXMLDocumentation(exception.Documentation, exceptionName+" - ")
	if err != nil {
		return tracerr.Wrap(err)
	}

	// Define the error type.
	f.Comment(WrapDocComment(docs))
	f.Type().Id(exceptionName).Struct(
		jen.Id("msg").String(),
	)

	// Define the constructor.
	constructorName := "New" + exceptionName
	f.Comment(fmt.Sprintf("%v creates a new %v.", constructorName, exceptionName))
	f.Func().Id(constructorName).Params(
		jen.Id("msg").String(),
	).Op("*").Id(exceptionName).Block(
		jen.Return(jen.Op("&").Id(exceptionName).Values(jen.Dict{
			jen.Id("msg"): jen.Id("msg"),
		})),
	)

	// Define the Error() function.
	f.Comment("Error returns a human-readable error.")
	f.Func().Params(
		jen.Err().Id(exceptionName),
	).Id("Error").Params().String().Block(
		jen.Return(jen.Err().Dot("msg")),
	)

	return nil
}
