package main

import (
	"fmt"
	"math"
)

/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

func findPrimes(max int64) []int64 {
	//return array of primes less then or equal to max
	results := []int64{2, 3}
	for i := int64(5); i <= max; i += 2 {
		maxPrime := int64(math.Sqrt(float64(i)) + 1)
		isPrime := true
		for _, f := range results {
			if i%f == 0 {
				isPrime = false
				break
			}
			if f > maxPrime {
				break
			}
		}
		if isPrime {
			results = append(results, i)
		}
	}
	return results
}

func main() {
	fmt.Println("What is the largest prime factor of the number 600851475143 ?")
	number := int64(600851475143)
	maxFactor := int64(math.Sqrt(float64(number)) + 1)
	factors := []int64{}
	possiblePrimes := findPrimes(maxFactor)

	for _, x := range possiblePrimes {
		if number%x == 0 {
			factors = append(factors, x)
		}
	}
	fmt.Println("Answer:", factors[len(factors)-1])

}
