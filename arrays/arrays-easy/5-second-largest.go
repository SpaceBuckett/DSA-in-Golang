package main

import "fmt"

func secondLargestInArray(arr []int) {
	secondLargest := arr[0]
	largest := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		}

		if arr[i] > secondLargest && arr[i] < largest {
			secondLargest = arr[i]
		}
	}

	if largest != secondLargest {
		fmt.Println(secondLargest)
		return
	}
	fmt.Println("No second largest found!")
}
