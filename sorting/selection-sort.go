package main

import "fmt"

func SelectionSort(arr []int) []int {
	fmt.Println("SelectionSortHere")

	for i := 0; i < len(arr)-1; i++ {
		indexOfMinima := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[indexOfMinima] {
				indexOfMinima = j
			}
		}
		arr[i], arr[indexOfMinima] = arr[indexOfMinima], arr[i]
	}

	return arr
}
