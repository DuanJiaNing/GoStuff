package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"sort"
	"time"
)

//import pk1 "GoStuff/pkg1"

type n1 struct {
}

type nn1 struct {
	n1
}

type Year int

// FormatTime format time as string
func FormatTime(t time.Time) string {
	const format = "20060101T01001Z"
	return t.Format(time.RFC3339)
}

var c map[string]string = nil

type vl int

func main() {

	handler := cors.AllowAll().Handler(handle.Handler)
	http.ListenAndServe("",handler)

	a, b, c, d := vl(1), vl(2), vl(3), vl(8)
	vs := []*vl{&a, &b, &c}
	dst := make([]*vl, len(vs))
	copy(dst, vs)
	vs[1] = &d

	fmt.Println(vs)
	fmt.Println(dst)

}

func main1() {

	// remove item from slice
	s := []int{2, 3, 4, 5, 6}
	i := 2
	fmt.Println(append(s[:i], s[i+1:]...)) // 2 3 5 6

	//for e := range c {
	//	fmt.Print(e)
	//}

	//fmt.Print(	time.Date(2008, time.January, 0, 0, 0, 0, 0, time.UTC))
	fmt.Println(MergeYears([]Year{2015, 2018, 2018, 2016, 2016}))
	//fmt.Println(math.MaxFloat32)
	//fmt.Println(float32(math.MaxFloat32))
	//fmt.Println(-math.MaxFloat32)

	//fmt.Println(float32(-math.MaxFloat32))

	//fmt.Println(FormatTime(time.Now()))

	//fmt.Println(time.Now())
	//var f float32
	//var f6 float64
	//var b bool
	//fmt.Println(f, f6, b)

	//t()

	//fmt.Println(mapMonthToQuarter(time.February))
	//fmt.Println(mapMonthToQuarter(time.May))
	//fmt.Println(mapMonthToQuarter(time.September))
	//fmt.Println(mapMonthToQuarter(time.October))
}

func t() {
	type C struct {
		d string
	}

	type a struct {
		A []*C
	}

	type A struct {
		b []byte
	}

	bytes, err := json.Marshal(a{A: []*C{{d: "a"}, {d: "b"}}}.A)
	fmt.Print(err)

	aa := &[]*C{} // 指针类型
	err = json.Unmarshal(bytes, aa)
	fmt.Print(err)

	fmt.Print(aa)
}
func MergeYears(years []Year) [][]int {
	if len(years) == 1 {
		return [][]int{{int(years[0]), 1}}
	}

	sort.SliceStable(years, func(i, j int) bool {
		return years[i] < years[j] // asc
	})

	var ys [][]int
	ysRange := []int{int(years[0]), 1}
	last := years[0]
	length := len(years)
	for i := 1; i < length; i++ {
		cur := years[i]
		if cur == last {
			goto append
		}

		if cur-last == 1 {
			ysRange[1]++
		} else {
			ys = append(ys, ysRange)
			ysRange = []int{int(cur), 1}
		}

	append:
		if i == length-1 {
			ys = append(ys, ysRange)
		}
		last = cur
	}

	return ys
}

type quarter int

const (
	q1 quarter = iota + 1
	q2
	q3
	q4
)

func mapMonthToQuarter(t time.Month) quarter {
	switch t {
	case time.January:
		fallthrough
	case time.February:
		fallthrough
	case time.March:
		return q1
	case time.April:
		fallthrough
	case time.May:
		fallthrough
	case time.June:
		return q2
	case time.July:
		fallthrough
	case time.August:
		fallthrough
	case time.September:
		return q3
	case time.October:
		fallthrough
	case time.November:
		fallthrough
	case time.December:
		return q4
	}

	return q1
}
