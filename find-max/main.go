package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 5, 7, 3, 2, 8, 9, 6, 1, 4}
	max := find_max(arr)
	fmt.Println(max)
}

func find_max(arr []int) int {
	arr_len := len(arr)
	curr_max := arr[0]
	for x := 0; x < arr_len; x += 1 {
		if arr[x] > curr_max {
			curr_max = arr[x]
		}
	}
	return curr_max
}
