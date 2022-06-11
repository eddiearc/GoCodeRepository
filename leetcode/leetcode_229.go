package leetcode

import "sort"

func majorityElement(nums []int) []int {
	if len(nums) < 0 {
		return nums
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	set := map[int]struct{}{}
	partLen := len(nums) / 3
	for i := 0; i < len(nums)-partLen; i++ {
		if nums[i] == nums[i+partLen] {
			set[nums[i]] = struct{}{}
		}
	}
	var ans []int
	for k := range set {
		ans = append(ans, k)
	}
	return ans
}
