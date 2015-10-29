/*


By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	for i := 2; i < 1<<31-1; i++ {
		if isprime(i) {
			count++
			if count == 10001 {
				fmt.Println(i)
				return
			}
		}
	}

}

func isprime(num int) bool {
	prime := true
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
			prime = false
			return prime
		}
	}
	return prime
}
