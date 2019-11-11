package main

import (
	"fmt"
	"sort"
)

type strz struct {
	dst int
	dt  int
	dsn string

	st int
	dk string
}

type byKey []*strz

func (b byKey) Len() int {
	return len(b)
}

func (b byKey) Less(i, j int) bool {
	o1 := b[i]
	o2 := b[j]

	// 排序优先级: dsn, dt, dst, dk, st
	return o1.dsn < o2.dsn ||
		(o1.dsn == o2.dsn && o1.dt < o2.dt) ||
		(o1.dsn == o2.dsn && o1.dt == o2.dt && o1.dst < o2.dst) ||
		(o1.dsn == o2.dsn && o1.dt == o2.dt && o1.dst == o2.dst && o1.dk < o2.dk) ||
		(o1.dsn == o2.dsn && o1.dt == o2.dt && o1.dst == o2.dst && o1.dk == o2.dk && o1.st < o2.st)

}

func (b byKey) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {

	ds := []*strz{
		{0, 20, "a", 100, "h",},
		{-1, 20, "a", 100, "h",},
		{-1, 20, "a", 101, "h",},
		{1, 20, "a", 100, "h",},
		{7, 21, "f", 101, "h",},
		{6, 22, "d", 102, "k",},
		{2, 29, "b", 103, "j",},
		{2, 39, "b", 103, "j",},
		{3, 28, "a", 104, "i",},
		{1, 26, "c", 107, "l",},
	}

	sort.Sort(byKey(ds))
	for _, d := range ds {
		fmt.Printf("%+v\n", d)
	}

}
