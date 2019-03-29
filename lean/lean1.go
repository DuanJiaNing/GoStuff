package main

import (
	"container/list"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math"
	_ "os" // 匿名导入，不使用包内方法或type，但包会被编译且 init 方法会调用
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

// ----------------------------------------------------------------------------------------------------基本语法
// 1 变量声明
var a_ int
var b_ string
var c_ []float32
var d_ func() bool // 函数变量，可用于回调
var e struct {
	x int
}
var (
	a1_ int
	b1  string
)

// 2 变量初始化
var hp int = 100
var hp1 = 100

type IntSlice []int

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p IntSlice) Len() int {
	return len(p)
}

// 3 匿名变量
var fu func(a int, b string)

func fn1(a int, b string) (int, int) {
	return a, len(b)
}

func fn2(f func(int, string) (int, int)) int {
	a, b := f(2, "abv")
	return b + a
}

// 4 字符串
var str = "str"

// 多行输入，反引号，所有的转义字符都无效
var str1 = ` 
<html>
<h1>abc is "ABC"</h1>
</html>
`

// 5 字符类型
var by uint8 = 'a' // ASCII 的一个字符
var by1 byte = 'a' // 同 uint8
var ch int32 = '我' // UTF-8 字符
var ch1 rune = 'a' // 同 int32

// 6 整形类型转换截断
var a3 int32 = math.MaxInt32
var a4 int16 = int16(a3) // 截断

// 7 指针
// 1 类型指针 2 切片指针
var a5 *int = &a_
var a6 int = *a5 // 指针取值
// new 创建指针
var a7 *int = new(int)

func fn3() {
	*a7 = a6

	a7 = &a_
	*a7 = 12 // 修改了 a 的值
}

// 8 变量生命周期由编译器决定，根据是否被取地址和是否发生逃逸决定分配在堆还是栈

// 9 const 常量
const pi = math.Pi

// 10 const+iota 模拟枚举
const (
	Arrow    ChipType = iota // 从 0 开始递增生成值
	Shuriken                 // 1
	Rifle                    // 2
)

type ChipType int

func (c ChipType) String() string {
	switch c {
	case Arrow:
		return "Arrow"
	case Shuriken:
		return "Shuriken"
	case Rifle:
		return "Rifle"
	}

	return `N/A`
}

func fn7(c ChipType) {
	fmt.Println(c)
}

// 11 类型别名
type MyInt int              // 新类型
type MyStringAlias = string // 类型别名（非本地类型），编译后原始类型代替，不能添加新方法(放入string定义所在包中后可以添加)

// ----------------------------------------------------------------------------------------------------容器
// 12 数组
var arr = [...]string{"a", "b"}

// 13 slice
var s1 = []int{1, 2, 3}
var s2 = s1[:] // 重置
var s3 []int
var s4 = make([]int, 5, 100)

func appendAndcopy() {
	s4 = append(s4, 1, 2, 3, 4)

	var copyData []int
	copy(copyData, s4)

}

// 14 map
var mp = make(map[string]int)

func fn8() {
	mp["a"] = 2
	mp["b"] = 3

	v, ok_ := mp["v"]
	if ok_ {
		fmt.Println(v)
	} else {
		fmt.Println("not exist")
	}

	for key, value := range mp {
		fmt.Println(key, value)
	}

	//sort.Sort()
	delete(mp, "a")

}

// 并发不安全的 map：fatal error: concurrent map read and map write
func fn9() {
	m := make(map[int]int)

	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	time.Sleep(5 * time.Minute)
}

// 并发安全的 map
func fn10() {
	var scene sync.Map // 不能使用 make 创建

	// 存 Store
	scene.Store("gree", 97)
	scene.Store("london", 27)
	scene.Store("egypt", 35)

	// 取
	fmt.Println(scene.Load("gree"))

	// 删
	scene.Delete("london")

	// 遍历
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true // true 继续迭代，false 终止迭代
	})
}

