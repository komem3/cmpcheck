package cmpcheck_test

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/komem3/cmpcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDeepequal(t *testing.T) {
	testData := analysistest.TestData()
	analysistest.Run(t, testData, cmpcheck.Analyzer, "deepequal")
}

func TestCmpDiff(t *testing.T) {
	mainPath := filepath.Join("cmd", "cmpcheck", "main.go")
	pkg := filepath.Join(analysistest.TestData(), "src", "cmpdiff")
	cmd := exec.Command("go", "run", mainPath, pkg)

	errs := []string{
		pkg + "/cmpdiff.go:14:2: string and int are not same type",
		pkg + "/cmpdiff.go:15:2: required cmpopts.EquateErrors when compare error",
		"exit status 3",
	}

	stdoutStderr, err := cmd.CombinedOutput()
	if _, ok := err.(*exec.ExitError); !ok {
		t.Errorf("command is not exit: %v", err)
		return
	}
	if diff := cmp.Diff(strings.Join(errs, "\n"), strings.TrimSpace(string(stdoutStderr))); diff != "" {
		t.Errorf("cmpcheck want(-), got(+)\n%s", diff)
	}
}
