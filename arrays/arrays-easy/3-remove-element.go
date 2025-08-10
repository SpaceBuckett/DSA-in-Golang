package main

import "fmt"

func removeElement(nums []int, val int) int {
	leftPointer := 0

	// SIMPLE SWAP EVERY ELEMENT TO THE PRIOR LOCATIONS
	// THE VAL WHICH ISN'T REQUIRED WILL GO TO THE END

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			temp := nums[i]
			nums[i] = nums[leftPointer]
			nums[leftPointer] = temp

			leftPointer++
		}
	}

	fmt.Println(nums)
	return 0
}