// 15 list
var li = list.New()
var li1 = list.List{}

func fn11() {
	ele := li.PushBack("a")
	li.PushFront(434)
	// insertAfter
	// insertBefore
	// pushBackList
	// pushFrontList

	li.Remove(ele)

	for i := li.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

// ----------------------------------------------------------------------------------------------------流程控制

// 16 if
func fn12() {

	// if前添加一个执行语句，返回值的作用范围被限制在 if、else
	if a := 100; a > 0 {

	}
}

// 17 for,for 循环可以通过 break goto return panic 强制退出循环,此时 for 中的结束语句不会执行
// for range 遍历数组、切片、字符串、map、通道
func fn13() {

	step := 2
	for ; step > 0; step-- {

	}

	for step := 2; step > 0; step-- {

	}

	for i := 0; ; i++ {
		if i == 100 {
			break
		}
	}

	i := 10
	for {
		if i > 100 {
			break
		}

		i++
	}

	for i < 10 {
		i++
	}

	for y := 1; y <= 9; y++ {
		for i := 1; i <= y; i++ {
			fmt.Printf("%d*%d=%d ", i, y, i*y)
		}
		fmt.Println()
	}
}

// 17 switch case 每个 case 都是独立代码块，无需用 break 显示跳出
func fn14() {
	a := "a"
	switch a {
	case "a", "b":
		fmt.Println(a)
	case "c":

	}

	b := 10
	switch {
	case b > 10 && b < 100:
		fmt.Println(a)
	}

	switch a {
	case "a":
		fmt.Println(a)
		fallthrough // a == "a" 时执行此 case 后继续执行后续 case（不建议使用）
	case "b":
		fmt.Println(a)
	}
}

func fn15() {

OuterLoop:
	for i := 1; i > 10; i++ {
	InnerLoop:
		for j := 1; j > 10; j++ {
			switch {
			case i+j == 9:
				break OuterLoop
			case i+j == 12:
				continue InnerLoop
			}
		}

	}
}

// ----------------------------------------------------------------------------------------------------函数
// 18 声明
func fn16(a, b int) (ra int, rb string, err error) {
	return // return ra, rb, err
}

func fn17() (int, int) {
	return 1, 2
}

func fn18() int {
	var f func() (int, int)
	f = fn17
	a, b := f()
	return a + b
}

// 19 链式处理器
func StringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}

		list[index] = result
	}
}

// 20 匿名函数
func fn19() {
	StringProcess([]string{"abced", "efgh"}, []func(string) string{
		func(str string) string { return str[:len(str)/2] },
		func(str string) string { return strings.ToLower(str) },
		func(str string) string { return string(len(str)) },
	})

	func(data int) {
		fmt.Println(data)
	}(100)

	anfunc := func(data int) {
		fmt.Println(data)
	}
	anfunc(9)

	skill := map[string]func(){
		"fire": func() {
			// fire
		},
		"water": func() {
			// water
		},
	}
	skill["water"]()

}

// 21 结构体实现接口，函数体实现接口
type Invoker interface {
	Call(interface{})
}

type StructCaller struct {
	a int
}

// 结构体实现接口
func (s *StructCaller) Call(p interface{}) {
	fmt.Println("struct call..", p)
}

type FuncCaller func(interface{})

// 函数体实现接口
func (f FuncCaller) Call(p interface{}) {
	f(p) // f 相当于 java 的 this
}

func fn20() {
	var invoker Invoker
	fmt.Println(&invoker) // 0xc04203a1c0
	fmt.Println(invoker)  // nil

	// StructCaller
	var caller *StructCaller = new(StructCaller)
	caller.a = 3

	//invoker = &StructCaller{2}
	invoker = caller
	fmt.Println(caller)   // &{3}
	fmt.Println(*caller)  // {3}
	fmt.Println(&caller)  // 0xc042004030
	fmt.Println(invoker)  // &{3}
	fmt.Println(&invoker) // 0xc04203a1c0
	//invoker1 := StructCaller{} // 对象

	invoker.Call("a") // struct call..a

	// funcCall
	//invoker = new(FuncCaller)
	var funcCall = func(it interface{}) {
		fmt.Println("func call..", it)
	}

	invoker = FuncCaller(funcCall)
	var call FuncCaller = funcCall

	call("1")      // func call..1
	call.Call("2") // func call..2
}

