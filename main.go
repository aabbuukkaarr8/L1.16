package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	someIndex := len(arr) / 2
	some := arr[someIndex]

	var left []int
	var right []int

	for i, v := range arr {
		if i == someIndex {
			continue
		}
		if v <= some {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), some), quickSort(right)...)
}

func main() {
	arr := []int{33, 10, 55, 71, 29, 3, 18, 90, 44}
	fmt.Println("Исходный массив:", arr)
	sorted := quickSort(arr)
	fmt.Println("Отсортированный:", sorted)
}
