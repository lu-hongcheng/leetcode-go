package main

func removeElement(nums []int, val int) int {
    n:=len(nums)
    i:=0
    for j:=0;j<n;j++{
        if nums[j]!=val{
            nums[i]=nums[j]
            i++
        }
    }
    return i
}