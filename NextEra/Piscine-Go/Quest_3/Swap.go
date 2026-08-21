package main

func Swap(a *int, b *int) {
	x := *a
	*a = *b
	*b = x
}