// 22 闭包 Closure
func closure() {
	str := "hello"

	fn := func() {
		str += " dude "
		p(str)
	}

	fn1 := func() {
		str += " girl "
		p(str)
	}

	fn()
	fn1()

	// value 为被捕获到闭包中的变量
	value := 0
	accm := accm(value)

	// accm 闭包会修改 value 的值，value 会跟随闭包的生命周期一直存在（闭包本身就拥有了记忆效应，状态）
	p(accm())
	p(accm())
	p(accm())

	playerGenertor := playerGen("tom")
	// 玩家死亡，生成一个新玩家
	p(playerGenertor())
	p(playerGenertor())
	p(playerGenertor())
}

// 累加器:返回一个闭包函数
func accm(value int) func() int {
	return func() int {
		value++
		//p(&value)
		return value
	}
}

// 玩家生成，只有 name 不同，hp 始终为 150
func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

// 23 类型判断
func fn21(a ...interface{}) {
	for _, a1 := range a {
		//switch a1.(type) {
		//case int:
		//case string:
		//}
		fmt.Print(a1)
		fmt.Print(" ")
	}
}

// 24 可变参数二次传参
func fn22(a ...interface{}) {
	fn21(a, a) // slice 作为整体传递，只传了两个参数
	fmt.Println()
	fn21(a...) // 多参数传递，传了多个参数
}

// 25 defer 延迟处理，可用于释放资源
// 延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。
func fn23() {
	// 将 defer 后的语句放入延迟调用栈
	defer p("last exec")
	defer p("exec 3")
	defer p("exec 2")
	defer p("exec 1")
	defer p("first exec")

	p("main code exec first")
}

// 26 运行时错误
type ErrorStruct struct {
	code int
}

func (err *ErrorStruct) Error() string {
	return fmt.Sprintf("error occ with code=%d", err.code)
}

var errorEquals0 = errors.New("code equals 0")

func mayError(code int) error {
	if code == 0 {
		return errorEquals0
	} else {
		return &ErrorStruct{code}
	}
}

func fn24() {
	p(mayError(500).Error())
	p(mayError(0))
}

// 27 panic 宕机
func fn25() {
	defer p("still exec too even panic") // 宕机信息处理
	defer p("still exec even panic")
	panic("error: crash")
}

func ProtectRun(run func()) (a int) {
	defer func() {
		err := recover() // 接收宕机
		switch err.(type) {
		case runtime.Error:
			p(fmt.Sprintf("runtime error: %s", err))
			a = -1
		case error:
			panic(err)
		default:
			p(err)
			a = -2
		}
	}()

	run()
	return

}

func fn26() {
	p(ProtectRun(func() {
		// 手动宕机
		fn25()
	}))

	p(ProtectRun(func() {
		var a *int // nil
		*a = 1     // 空指针宕机
	}))

	p(ProtectRun(func() {
		panic(mayError(0))
	}))
}

// ----------------------------------------------------------------------------------------------------结构体
// 28 实例化
type Point struct {
	X   int
	Y   int
	err ErrorStruct
}

type XPoint struct {
	Point // 嵌入 Point，类似继承
	Z     int
}

