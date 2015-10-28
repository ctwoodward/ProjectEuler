/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	biggest := 0
	bigI := 0
	bigJ := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			str := strconv.Itoa(i * j)
			if str == reverse(str) {
				//fmt.Print(i, "*", j, "=", i*j, "=", reverse(str), ",   ")
				if i*j > biggest {
					bigI = i
					bigJ = j
					biggest = i * j
				}

			}
		}
	}
	fmt.Println("Largest Palindrome is a product of ", bigI, "& ", bigJ, " = ", bigI*bigJ)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
