package cmpdiff

import (
	"errors"

	"github.com/google/go-cmp/cmp"
)

func f() {
	s := "hello world"
	i := 1
	var err error

	cmp.Diff(s, i)
	cmp.Diff(errors.New("error"), err)
}
