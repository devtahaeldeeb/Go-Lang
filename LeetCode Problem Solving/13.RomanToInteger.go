package main

import "fmt"

func romanToInt(s string) int {

	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	total := 0

	for i := 0; i < n-1; i++ {
		current := values[s[i]]
		next := values[s[i+1]]

		if current < next {
			total -= current
		} else {
			total += current
		}
	}

	total += values[s[n-1]]

	return total
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
