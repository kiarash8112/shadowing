package main

import (
	shadowing "github.com/kiarash8112/shadowing/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(shadowing.ShadowedVarAnalyzer)
}
