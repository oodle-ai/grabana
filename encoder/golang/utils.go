package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/oodle-ai/grafana-sdk"
)

func panelSpan(panel sdk.Panel) float32 {
	span := panel.Span
	if span == 0 && panel.GridPos.H != nil {
		span = float32(*panel.GridPos.W / 2) // 24 units per row to 12
	}

	return span
}

func qual(pkg string, name string) *jen.Statement {
	return jen.Qual(packageImportPath+"/"+pkg, name)
}

func lit(v interface{}) *jen.Statement {
	return jen.Lit(v)
}

func Map[I any, O any](input []I, mapFunc func(item I) O) []O {
	results := make([]O, 0, len(input))

	for _, item := range input {
		results = append(results, mapFunc(item))
	}

	return results
}
