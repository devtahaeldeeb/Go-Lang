package main

import "fmt"

func GenerateSubsets(nums []int, index int, current []int) {
	if index == len(nums) {
		fmt.Println(current)
		return
	}

	current = append(current, nums[index])
	GenerateSubsets(nums, index+1, current)

	current = current[:len(current)-1]

	GenerateSubsets(nums, index+1, current)
}


func main() {
	GenerateSubsets([]int{1, 2, 3}, 0, []int{})
}
