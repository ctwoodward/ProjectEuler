/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 33; i++ {
		//		if fib(i) > 4000000 {
		//		fmt.Println("Over 4 million at: ", i)
		//}
		if fib(i)%2 == 0 {
			sum = sum + fib(i) //yeah yeah doubling the work who cares
		}
	}
	fmt.Println(sum)
}

func fib(iteration int) (sum int) {
	if iteration == 1 {
		return 1
	}
	if iteration == 2 {
		return 2
	}
	a := fib(iteration - 2)
	b := fib(iteration - 1)
	return a + b
}
