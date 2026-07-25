package main

import "fmt"

// ========== 1. Factorial ==========
func factorial(n int) int {
	if n <= 1 {
		fmt.Print("1")
		return 1
	}
	fmt.Print(n, " * ")
	return n * factorial(n-1)
}

// ========== 2. Power (base^exp) ==========
func power(base, exp int) int {
	if exp == 0 {
		fmt.Print("1")
		return 1
	}
	fmt.Print(base, " * ")
	return base * power(base, exp-1)
}

// ========== 3. Count Digits ==========
func countDigits(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + countDigits(n/10)
}

// ========== 4. Print Even Descending ==========
func printEven(n int) {
	if n < 2 {
		return
	}
	if n%2 == 0 {
		fmt.Print(n, " ")
		printEven(n - 2)
	} else {
		printEven(n - 1)
	}
}

// ========== 5. Print Ascending ==========
func printAsc(n int) {
	if n == 0 {
		return
	}
	printAsc(n - 1)
	fmt.Print(n, " ")
}

func main() {
	// 1. Factorial
	fmt.Print("factorial(5) = ")
	fmt.Println(" =", factorial(5))

	fmt.Print("factorial(4) = ")
	fmt.Println(" =", factorial(4))

	fmt.Println()

	// 2. Power
	fmt.Print("power(2, 4) = ")
	fmt.Println(" =", power(2, 4))

	fmt.Print("power(3, 3) = ")
	fmt.Println(" =", power(3, 3))

	fmt.Println()

	// 3. Count Digits
	fmt.Println("countDigits(12345) =", countDigits(12345))
	fmt.Println("countDigits(7) =", countDigits(7))
	fmt.Println("countDigits(999999) =", countDigits(999999))

	fmt.Println()

	// 4. Print Even Descending
	fmt.Print("printEven(9): ")
	printEven(9)
	fmt.Println()

	fmt.Print("printEven(12): ")
	printEven(12)
	fmt.Println()

	fmt.Println()

	// 5. Print Ascending
	fmt.Print("printAsc(5): ")
	printAsc(5)
	fmt.Println()

	fmt.Print("printAsc(8): ")
	printAsc(8)
	fmt.Println()

	fmt.Print("printAsc(1): ")
	printAsc(1)
	fmt.Println()
}
