package main

import (
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

func main() {

	ys := []Year{2011, 2013, 2004, 2017, 2014, 2015, 2010, 2018, 2019, 2020}
	sort.SliceStable(ys, func(i, j int) bool {
		return ys[i] < ys[j]
	})
	fmt.Println(ys)
	// [2004 2010 2011 2013 2014 2015 2017]

	var ys1 [][]int
	ysc := []int{int(ys[0]), 1}
	las := ys[0]
	for i := 1; i < len(ys); i++ {
		cur := ys[i]
		if cur-las == 1 {
			ysc[1]++
		} else {
			ys1 = append(ys1, ysc)
			ysc = []int{int(cur), 1}
		}

		if i == len(ys)-1 {
			ys1 = append(ys1, ysc)
		}
		las = cur
	}
	fmt.Println(ys1)

	//fmt.Println(mapMonthToQuarter(time.February))
	//fmt.Println(mapMonthToQuarter(time.May))
	//fmt.Println(mapMonthToQuarter(time.September))
	//fmt.Println(mapMonthToQuarter(time.October))
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
