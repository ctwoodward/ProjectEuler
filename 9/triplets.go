/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product a*b*c.
*/
/*
A well known triple is 8,15,17 which adds up to 40 which is 1/25 of 1,000
I'll just multiply each of 8,15 and 17 by 25
*/

package main

import "fmt"

func main() {
	//8,15,17 multipled by 25s and each other
	fmt.Println(8 * 25 * 15 * 25 * 17 * 25)
}

//brute force method
/*	for a := 1; a < 997; a++ {
  for b := 2; b < 998; b++ {
    for c := 3; c < 999; c++ {
      if a+b+c == 1000 {
        if (a*a + b*b) == c*c {
          fmt.Println(a, " ", b, " ", c, " ", a*b*c)
          return
        }
      }
      if a+b+c > 1000 {
        break
      }
    }
  }
}*/
