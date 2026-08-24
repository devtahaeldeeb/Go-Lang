package main

func BasicAtoi2(s string) int {
	result := 0

	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}

		digit := c - '0'
		result = result*10 + int(digit)
	}

	return result
}
