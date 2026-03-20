package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// fmt.Println("Number analyzer task - Odd/Even/Prime")
	// var num int

	// fmt.Print("Enter a number: ")
	// fmt.Scan(&num)

	// // Even or Odd
	// if num%2 == 0 {
	// 	fmt.Print("Even, ")
	// } else {
	// 	fmt.Print("Odd, ")
	// }
	// // Prime or Not
	// if isPrime(num) {
	// 	fmt.Println("Prime")
	// } else {
	// 	fmt.Println("Not Prime")
	// }

	///word count
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	input, _ := reader.ReadString('\n')

	input = strings.ToLower(input)

	words := strings.Fields(input)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("\nWord Frequency:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

}

// function to check prime
// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
