package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		j, ok := mp[target-v]
		if ok {
			return []int{i, j}
		}
		mp[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
