package main

import "fmt"

func main() {
	is := []int{
		132,
		3,
		4,
		4,
		90,
		130,
		104,
		111,
		60,
		131,
		103,
		50,
		1,
		5,
		4,
		276,
		7,
		54,
		64,
		23,
		21,
		44,
		2,
		7,
		49,
		504,
		210,
		14,
		98,
		86,
		4,
		9,
		1,
		7,
		4,
		2,
		2,
		3,
		9,
		3,
		9,
		4,
	}

	sum := 0
	for _, v := range is {
		sum += v
	}
	fmt.Println(sum)

}
