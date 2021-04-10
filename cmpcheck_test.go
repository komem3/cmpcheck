package cmpcheck_test

import (
	"testing"

	"github.com/komem3/cmpcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testData := analysistest.TestData()
	analysistest.Run(t, testData, cmpcheck.Analyzer, "a")
}
