package main

import (
	"fmt"
	"reflect"
)

type cat int

type pig struct {
	cat
	id   int
	name string `k1:"acd" k2:"def"`
}

func main() {
	//t()
	//t1()
	//t2()
	//t3() t
	t4()
}

func t4() {
	of := reflect.ValueOf(add)
	fmt.Println(of.Type().Name(), of.Type().Kind()) // "" func
	crs := of.Call([]reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	})
	for _, va := range crs {
		fmt.Println(va)
	}

}

func add(a, b int) int {
	return a + b
}

func t3() {

	c := &pig{cat: cat(2), id: 13, name: "tom"}
	c1 := &pig{cat: cat(3), id: 14, name: "Jim"}
	cv := reflect.ValueOf(c)
	pig := cv.Elem()
	fmt.Println(pig, pig.CanAddr(), pig.CanSet())
	pig.Set(reflect.ValueOf(*c1))
	fmt.Println(pig)

	v := reflect.New(reflect.TypeOf(pig))
	fmt.Println(v.Kind())

}

func t2() {
	c := pig{cat: cat(2), id: 13, name: "tom"}
	vc := reflect.ValueOf(c)
	fmt.Println(vc.Interface(), vc.String())
	fmt.Println(vc.FieldByName("cat").Interface())

}

func t1() {

	cc := &pig{}
	tfcc := reflect.TypeOf(cc)
	ptre := tfcc.Elem()                   // ptr point type
	fmt.Println(tfcc.Name(), tfcc.Kind()) // "" ptr
	fmt.Println(ptre.Name(), ptre.Kind()) // pig struct

}

func t() {

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
