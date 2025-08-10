package main

func RemoveDuplicates(nums []int) int {
	leftPointer := 0
	for i := 1; i < len(nums); i++ {

		if nums[leftPointer] != nums[i] {
			leftPointer++
			nums[leftPointer] = nums[i]
		}
	}

	//fmt.Println(nums)

	for i := leftPointer + 1; i < len(nums); i++ {
		nums[i] = 0
	}

	//fmt.Println(nums)

	return leftPointer + 1
}
