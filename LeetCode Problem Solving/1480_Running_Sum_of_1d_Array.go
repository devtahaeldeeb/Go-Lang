package main
import "fmt"

// runningSum calculates the cumulative sum in place.
// Time Complexity: O(n)
// Space Complexity: O(1)

func runningSum(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i-1]
    }
    return nums
}

func main() {
    nums := []int{1, 2, 3, 4}
    result := runningSum(nums)
    fmt.Println(result) // [1 3 6 10]
}
