package a

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

	if reflect.DeepEqual(s, i) { // want "string and int are not same type"
		print(0)
	}
	if reflect.DeepEqual(custom{}, os.File{}) { // want "a.custom and os.File are not same type"
		print(0)
	}
	if reflect.DeepEqual(Int(0), i) { // want "Int and int are not same type"
		print(0)
	}
	if reflect.DeepEqual(i, 1) {
		print(0)
	}
	if reflect.DeepEqual(errors.New("error"), err) { // want "error comparison should use errors.Is"
		print(0)
	}
}
