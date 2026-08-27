package main

func FirstRune(s string) rune {
	if s == "" {
		return 0
	}

	for _, c := range s {
		return c
	}

	return 0
}
