package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := &A{i: 5}
	a.AF()
	fmt.Println(a.i)
	a.AFP()
	fmt.Println(a.i)

	var b D = &A{i: 5}
	b.AF()
	fmt.Println(b)

	var s *A
	fmt.Println(s, reflect.TypeOf(s), s == nil)
	fmt.Println(interface{}(s), reflect.TypeOf(interface{}(s)), interface{}(s) == nil)

	fmt.Println(unsafe.Sizeof("s"))
}

type D interface {
	AF()
}

type A struct {
	i int
}

func (a A) AF() {
	a.i = 2
}

func (a *A) AFP() {
	a.i = 3
}
