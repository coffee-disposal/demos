package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 5, 7, 3, 2, 8, 9, 6, 1, 4}
	sorted_arr := bubble_sort(arr)
	fmt.Println(sorted_arr)
}

func bubble_sort(arr []int) []int {
	arr_len := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for x := 0; x < arr_len-1; x += 1 {
			if arr[x] > arr[x+1] {
				arr[x], arr[x+1] = arr[x+1], arr[x]
				swapped = true
			}
		}
	}
	return arr
}
