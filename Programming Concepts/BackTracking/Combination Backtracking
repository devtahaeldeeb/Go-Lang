package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 0 || n >= 10 {
		return
	}

	BackTrack([]int{}, n)
}

func BackTrack(current []int, n int) {
	if len(current) == n {
		for _, c := range current {
			z01.PrintRune(rune(c + '0'))
		}

		if current[0] != 10-n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	start := 0
	if len(current) > 0 {
		start = current[len(current)-1] + 1
	}

	for i := start; i < 10; i++ {
		current = append(current, i)
		BackTrack(current, n)
		current = current[:len(current)-1]
	}
}

func main() {
	PrintCombN(3)
}
