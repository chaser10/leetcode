package 二分

func findClosestElements(arr []int, k int, x int) []int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] >= x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	cur := left
	ret := []int{}
	left, right = cur-1, cur
	for k > 0 {
		k--
		if left < 0 {
			ret = append(ret, arr[right])
			right++
			continue
		}

		if right >= len(arr) {
			ret = append(ret, arr[left])
			left--
			continue
		}

		if abs(arr[left]-x) <= abs(arr[right]-x) {
			ret = append(ret, arr[left])
			left--
		} else {

			ret = append(ret, arr[right])

			right++
		}

	}

	return arr[left+1 : right]
}
