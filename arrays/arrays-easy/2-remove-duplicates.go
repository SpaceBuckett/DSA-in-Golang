package main

func RemoveDuplicates(nums []int) int {
	mapData := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, exists := mapData[nums[i]]

		if exists {
			nums[i]++
		} else {
			mapData[nums[i]] = 1
		}
	}
	return len(mapData)
}
