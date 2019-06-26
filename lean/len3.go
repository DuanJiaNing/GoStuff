package main

import "fmt"

func main() {

	fc()

}

// 数组传值 切片传值（隐式传递指针），参照 reflect.sliceHeader uintptr
func fc() {
	tc := func(arr [3]int,arr2 []int) {
		arr[1] = 5
		arr2[1] = 5
	}

	arr := [...]int{2,3,4} // array
	arr2 := []int{2,3,4} // slice

	fmt.Println(arr,arr2) // [2 3 4] [2 3 4]
	tc(arr,arr2)
	fmt.Println(arr,arr2) // [2 3 4] [2 5 4]
}


