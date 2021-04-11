package main

import (
	"github.com/komem3/cmpcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(cmpcheck.Analyzer)
}
