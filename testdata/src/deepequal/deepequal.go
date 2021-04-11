package deepequal

import (
	"errors"
	"os"
	"reflect"
)

type custom struct {
	int
}

type Int int

func f() {
	s := "hello world"
	i := 1
	var err error

	reflect.DeepEqual(s, i)                // want "string and int are not same type"
	reflect.DeepEqual(custom{}, os.File{}) // want "deepequal.custom and os.File are not same type"
	reflect.DeepEqual(Int(0), i)           // want "Int and int are not same type"
	reflect.DeepEqual(i, 1)
	reflect.DeepEqual(errors.New("error"), err) // want "error comparison should use errors.Is"
}
