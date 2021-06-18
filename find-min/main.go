package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 5, 7, 3, 2, 8, 9, 6, 1, 4}
	min := find_min(arr)
	fmt.Println(min)
}

func find_min(arr []int) int {
	arr_len := len(arr)
	curr_min := arr[0]
	for x := 0; x < arr_len; x += 1 {
		if arr[x] < curr_min {
			curr_min = arr[x]
		}
	}
	return curr_min
}
