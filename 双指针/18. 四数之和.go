package 双指针

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for first := 0; first < len(nums)-3; first++ {
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < len(nums)-2; second++ {
			if second != first+1 && nums[second] == nums[second-1] {
				continue
			}

			third := second + 1
			fourth := len(nums) - 1
			for third < fourth {
				if third != second+1 && nums[third] == nums[third-1] {
					third++
					continue
				}

				sum := nums[first] + nums[second] + nums[third] + nums[fourth]
				if sum == target {
					ret = append(ret, []int{nums[first], nums[second], nums[third], nums[fourth]})
					third++
				} else if sum > target {
					fourth--
				} else {
					third++
				}
			}
		}
	}

	return ret
}
