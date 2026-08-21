package main

func StrRev(s *string) {
	r := []rune(*s)
	left := 0
	right := len(r) - 1

	for left < right {
		temp := r[left]
		r[left] = r[right]
		r[right] = temp

		left++
		right--
	}
  
	*s = string(r)
}