func fn27() {
	// 1
	var point1 Point = Point{}
	var point Point // 结构体实例化后字段的默认值是字段类型的默认值，例如：数值为 0，字符串为空字符串，布尔为 false，指针为 nil 等。
	p(point)
	point.X = 1 + point1.Y
	point = Point{
		X:   1,
		Y:   1,
		err: ErrorStruct{122}}

	// 2
	var po *Point = new(Point)
	po.X = 10 // go 语法糖：(*po).X
	var i *int = new(int)
	p(po.X + *i)

	// 3
	var p1 *Point = &Point{}
	var p2 *Point = &Point{1, 2, ErrorStruct{100}}
	p1.err.code = p1.X + p2.Y

	// 匿名结构体
	msg := &struct {
		id   int
		name string
	}{1, "tom"}

	p(*msg)

}

// 29 方法和接收器
// 方法与函数的最大区别在于方法是面向对象的概念，一个方法需要归属于一个类，即方法需要作用于某一对象（接收器），而函数没有作用对象
// 接收器 + 函数 -> 方法
type Property struct {
	value int
	x     int
}

// p 称之为指针类型接收器，可以是任何类型，类似面向对象的 this，self 关键字，通过 p 能实际修改对象的值
func (p *Property) SetValue(v int) { // 面向对象的设置对象属性
	p.value = v
}

func (p *Property) Value() int { // 面向对象的获取属性值
	return p.value
}

// 运行时接收器的值会复制一份
func (p Property) SetValue_(v int) {
	// 只改变了实际入参的拷贝 p 的 value
	p.value = v
	// 方法调用结束 p 就会回收
}

func (p Property) Copy() Property {
	return Property{p.value, p.x} // 新对象
}

// 小对象值复制速度较快，大对象适合用指针接收器
func fn29() {
	// 指针类型接收器（符合面向对象的写法）
	p2 := new(Property)
	p2.SetValue(12)
	p2.x = p2.Value()

	// 非指针类型接收器
	p1 := &Property{2, 1}
	p1.SetValue_(3) // 修改无效，p1.value 还是 2
	p(p1.value)

	p3 := p1.Copy()
	p(&p3)
	p(&p1)

}

// 30 为任意类型添加方法
type Integer int

// 类型方法
func (in Integer) ValueOf(str string) Integer {
	return -1
}

func ValueOf(str string) Integer {
	return -2
}

func fn30() {
	var del func(str string) Integer // 方法签名一致就可（函数或类型方法都可以）

	in := new(Integer)
	del = in.ValueOf
	p(del("a"))

	del = ValueOf
	p(del("b"))
}

// 事件系统
var eventMapFunc = make(map[int][]func(interface{}))

func RegisterEvent(code int, callback func(interface{})) {
	list := eventMapFunc[code]
	list = append(list, callback)
	eventMapFunc[code] = list
}

func CallEvent(code int, param interface{}) {
	for _, event := range eventMapFunc[code] {
		event(param)
	}
}

func fn31() {
	RegisterEvent(1, func(pa interface{}) {
		p(1, "event execed:", pa)
	})

	RegisterEvent(2, func(pa interface{}) {
		p(2, "event execed:", pa)
	})

	RegisterEvent(3, func(pa interface{}) {
		p(3, "event execed:", pa)
	})

	CallEvent(1, "aa")
	CallEvent(2, "bb")
	CallEvent(3, "cc")
}

// 31 类型内嵌和结构体内嵌
type BasicColor struct {
	R, G, B float32
}

type Color struct {
	Basic BasicColor
	Alpha float32
}
type ColorPlus struct {
	int        // 类型内嵌：字段名也就是类型名，此时类型不能重复
	BasicColor // 结构体内嵌：BasicColor 内嵌到 ColorPlus 中，可忽略字段名直接访问成员，多层内嵌同理

	Alpha float32
	R     string // 覆盖 BasicColor.R
}

type Car struct {
	Engine struct {
		// 内嵌匿名结构体
		Power int
		Type  string
	}
}

