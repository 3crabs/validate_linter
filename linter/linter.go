package linter

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "validate_linter",
		Doc:   "validate rules for dto",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			switch s := node.(type) {
			case *ast.StructType:
				checkStruct(pass, s.Fields)
			}
			return true
		})
	}
	return nil, nil
}

func checkStruct(pass *analysis.Pass, fields *ast.FieldList) {
	for _, field := range fields.List {
		fieldType := fmt.Sprintf("%v", field.Type)
		fieldValidatorValue := getValidateRule(field.Tag.Value)
		switch fieldType {
		case "string":
			if err := checkString(fieldValidatorValue); err != nil {
				pass.Reportf(field.Pos(), "%v", err)
			}
		}
	}
}

func checkString(tag string) error {
	if strings.Contains(tag, "gte=0") {
		return errors.New("gte=0 -> omitempty")
	}
	return nil
}

func getValidateRule(tag string) string {
	validateTag := "validate:\""
	i := strings.Index(tag, validateTag)
	tag = tag[i+len(validateTag):]
	i = strings.Index(tag, "\"")
	tag = tag[:i]
	return tag
}
