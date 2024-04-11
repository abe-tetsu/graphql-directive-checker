package directivechecker

import (
	"fmt"
	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
	"slices"
)

func Analyzer(types, directives, excludeFieldNames []string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "directivechecker",
		Doc:  "directivechecker determines whether a given type has a given directive.",
		Run:  run(types, directives, excludeFieldNames),
	}
}

func isTargetType(types []string, t *ast.Type) bool {
	if len(types) == 0 {
		return true
	}
	if t == nil {
		return false
	}
	if slices.Contains(types, t.NamedType) {
		return true
	}
	return isTargetType(types, t.Elem)
}

func isExcludeFieldName(ignoreFieldNames []string, fieldName string) bool {
	return slices.Contains(ignoreFieldNames, fieldName)
}

func run(types, directives, excludeFieldNames []string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		fmt.Println("test start")
		fmt.Println("types: ", types)
		fmt.Println("directives: ", directives)
		fmt.Println("excludeFieldNames: ", excludeFieldNames)
		for _, t := range pass.Schema.Types {
			if t.BuiltIn {
				continue
			}

			switch t.Kind {
			case ast.InputObject:
				for _, field := range t.Fields {
					// 除外するフィールドかどうかを判定
					fmt.Println("field: ", field.Name)
					if isExcludeFieldName(excludeFieldNames, field.Name) {
						continue
					}

					if field != nil && field.Type != nil {
						// 引数で渡ってきた型に一致するかどうかを判定
						if isTargetType(types, field.Type) {
							for _, directive := range directives {
								// 引数で渡ってきたディレクティブに一致するかどうかを判定
								if field.Directives.ForName(directive) == nil {
									pass.Reportf(field.Position, "%s has no %s directive", field.Name, directive)
								}
							}
						}
					}
				}

			case ast.Object:
				if t.Kind == ast.Object {
					for _, field := range t.Fields {
						// 除外するフィールドかどうかを判定
						if isExcludeFieldName(excludeFieldNames, field.Name) {
							continue
						}

						for _, arg := range field.Arguments {
							// 引数で渡ってきた型に一致するかどうかを判定
							if isTargetType(types, arg.Type) {
								for _, directive := range directives {
									// 引数で渡ってきたディレクティブに一致するかどうかを判定
									if field.Directives.ForName(directive) == nil {
										pass.Reportf(field.Position, "%s has no %s directive", field.Name, directive)
									}
								}
							}
						}
					}
				}
			}
		}
		return nil, nil
	}
}
