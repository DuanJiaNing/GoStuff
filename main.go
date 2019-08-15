package main

import (
	"encoding/json"
	"fmt"
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

func main() {

	//fmt.Println(MergeYears([]Year{2011, 2013, 2004, 2017, 2014, 2015, 2010, 2018, 2019, 2020}))
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

	t()

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
	sort.SliceStable(years, func(i, j int) bool {
		return years[i] < years[j] // asc
	})

	var ys [][]int
	ysRange := []int{int(years[0]), 1}
	last := years[0]
	length := len(years)
	for i := 1; i < length; i++ {
		cur := years[i]
		if cur-last == 1 {
			ysRange[1]++
		} else {
			ys = append(ys, ysRange)
			ysRange = []int{int(cur), 1}
		}

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
