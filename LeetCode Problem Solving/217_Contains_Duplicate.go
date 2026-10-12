package main
import "fmt"

// containsDuplicate returns true if any number
// appears more than once
// Time Complexity: O(n) expected
// Space Complexity: O(n)

func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
        _, found := seen[nums[i]]
        if found {
            return true
        }
        seen[nums[i]] = true
    }
    return false
}

func main() {
    testCases := [][]int{
        {1, 2, 3, 1},
        {1, 2, 3, 4},
        {7, 7},
        {5},
        {-1, 0, -1},
    }

    for _, nums := range testCases {
        fmt.Printf(
            "Input: %v | Duplicate: %t\n",
            nums,
            containsDuplicate(nums),
        )
    }
}
