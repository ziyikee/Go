package main

import (
	"Golint/linters"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linters.HGAnalyzer)
}
