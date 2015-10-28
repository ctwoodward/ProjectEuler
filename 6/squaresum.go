/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "fmt"

func sumsquare(i int) (val int) {
	if i > 1 {
		val = (i * i) + sumsquare(i-1)
	} else {
		val = 1
	}
	return val
}
func squaresum(i int) (val int) {

	for j := i; j > 0; j-- {
		val = val + j
	}
	val = val * val
	return val
}

func main() {
	fmt.Println(squaresum(100) - sumsquare(100))
}
