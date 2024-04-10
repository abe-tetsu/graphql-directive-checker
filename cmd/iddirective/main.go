package main

import (
	"fmt"
	"github.com/abe-tetsu/graphql-directive-checker/iddirective"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	fmt.Println("test")
	multichecker.Main(
		iddirective.Analyzer(),
	)
}
