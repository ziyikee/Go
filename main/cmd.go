package main

import (
	"Golint/linters"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linters.PassMutexByValueAnalyzer)
	//multichecker.Main(linters.HGAnalyzer, linters.WgAddAnalyzer, linters.WaitInLoopAnalyzer, linters.ClosureErrAnalyzer)
}
