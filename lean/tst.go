package main

import (
	"expvar"
	"fmt"
	"reflect"
)

//type Ficker interface {
//	String() string
//}

type Int struct {
}

func (i Int) String() string {
	return "Int"
}

func main() {
	i2 := new(Int)
	get1(i2)
	i := Int{}
	get1(i)
	ints := make([]Int,2)
	get1(ints)
}

func get1(c interface{}) {
	ser := c.(fmt.Stringer)
	ser1 := c.(expvar.Var)
	fmt.Println(reflect.TypeOf(ser).Name())
	fmt.Println(reflect.TypeOf(ser1).Name())

	switch c.(type) {
	case fmt.Stringer:
		fmt.Println("Stringer")
	case expvar.Var:
		fmt.Println("Var")
	}
}
