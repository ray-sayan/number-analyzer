package main

import "fmt"

func main() {

	fmt.Println("Number analyzer task - Odd/Even/Prime")
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Even or Odd
	if num%2 == 0 {
		fmt.Print("Even, ")
	} else {
		fmt.Print("Odd, ")
	}

	// Prime or Not
	if isPrime(num) {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}

// function to check prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
