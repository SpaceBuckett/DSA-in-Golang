package main

import "fmt"

func main() {
	arr := []int{3, 5, 81, 1, 0, 2}
	max := maxInArray(arr)

	fmt.Println(max)
}

func maxInArray(arr []int) int {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
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
