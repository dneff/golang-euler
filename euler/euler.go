package euler

import (
	"math"
)

//Less returns lesser integer
func Less(x int64, y int64) int64 {

	if x <= y {
		return x
	}
	return y
}

// More returns greater integer
func More(x int64, y int64) int64 {

	if x >= y {
		return x
	}
	return y
}

//Fibonacci returns ordered fib list up to max value
func Fibonacci(max int32) []int32 {
	result := []int32{1, 1}
	a := int32(1)
	b := int32(1)
	next := a + b
	for (a + b) <= max {
		next = a + b
		result = append(result, next)
		a = b
		b = next
	}
	return result
}

// IsPalindrome verifies if positive integer is the same value when reversed
// Could be made more efficient by comparing values as we go
func IsPalindrome(x int64) bool {
	var reversed int64
	for y := int64(x); y > 0; {
		reversed *= 10
		reversed += y % 10
		y /= 10
	}
	return x == reversed
}

// FindPrimes returns array of primes less than or equal to max
func FindPrimes(max int64) []int64 {
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

//FindFactors returns array of factors of an integer
func FindFactors(x int64) []int64 {
	// get all possible primes
	if x == 1 {
		return []int64{1}
	}
	primes := FindPrimes(x)
	factors := []int64{}
	// as long as we find factors, keep circling through...
	factoring := true
	for factoring {
		factoring = false
		for _, f := range primes {
			if x%int64(f) == 0 {
				factors = append(factors, int64(f))
				x = x / int64(f)
				factoring = true
				break
			}
		}
	}
	// if there were no factors, return x
	if len(factors) < 1 {
		factors = append(factors, x)
	}
	return factors
}
