package main

func mergeSort(arr []int, n int) {
	mS(arr, 0, n-1)
}

func mS(arr []int, low int, high int) {
	if low >= high {
		return
	}

	mid := int((low + high) / 2)

	mS(arr, low, mid)
	mS(arr, mid+1, high)

	merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int) {
	left := low
	right := mid + 1
	temp := make([]int, 0, high-low+1)

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, (arr)[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}
