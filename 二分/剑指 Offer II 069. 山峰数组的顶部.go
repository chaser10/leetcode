package 二分

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) >> 1
		if arr[mid+1] > arr[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
