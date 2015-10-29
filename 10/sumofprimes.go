/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
package main

import (
	"fmt"
	"math"
)

func main() {

	sum := 0
	for i := 2; i < 2000000; i++ {
		if isprime(i) {
			sum = sum + i
		}
	}
	fmt.Println(sum)

}

func isprime(num int) bool {
	if num == 2 {
		return true
	}
	if num == 3 {
		return true
	}
	if num == 4 {
		return false
	}
	if num == 5 {
		return true
	}

	for j := 2; j <= int(math.Sqrt(float64(num))); j++ {
		if num%j == 0 {
			return false
		}
	}
	return true
}
