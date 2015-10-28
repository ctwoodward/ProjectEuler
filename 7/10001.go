/*


By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

*/
package main

import "fmt"

func main() {
	count := 0
	for i := 0; i < 1<<31-1; i++ {
		if isprime(i) {
			count++
			if count == 10003 {
				fmt.Println(i)
				return
			}
		}
	}

}

func isprime(num int) bool {
	prime := true
	for j := 2; j <= num/2; j++ {
		if num%j == 0 {
			prime = false
		}
	}
	return prime
}
