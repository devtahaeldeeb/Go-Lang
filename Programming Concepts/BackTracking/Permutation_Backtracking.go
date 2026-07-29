package main

import "github.com/01-edu/z01"

func GenPer(current []int, used []bool) {
	if len(current) == len(used) {
		for _, i := range current {
			z01.PrintRune(rune(i + '0'))
		}
		z01.PrintRune(',')
		z01.PrintRune('\n')
		return
	}

	for i := 0; i < len(used); i++ {
		if used[i] == true {
			continue
		} else {
			current = append(current, i)
			used[i] = true
			GenPer(current, used)
			current = current[:len(current)-1]
			used[i] = false
		}
	}
}

func main() {
	GenPer([]int{}, []bool{false, false, false, false})
}
