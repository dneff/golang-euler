package main

/*
Each new term in the Fibonacci sequence is generated by
adding the previous two terms. By starting with 1 and 2,
the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose
values do not exceed four million, find the sum of the
even-valued terms.
*/

import (
	"fmt"
)

func fibonacci(max int32) []int32 {
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

func main() {
	fmt.Print("By considering the terms in the Fibonacci\nsequence whose values do not exceed four million,\nfind the sum of the even-valued terms.\n")
	maxValue := int32(3999999)
	result := int32(0)

	for _, x := range fibonacci(maxValue) {
		if x%2 == 0 {
			result += int32(x)
		}
	}

	fmt.Println("Answer:", result)

}