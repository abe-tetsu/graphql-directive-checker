package directivechecker

import (
	"github.com/gqlgo/gqlanalysis/analysistest"
	"testing"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	types := []string{"ID"}
	directives := []string{"id"}
	excludeFieldNames := []string{}
	analysistest.Run(
		t,
		testdata,
		Analyzer(types, directives, excludeFieldNames),
		"a",
	)
}
