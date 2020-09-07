package main

import (
	"fmt"
)

func isPythagorean(a, b, c int64) bool {
	return (a+b+c == 1000) && (a*a+b*b == c*c)
}

func main() {
	fmt.Println("There exists exactly one Pythagorean triplet for which a + b + c = 1000.")
	fmt.Println("Find the product abc")
	result := int64(0)

	// BRUTE FORCE (with some slight common sense)
	// could generate squares and permutate, but still O(n^2), just n is always smaller
	// TODO - optimize based on pythogorean triplet properties
	maxValue := int64(1000)
finished:
	for a := int64(2); a <= maxValue-2; a += 2 {
		for b := a + 1; b <= maxValue-3; b += 3 {
			c := maxValue - a - b
			if isPythagorean(a, b, c) {
				result = a * b * c
				break finished
			}
		}
	}

	fmt.Println("Answer:", result)
}
