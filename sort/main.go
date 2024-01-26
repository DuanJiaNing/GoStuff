package main

import "fmt"

func main() {
	nums := []int{3, 4, 1, 2, 8, 5}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums) // [1 2 3 4 5 8]
}
func quickSort(arr []int, l, r int) []int {
	if l >= r {
		return arr
	}

	var p int
	arr, p = partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
	return arr
}

func partition(arr []int, l, r int) ([]int, int) {
	pivot := arr[r]
	i := l
	for j := l; j < r; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return arr, i
}
