package main

func BasicAtoi(s string) int {
	result := 0

	for _, c := range s {
		digit := c - '0'
		result = result*10 + int(digit)
	}

	return result
}
