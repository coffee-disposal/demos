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
	for x := 0; x < arr_len; x += 1 {
		for y := 0; y < arr_len-1; y += 1 {
			if arr[y] > arr[y+1] {
				temp := arr[y]
				arr[y] = arr[y+1]
				arr[y+1] = temp
			}
		}
	}
	return arr
}
