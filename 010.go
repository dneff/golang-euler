package main

import (
	"fmt"

	"github.com/dneff/golang-euler/euler"
)

func main() {
	fmt.Println("Find the sum of all the primes below two million.")
	// fast with existing euler function

	result := int64(0)
	primes := euler.FindPrimes(2000000)

	for _, x := range primes {
		result += int64(x)
	}

	fmt.Println("Answer:", result)
}
