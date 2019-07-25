package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 2
var f = "string"

//数组
var arr1 [10]int
var arr2 = [3]int{1, 2, 3}
var arr3 = [...]int{1, 2, 3}

var arr4 = []int{1, 2, 3}
var arr4i = arr4[2]

// 切片
var books []Book // nil slice
var books_ []Book = make([]Book, 20)
var books1 []Book = make([]Book, 20, 30)

//接口
type Describer interface {
	Describe()
}
type St string

func (s St) Describe() {
	fmt.Println("被调用le!")
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	case string:
		fmt.Println("String 变量")
	default:
		fmt.Printf("unknown type\n")
	}
}

func main7() {
	findType("Naveen")
	st := St("我的字符串")
	findType(st)
}

//定义interface
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//实现接口
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main3() {
	name := MyString("Sam Anderson") // 类型转换
	var v VowelsFinder               // 定义一个接口类型的变量
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())

	n := MyString("aaa")
	n.FindVowels()
}

type NilInterface interface {
	call(int) int
}

type Type struct {
	a int
}

func (t Type) call(c int) int {
	return t.a * c
}

func cousAll(is interface{}) {
	//v := is.(int)
	//P(v)
	//P(is)

	//c := is.(Type)
	//P(c.a)

	switch is.(type) {
	case int:
		P("int", is.(int))
	case Type:
		P("Type", is.(NilInterface).call(12))
	default:
		P("default_type", is)
	}

}

func main4() {
	cousAll(2)
	v := Type{a: 2}
	cousAll(v)

	var s interface{} = 23.3
	cousAll(s)
}

func P(p ...interface{}) {
	fmt.Println(p)
}

func cast() {
	var sum int = 12
	var count float32 = 20
	mean := float32(sum) / count
	P(mean)
}

func map_() {
	//var map_ = make(map[int]string) // nil map
	//var map2 map[int]string
	var m map[int]string
	m = make(map[int]string)
	m[1] = "a"
	m[11] = "b"
	m[111] = "bc"

	for key, value := range m {
		P(key, value)
	}

	capital, ok := m[1]
	if ok {
		P(capital)
	}

	P(m)
	delete(m, 1)
	P(m)

}

func range_() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for index, num := range nums {
		if num == 3 {
			P(index)
		}
	}

	kvs := map[string]int{"a": 1, "b": 2}
	for key, value := range kvs {
		P(key)
		P(value)
	}

	for _, v := range []int{1, 2, 3} {
		P(v)
	}

	for index, char := range "abcdef" {
		P(index)
		P(char)
	}

}

func main() {
	//fmt.Println([...]int{3: 4, 5: 3}) // a: b 表示下标为 a 的位置值为 b
	//fmt.Println([...]int{2: 2})
	fmt.Println(test())  // 13
	fmt.Println(test1()) // 1 13 ?
	fmt.Println(test2()) // C
	fmt.Println(test3()) // A
	fmt.Println(test4()) // 13 13
	fmt.Println(test5()) // 13 13
}

func test3() error {
	err := errors.New("B")
	defer func() {
		err = errors.New("C") // write original var, inner func scope
		// write will influence err, return C
	}()

	err = errors.New("A")
	return err // return value is copied out of func scape(fixed)
}

func test2() (err error) {
	err = errors.New("B")
	defer func() {
		err = errors.New("C") // write to err which defined where out of func scape
		// write will influence err, return C
	}()

	err = errors.New("A")
	return err // temporary store, defer write will change it
}

func test() (i int) {
	i++
	defer func() {
		i = 13
	}()

	return
}

func test4() (i int) {
	defer func() {
		fmt.Print(i, " ") // 闭包，执行匿名方法时才去取 i 的值
	}()

	i = 13
	return
}

func test5() int {
	i := 0
	defer func() {
		fmt.Print(i, " ")
	}()

	i = 13
	return i
}

func test1() (i int) {
	defer fmt.Print(i, " ") // 此时值 0 立刻传递给Print方法了，会输出 0

	i = 13
	return
}

func main41() {
	arr := []int{}
	arr = append(arr, 1, 2, 3, 4, 5, 6, 7)
	//s1 := arr[:]
	s2 := arr[2:5]
	s3 := arr[2:]
	s4 := arr[:3]

	//fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(append(arr[:2], arr[3:]...))

}

func slice() {
	s := []int{1, 2, 3} // cp = len = 3

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr[:]
	//s2 := arr[2:5]
	//s3 := arr[2:]
	//s4 := arr[:3]

	s_ := s[1:2]
	//s1_ := s[:2]

	ss := make([]int, 2, 3)

	P(ss)
	P(cap(ss))
	apd := append(ss, 1, 2, 3, 4)
	P(ss)
	P(len(apd))
	P(apd)

	ss1 := make([]int, len(ss)+1, cap(ss)<<1)
	copy(ss1, apd)
	P((ss1))

	P(len(s1))
	P(cap(s1))
	P(s_)
	P(s_[0])

}

type Book struct {
	title   string
	author  string
	subject int
}

func main2() {
	var book = Book{"a", "b", 1}
	var book1 = Book{title: "a", author: "b", subject: 1}
	//var book2 = Book{title: "a", author: "b"}

	book.subject = 3
	book.author = "sd"

	printBook(book)
	//printBook(book2)
	printBook_(&book1)
}

func printBook(book Book) {
	P(book)
	P(book.author)
}

func printBook_(book *Book) {
	P(book)
	P(book.author)
}

func m3() {
	var a = 3
	var pt *int

	pt = &a
	P(&a)
	P(&pt)
	P(pt)
	P(*pt)

}

const (
	_1 = iota
	_2
	_3
	//c_c = len("strstr")
)
const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)
const (
	i_ = 1
	j_
	k_
	l_
)
const (
	c_1 = 2
	c_2
	c_3 = iota
	c_4
	c_5 = 4
	c_6
	c_7 = iota
)

func get() {
	//return "str"
}

func m1() {
	a, c = c, a
	_, c = c, a

	const v1 = "str"

	//var1 := 1
	//var a = var1 + 1
}

//func show(attr)  {
//	fmt.Print(attr)
//}

func init() {
	//P("Init first")
}

//func main() {
/*P("Hello, World!")
P(&a)
P(_2)
P(_1, _2, _3)
P(c_1, c_2, c_3, c_4, c_5, c_6, c_7)
P(i, j, k, l)
P(i_, j_, k_, l_)
*/
//	f1.M1()
/*
	var i int
	for true {
		fmt.Print(i)
		fmt.Print("\n")
		i++
		if (i == 1) {
			goto ctag
		}
	}
ctag:
	P(m2(1, "str1"))
*/

func sqrt(f float64) (float32, error) {
	if f < 0 {
		return float32(f), errors.New("error msg")
	}

	return -1, nil
}

func main5() {
	P(sqrt(-1))

	i2, e := sqrt(-1)
	P(i2, e.Error())
	P(sqrt(1))
}
