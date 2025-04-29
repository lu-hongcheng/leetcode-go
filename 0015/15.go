package main

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			x := nums[i] + nums[l] + nums[r]
			if x < 0 {
				l++
			} else if x > 0 {
				r--
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l, r = l+1, r-1
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return
}
