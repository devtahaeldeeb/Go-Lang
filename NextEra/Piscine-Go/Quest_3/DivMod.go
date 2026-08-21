package main

func DivMod(a *int, b *int) {
	x := *a
	y := *b

	*a = x / y
	*b = x % y
}
