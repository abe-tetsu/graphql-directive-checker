package iddirective_test

import (
	"testing"

	"github.com/abe-tetsu/graphql-directive-checker/iddirective"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, iddirective.Analyzer(), "a")
}
