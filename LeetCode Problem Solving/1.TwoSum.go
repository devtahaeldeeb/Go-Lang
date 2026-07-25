func twoSum(nums []int, target int) []int {
	result := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				return result
			}
		}
	}

	return result
}
