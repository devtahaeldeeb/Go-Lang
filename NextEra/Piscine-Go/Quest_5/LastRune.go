package main

func LastRune(s string) rune {
	if s == "" {
		return 0
	}

	x := 0
	for _, c := range s {
		x = c
	}

	return x
}
