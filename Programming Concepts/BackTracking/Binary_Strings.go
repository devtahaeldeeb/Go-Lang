package main

import "fmt"

func GenerateBinary(current []int, n int) {
	if len(current) == n {
		fmt.Println(current)
		return
	}

	for i := 0; i <= 1; i++ {
		current = append(current, i)

		GenerateBinary(current, n)

		current = current[:len(current)-1]
	}
}

func main() {
	GenerateBinary([]int{}, 3)
}
