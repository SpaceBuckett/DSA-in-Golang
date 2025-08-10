package main

func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
