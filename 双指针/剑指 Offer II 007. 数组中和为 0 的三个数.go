package 双指针

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for first := 0; first < len(nums)-2; first++ {
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}

		third := len(nums) - 1
		for second := first + 1; second < len(nums)-1; second++ {
			if second != first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}

			if second == third {
				break
			}

			if nums[first]+nums[second]+nums[third] == 0 {
				ret = append(ret, []int{first, second, third})
			}
		}
	}

	return ret
}
