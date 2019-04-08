package pk1

import (
	"fmt"
	"testing"
)

var owner int

// go style getter
func Owner() int {
	return owner
}

type Len struct {
}

func New() Len {
	return Len{}
}

func fn() {
	x, y := 1, 2

	var a = 1;
	var b = 3;
	var c = 4; // should you write code this way ? you should not
	x = a + b + c

	// lexer: go 的词法分析器
	// semicolons: ';'
	if x < y+12<<20 { // lexer wil add semicolons before '{' token automatically, so dont add '\r\n' before '{'

	}

	cc := 12
	ss, cc := 23, 34 // cc only re-assigned
	//cc := 23 // gramme error
	if 1 == 1 {
		cc := 33 // differ with line 33, cc is a new variable in if scope
		cc++
	}
	cc += ss

	for i := 0; i < 10; i++ {
	}

	// like while
	for a < 10 {
		a++
	}

	//for {} // infinite loop

	switch a {
	case 1, 2, 3, 4, 5:
		return
	case -1:
		fallthrough // 实际会执行 case 0 的操作
	case 0:
		return
	}

	// both pointer type
	//file := new(os.File)
	//f2 := os.File{}

}

func Run() {
	var c = new(int) // pointer type, zeroed value
	fmt.Println(*c)
	fmt.Println(c)

	ms := make(map[int]string, 10) // instance vale, value initialized already,can make map, slice, channel only
	fmt.Println(ms)

}

func Test_t1(t *testing.T) {
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}
