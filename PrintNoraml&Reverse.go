package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
