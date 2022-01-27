package main

import (
	"github.com/3crabs/validate_linter/linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linter.NewAnalyzer())
}
