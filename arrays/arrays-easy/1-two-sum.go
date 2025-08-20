package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 10, 10}
	isSorted := isArraySorted(arr)

	fmt.Println(isSorted)
}

func isArraySorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			return false
		}
	}
	return true
}

// BRUTE FORCE METHOD - O(n2)
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{0, 0}
//}

// OPTIMAL METHOD O(N)
func twoSum(nums []int, target int) []int {
	mapData := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		currentValue := nums[i]
		moreNeeded := target - currentValue
		_, exists := mapData[moreNeeded]
		if exists {
			return []int{mapData[moreNeeded], i}
		} else {
			mapData[currentValue] = i
		}
	}
	return []int{0, 0}
}