func fn32() {
	car := Car{
		Engine: struct {
			// 内嵌匿名结构体的初始化
			Power int
			Type  string
		}{Power: 12, Type: "a"},
	}

	var c Color
	c.Basic.R = float32(car.Engine.Power)
	c.Basic.G = 255
	c.Basic.B = 1
	c.Alpha = 0.5

	var cp_ ColorPlus = ColorPlus{
		int:        1,
		BasicColor: BasicColor{R: 1, G: 1, B: 1},
		Alpha:      2.1,
		R:          "as",
	}
	p(cp_.R)
	p(cp_.BasicColor.R)

	var cp ColorPlus
	cp.R = "1" // 直接访问，多层结构体内嵌时也可以直接访问
	cp.G = 255
	cp.BasicColor.G = 244 // 也可全名访问
	cp.B = 1
	cp.int = 12 // 类型名也就是字段名
	cp.Alpha = 0.5
}

// 32 结构体内嵌模拟继承
type Walkable struct{}
type Flying struct{}
type Human struct {
	Walkable
}
type Bird struct {
	Walkable
	Flying
}

// ----------------------------------------------------------------------------------------------------接口
// 33 声明
type DataWriter interface {
	WriteData(File) error
	WriteDataPlus(string)
	WriteDataPlus_(string)
}

// 隐式接口实现，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计。
// File 实现 WriteDataPlus 和 WriteDataPlus 接口
// FilePlus 实现 WriteDataPlus_ 接口
type File struct {
	FilePlus
}

type FilePlus int

func (file *FilePlus) WriteDataPlus_(str string) {
	p(str)
}

func (file *File) WriteData(fi File) error {
	p(fi)
	return nil
}

func (file *File) WriteDataPlus(str string) {
	p(str)
}

func fn33() {
	var writer DataWriter = new(File)
	writer.WriteData(File{12})

}

// 34 排序
// 实现 sort.Interface 接口才能使用排序
type MyStringList []string

func (m *MyStringList) Len() int {
	return len(*m)
}

func (m *MyStringList) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *MyStringList) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func fn34() {
	ns := &MyStringList{
		"2 as",
		"1 as",
		"3 as",
	}

	sort.Sort(ns)
	p(*ns)

	// sort 包中已经封装的排序
	nms := sort.StringSlice{"a", "d", "c", "b"}
	sort.Sort(nms)
	p(nms)

	is := sort.IntSlice{1, 4, 666, 2, 3, 56, 7}
	sort.Sort(is)
	p(is)

	// sort.Strings(a []string), sort.Ints(a []int), sort.Float64s(a []float64)
	sort.Slice(nms, func(i, j int) bool {
		return nms[i] < nms[j]
	})
}

// 35 嵌套接口 + 接口类型转换
type EmmdInterface string

func (e *EmmdInterface) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (e *EmmdInterface) Close() error {
	return nil
}

func fn35() {
	var cls io.Closer = new(EmmdInterface)
	var clo io.Writer = new(EmmdInterface)
	var cle io.WriteCloser = new(EmmdInterface)
	cls.Close()
	clo.Write([]byte{1, 2, 3})

	cle.Write([]byte{'a', 'b', 2, 3})
	cle.Close()

	// 接口转换
	cl := cle.(io.Closer) // 可能宕机
	cl1 := io.Closer(cle) // 可能宕机
	cl1.Close()
	cl, ok := cle.(io.Closer)
	if ok {
		cl.Close()
	}

	switch clo.(type) { // 三个 case 都满足
	case io.Closer:
		p("Close")
	case io.Writer:
		p("Writer")
	case io.WriteCloser:
		p("WriteCloser")

	}
}

func p(i ...interface{}) {
	fmt.Println(i...) // i -> slice
}

func main() {
	//fn := fn1
	//len := fn2(fn)
	//fmt.Println(len)
	//fn8()
	//fn9()
	//fn10()
	//fn11()
	//fn13()
	//fn20()
	//closure()
	//fn22(1, 2, "aa", 3.33)
	//fn23()
	//fn24()
	//fn25()
	//fn26()
	//fn27()
	//fn29()
	fn35()
}
