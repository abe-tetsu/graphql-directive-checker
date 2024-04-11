package main

import (
	"flag"
	"github.com/abe-tetsu/graphql-directive-checker/directivechecker"
	"github.com/gqlgo/gqlanalysis/multichecker"
	"strings"
)

func main() {
	var types, directives, excludes string
	flag.StringVar(&types, "types", "", "target types. it can specify multiple values separated by `,`")
	flag.StringVar(&directives, "directives", "", "target directives. it can specify multiple values separated by `,`")
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL field name. it can specify multiple values separated by `,`")
	flag.Parse()

	multichecker.Main(
		directivechecker.Analyzer(strings.Split(types, ","), strings.Split(directives, ","), strings.Split(excludes, ",")),
	)
}
