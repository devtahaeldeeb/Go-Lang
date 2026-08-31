package main

func RuneAt(s string, n int) rune {
	if s == "" {
		return 0
	}

	counter := 1

	for _, c := range s {
		if counter == n {
			return c
		}
		counter++
	}

	return 0
}
