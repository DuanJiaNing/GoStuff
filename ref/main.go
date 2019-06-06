package main

import (
	"fmt"
	"reflect"
)

type cat int

type pig struct {
	cat
	id   int
	name string
}

func main() {
	c := cat(1)
	t := reflect.TypeOf(c)
	fmt.Println(t.Name(), t.Kind()) // cat int

	cc := &c
	t = reflect.TypeOf(cc)
	fmt.Println(t.Name(), t.Kind()) // "" ptr

	ce := t.Elem()
	fmt.Println(ce.Name(), ce.Kind()) // cat int

	p := reflect.TypeOf(pig{})
	co := p.NumField()
	for i := 0; i < co; i++ {
		f := p.Field(i)
		fmt.Println(f.Name, f.Type.Name(), f.Type.Kind(), f.Anonymous, f.PkgPath, f.Tag)
	}

}
