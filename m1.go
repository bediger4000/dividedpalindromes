package main

import (
	"fmt"
	"os"
)

func main() {
	var list []string

	for _, str := range os.Args[1:] {
		list = append(list, str)
	}

	comparisonCount := 0

	max := len(list)

	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			if i == j {
				continue
			}
			comparisonCount++
			if palindrome(list[i], list[j]) {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}

	// Combinatorically, it should make N Permute 2 comparisons,
	// or n!/(n-2)! = n*(n-1) = n^2 - n

	fmt.Printf("made %d comparisons\n", comparisonCount)
	n := len(os.Args) - 1
	fmt.Printf("(n^2 - n) = %d\n", (n*n - n))
}

func palindrome(a, b string) bool {
	runes := append([]rune(a), []rune(b)...)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
