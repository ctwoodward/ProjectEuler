package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	primesBelow1000 []int
	primesBelow8000 []int
)

func init() {
	primesBelow1000 = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
		83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181,
		191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283,
		293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523,
		541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647,
		653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773,
		787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911,
		919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}
	primesBelow8000 = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
		83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181,
		191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283,
		293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523,
		541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647,
		653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773,
		787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911,
		919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997, 1009, 1013, 1019, 1021, 1031,
		1033, 1039, 1049, 1051, 1061, 1063, 1069, 1087, 1091, 1093, 1097, 1103, 1109, 1117, 1123,
		1129, 1151, 1153, 1163, 1171, 1181, 1187, 1193, 1201, 1213, 1217, 1223, 1229, 1231, 1237,
		1249, 1259, 1277, 1279, 1283, 1289, 1291, 1297, 1301, 1303, 1307, 1319, 1321, 1327, 1361,
		1367, 1373, 1381, 1399, 1409, 1423, 1427, 1429, 1433, 1439, 1447, 1451, 1453, 1459, 1471,
		1481, 1483, 1487, 1489, 1493, 1499, 1511, 1523, 1531, 1543, 1549, 1553, 1559, 1567, 1571,
		1579, 1583, 1597, 1601, 1607, 1609, 1613, 1619, 1621, 1627, 1637, 1657, 1663, 1667, 1669,
		1693, 1697, 1699, 1709, 1721, 1723, 1733, 1741, 1747, 1753, 1759, 1777, 1783, 1787, 1789,
		1801, 1811, 1823, 1831, 1847, 1861, 1867, 1871, 1873, 1877, 1879, 1889, 1901, 1907, 1913,
		1931, 1933, 1949, 1951, 1973, 1979, 1987, 1993, 1997, 1999, 2003, 2011, 2017, 2027, 2029,
		2039, 2053, 2063, 2069, 2081, 2083, 2087, 2089, 2099, 2111, 2113, 2129, 2131, 2137, 2141,
		2143, 2153, 2161, 2179, 2203, 2207, 2213, 2221, 2237, 2239, 2243, 2251, 2267, 2269, 2273,
		2281, 2287, 2293, 2297, 2309, 2311, 2333, 2339, 2341, 2347, 2351, 2357, 2371, 2377, 2381,
		2383, 2389, 2393, 2399, 2411, 2417, 2423, 2437, 2441, 2447, 2459, 2467, 2473, 2477, 2503,
		2521, 2531, 2539, 2543, 2549, 2551, 2557, 2579, 2591, 2593, 2609, 2617, 2621, 2633, 2647,
		2657, 2659, 2663, 2671, 2677, 2683, 2687, 2689, 2693, 2699, 2707, 2711, 2713, 2719, 2729,
		2731, 2741, 2749, 2753, 2767, 2777, 2789, 2791, 2797, 2801, 2803, 2819, 2833, 2837, 2843,
		2851, 2857, 2861, 2879, 2887, 2897, 2903, 2909, 2917, 2927, 2939, 2953, 2957, 2963, 2969,
		2971, 2999, 3001, 3011, 3019, 3023, 3037, 3041, 3049, 3061, 3067, 3079, 3083, 3089, 3109,
		3119, 3121, 3137, 3163, 3167, 3169, 3181, 3187, 3191, 3203, 3209, 3217, 3221, 3229, 3251,
		3253, 3257, 3259, 3271, 3299, 3301, 3307, 3313, 3319, 3323, 3329, 3331, 3343, 3347, 3359,
		3361, 3371, 3373, 3389, 3391, 3407, 3413, 3433, 3449, 3457, 3461, 3463, 3467, 3469, 3491,
		3499, 3511, 3517, 3527, 3529, 3533, 3539, 3541, 3547, 3557, 3559, 3571, 3581, 3583, 3593,
		3607, 3613, 3617, 3623, 3631, 3637, 3643, 3659, 3671, 3673, 3677, 3691, 3697, 3701, 3709,
		3719, 3727, 3733, 3739, 3761, 3767, 3769, 3779, 3793, 3797, 3803, 3821, 3823, 3833, 3847,
		3851, 3853, 3863, 3877, 3881, 3889, 3907, 3911, 3917, 3919, 3923, 3929, 3931, 3943, 3947,
		3967, 3989, 4001, 4003, 4007, 4013, 4019, 4021, 4027, 4049, 4051, 4057, 4073, 4079, 4091,
		4093, 4099, 4111, 4127, 4129, 4133, 4139, 4153, 4157, 4159, 4177, 4201, 4211, 4217, 4219,
		4229, 4231, 4241, 4243, 4253, 4259, 4261, 4271, 4273, 4283, 4289, 4297, 4327, 4337, 4339,
		4349, 4357, 4363, 4373, 4391, 4397, 4409, 4421, 4423, 4441, 4447, 4451, 4457, 4463, 4481,
		4483, 4493, 4507, 4513, 4517, 4519, 4523, 4547, 4549, 4561, 4567, 4583, 4591, 4597, 4603,
		4621, 4637, 4639, 4643, 4649, 4651, 4657, 4663, 4673, 4679, 4691, 4703, 4721, 4723, 4729,
		4733, 4751, 4759, 4783, 4787, 4789, 4793, 4799, 4801, 4813, 4817, 4831, 4861, 4871, 4877,
		4889, 4903, 4909, 4919, 4931, 4933, 4937, 4943, 4951, 4957, 4967, 4969, 4973, 4987, 4993,
		4999, 5003, 5009, 5011, 5021, 5023, 5039, 5051, 5059, 5077, 5081, 5087, 5099, 5101, 5107,
		5113, 5119, 5147, 5153, 5167, 5171, 5179, 5189, 5197, 5209, 5227, 5231, 5233, 5237, 5261,
		5273, 5279, 5281, 5297, 5303, 5309, 5323, 5333, 5347, 5351, 5381, 5387, 5393, 5399, 5407,
		5413, 5417, 5419, 5431, 5437, 5441, 5443, 5449, 5471, 5477, 5479, 5483, 5501, 5503, 5507,
		5519, 5521, 5527, 5531, 5557, 5563, 5569, 5573, 5581, 5591, 5623, 5639, 5641, 5647, 5651,
		5653, 5657, 5659, 5669, 5683, 5689, 5693, 5701, 5711, 5717, 5737, 5741, 5743, 5749, 5779,
		5783, 5791, 5801, 5807, 5813, 5821, 5827, 5839, 5843, 5849, 5851, 5857, 5861, 5867, 5869,
		5879, 5881, 5897, 5903, 5923, 5927, 5939, 5953, 5981, 5987, 6007, 6011, 6029, 6037, 6043,
		6047, 6053, 6067, 6073, 6079, 6089, 6091, 6101, 6113, 6121, 6131, 6133, 6143, 6151, 6163,
		6173, 6197, 6199, 6203, 6211, 6217, 6221, 6229, 6247, 6257, 6263, 6269, 6271, 6277, 6287,
		6299, 6301, 6311, 6317, 6323, 6329, 6337, 6343, 6353, 6359, 6361, 6367, 6373, 6379, 6389,
		6397, 6421, 6427, 6449, 6451, 6469, 6473, 6481, 6491, 6521, 6529, 6547, 6551, 6553, 6563,
		6569, 6571, 6577, 6581, 6599, 6607, 6619, 6637, 6653, 6659, 6661, 6673, 6679, 6689, 6691,
		6701, 6703, 6709, 6719, 6733, 6737, 6761, 6763, 6779, 6781, 6791, 6793, 6803, 6823, 6827,
		6829, 6833, 6841, 6857, 6863, 6869, 6871, 6883, 6899, 6907, 6911, 6917, 6947, 6949, 6959,
		6961, 6967, 6971, 6977, 6983, 6991, 6997, 7001, 7013, 7019, 7027, 7039, 7043, 7057, 7069,
		7079, 7103, 7109, 7121, 7127, 7129, 7151, 7159, 7177, 7187, 7193, 7207, 7211, 7213, 7219,
		7229, 7237, 7243, 7247, 7253, 7283, 7297, 7307, 7309, 7321, 7331, 7333, 7349, 7351, 7369,
		7393, 7411, 7417, 7433, 7451, 7457, 7459, 7477, 7481, 7487, 7489, 7499, 7507, 7517, 7523,
		7529, 7537, 7541, 7547, 7549, 7559, 7561, 7573, 7577, 7583, 7589, 7591, 7603, 7607, 7621,
		7639, 7643, 7649, 7669, 7673, 7681, 7687, 7691, 7699, 7703, 7717, 7723, 7727, 7741, 7753,
		7757, 7759, 7789, 7793, 7817, 7823, 7829, 7841, 7853, 7867, 7873, 7877, 7879, 7883, 7901,
		7907, 7919}
}

func main() {
	//will take input in the form of project numbers and run the code associated with that project.
	funcs := map[string]interface{}{
		"1": Problem1, "2": Problem2, "3": Problem3, "4": Problem4,
		"5": Problem5, "6": Problem6, "7": Problem7, "8": Problem8,
		"9": Problem9, "10": Problem10, "11": Problem11, "12": Problem12,
		"13": Problem13, "14": Problem14, "15": Problem15, "16": Problem16,
		"17": Problem17, "18": Problem18, "19": Problem19, "20": Problem20,
		"21": Problem21, "22": Problem22, "23": Problem23, "24": Problem24,
		"25": Problem25, "26": Problem26, "67": Problem67}
	fmt.Println("Working on a maximum of ", runtime.GOMAXPROCS(0), " CPUs")
	for {
		var choice string
		fmt.Println("Which project would you like to run? [1-26,67], 0 for quit")
		fmt.Scanln(&choice)
		if choice == "0" {
			break
		}
		//in := make([]reflect.Value, 1)
		//reflect.ValueOf().MethodByName(strconv.Itoa(choice)).Call(in)
		CallByName(funcs, choice)
	}

	t := time.Now()
	//Do work
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")

}

//CallByName calls a function through reflection by name
//need to add error handling in there.
func CallByName(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

/*Problem1 If we list all the natural numbers below 10 that are multiples of
//3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 Find the sum of all the multiples of 3 or 5 below 1000.
*/
func Problem1() int {
	var sum = 0
	t := time.Now()
	for i := 0; i < 1000; i++ {

		switch {
		case i%3 == 0:
			sum = sum + i
		case i%5 == 0:
			sum = sum + i
		}
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(sum)
	return sum
}

/*Problem2 Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.*/
func Problem2() int {
	sum := 0
	t := time.Now()
	for i := 1; i < 33; i++ {
		fibNum := fibProblem2(i)
		if fibNum%2 == 0 {
			sum = sum + fibNum
		}
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(sum)
	return sum
}
func fibProblem2(iteration int) (sum int) {
	if iteration == 1 {
		return 1
	}
	if iteration == 2 {
		return 2
	}
	a := fibProblem2(iteration - 2)
	b := fibProblem2(iteration - 1)
	return a + b
}

/*Problem3 The prime factors of 13195 are 5, 7, 13 and 29.
  What is the largest prime factor of the number 600851475143 ?
*/
func Problem3() int {

	num := 600851475143
	largestfactor := 0

	t := time.Now()
	for _, val := range primesBelow8000 {
		//fmt.Print(val, ",")
		if val > 0 {
			if num%val == 0 {
				largestfactor = val
			}
		}
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println("Largest factor is: ", largestfactor)
	return largestfactor
}

/*Problem4 A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.*/
func Problem4() int {
	biggest := 0
	bigI := 0
	bigJ := 0
	t := time.Now()
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			str := strconv.Itoa(i * j)
			if str == reverseString(str) {
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

	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	return bigI * bigJ
}

/*Problem5 2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder. What is the smallest positive number that is
evenly divisible by all of the numbers from 1 to 20? */
func Problem5() int {
	fmt.Println("Working on it, this will take 5-6 seconds.")
	good := true
	t := time.Now()
	for i := 1; i < 10000000000; i++ {
		for j := 1; j < 21; j++ {
			if i%j != 0 {
				good = false
			}
			if !good {
				break
			}
		}
		if good {
			d := time.Since(t)
			fmt.Println("Completed in ", d.Seconds(), "seconds")
			fmt.Println(i)
			return i
		}
		good = true
	}
	return 0
}

/*Problem6 Sum of squares/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func Problem6() int {
	t := time.Now()
	meh := squareSumProblem6(100) - sumSquareProblem6(100)
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(meh)
	return meh
}

func sumSquareProblem6(i int) (val int) {
	if i > 1 {
		val = (i * i) + sumSquareProblem6(i-1)
	} else {
		val = 1
	}
	return val
}
func squareSumProblem6(i int) (val int) {

	for j := i; j > 0; j-- {
		val = val + j
	}
	val = val * val
	return val
}

/*Problem7 By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13. What is the 10 001st prime number?*/
func Problem7() {
	count := 0
	t := time.Now()
	for i := 2; i < 1<<31-1; i++ {
		if isPrime(i) {
			count++
			if count == 10001 {
				d := time.Since(t)
				fmt.Println("Completed in ", d.Seconds(), "seconds")
				fmt.Println(i)
				return
			}
		}
	}
}

/*Problem8 The four adjacent digits in the 1000-digit number that have the
greatest product are 9 × 9 × 8 × 9 = 5832.
73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450
Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
*/
func Problem8() {
	rawdigits := `7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450`
	digits := []byte(rawdigits)
	maxsum := 0
	t := time.Now()
	for i := 13; i < len(digits); i++ {
		sum := 1
		products := digits[i-13 : i]
		for j := 0; j < len(products); j++ {
			c := string(products[j])
			num, _ := strconv.Atoi(c)
			sum = sum * num
		}
		if sum > maxsum {
			maxsum = sum
		}
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(maxsum)
}

/*Problem9 A Pythagorean triplet is a set of three natural numbers, a < b < c,
for which,a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product a*b*c.
A well known triple is 8,15,17 which adds up to 40 which is 1/25 of 1,000
I'll just multiply each of 8,15 and 17 by 25
*/
func Problem9() {
	t := time.Now()
	sum := (8 * 25) * (15 * 25) * (17 * 25)
	d := time.Since(t)
	fmt.Println("run took ", d.Seconds(), "seconds")
	fmt.Println(sum)
}

/*Problem10 The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
func Problem10() {
	t := time.Now()
	sum := 0
	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum = sum + i
		}
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(sum)
}

/*Problem11 In the 20×20 grid below, four numbers along a diagonal
line have been marked in red.
08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48
The product of these numbers is 26 × 63 × 78 × 14 = 1788696.
What is the greatest product of four adjacent numbers in the same direction
(up, down, left, right, or diagonally) in the 20×20 grid?
*/
func Problem11() {
	max := 0
	//Load Problem Data
	grid := make([][]int, 20)
	grid[0] = []int{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8}
	grid[1] = []int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0}
	grid[2] = []int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65}
	grid[3] = []int{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91}
	grid[4] = []int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80}
	grid[5] = []int{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50}
	grid[6] = []int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70}
	grid[7] = []int{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21}
	grid[8] = []int{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72}
	grid[9] = []int{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95}
	grid[10] = []int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92}
	grid[11] = []int{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57}
	grid[12] = []int{86, 56, 0, 48, 35, 71, 89, 7, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58}
	grid[13] = []int{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40}
	grid[14] = []int{04, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66}
	grid[15] = []int{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69}
	grid[16] = []int{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36}
	grid[17] = []int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16}
	grid[18] = []int{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54}
	grid[19] = []int{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48}
	//Ok time to start some work.
	t := time.Now()
	//Check horizontal Products
	for y := 0; y < 20; y++ {
		for x := 0; x < 17; x++ {
			result := grid[y][x] * grid[y][x+1] * grid[y][x+2] * grid[y][x+3]
			if max < result {
				//				fmt.Println("horizontal", y, ",", x, " is ", grid[y][x], grid[y][x+1], grid[y][x+2], grid[y][x+3], "=", result)
				max = result
			}
		}
	}
	//Check vertical Products
	for x := 0; x < 20; x++ {
		for y := 0; y < 17; y++ {
			result := grid[y][x] * grid[y+1][x] * grid[y+2][x] * grid[y+3][x]
			if max < result {
				//				fmt.Println("vertical", y, ",", x, " is ", grid[y][x], grid[y+1][x], grid[y+2][x], grid[y+3][x], "=", result)
				max = result
			}
		}
	}
	//Check right->left down-> up Products Products
	for y := 19; y > 2; y-- {
		for x := 0; x < 17; x++ {
			result := grid[y][x] * grid[y-1][x+1] * grid[y-2][x+2] * grid[y-3][x+3]
			if max < result {
				//				fmt.Println("diag up", y, ",", x, " is ", grid[y][x], grid[y-1][x+1], grid[y-2][x+2], grid[y-3][x+3], "=", result)
				max = result
			}
		}
	}
	//Check left->right up->down Products
	for y := 0; y < 17; y++ {
		for x := 0; x < 17; x++ {
			result := grid[y][x] * grid[y+1][x+1] * grid[y+2][x+2] * grid[y+3][x+3]
			if max < result {
				//				fmt.Println("diag down", y, ",", x, " is ", grid[y][x], grid[y+1][x+1], grid[y+2][x+2], grid[y+3][x+3], "=", result)
				max = result
			}
		}
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(max)
}

/*Problem12 The sequence of triangle numbers is generated by adding the natural
numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.
The first ten terms would be:
1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
Let us list the factors of the first seven triangle numbers:
     1: 1
     3: 1,3
     6: 1,2,3,6
    10: 1,2,5,10
    15: 1,3,5,15
    21: 1,3,7,21
    28: 1,2,4,7,14,28
We can see that 28 is the first triangle number to have over five divisors.
What is the value of the first triangle number to have over five hundred divisors?
*/
func Problem12() {
	t := time.Now()
	triNum := 0
	for i := 1; i < 15000; i++ {
		triNum = triNum + i
		if problem12NumDivisors(triNum) > 500 {
			d := time.Since(t)
			fmt.Println("Completed in ", d.Seconds(), "seconds")
			fmt.Println(triNum)
			return
		}
	}
}
func problem12NumDivisors(num int) (div int) {
	div = 0
	sqrt := int(math.Sqrt(float64(num)))
	if num%sqrt == 0 {
		div = 1
	}
	for i := sqrt - 1 + 1; i >= 1; i-- {
		if num%i == 0 {
			//		fmt.Printf("%d,", i)
			div = div + 2 //since divisors are always paired.
		}
	}
	//	fmt.Println(div)
	return div
}

/*Problem13 Copy and paste the below above here and rename the function
 */
func Problem13() {
	sum := big.NewInt(0)
	line := big.NewInt(0)
	inFile, _ := os.Open("problem13data.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	t := time.Now()
	for scanner.Scan() {
		line.SetString(scanner.Text(), 10)
		sum.Add(sum, line)
		//fmt.Println(scanner.Text())
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(sum)
}

/*Problem14 The following iterative sequence is defined for the set of positive
integers:
n → n/2 (n is even)
n → 3n + 1 (n is odd)
Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?
*/
func Problem14() {
	maxterms := 0
	maxtermsnum := 0
	t := time.Now()

	//start high, work on odd numbers only since they grow first then descend they will always be longer
	for i := 999999; i > 800000; i = i - 2 {
		num := i
		terms := 1
		for {
			if num%2 == 0 {
				num = num / 2
			} else {
				num = (num * 3) + 1
			}
			terms++
			if num == 1 {
				break
			}
		}
		if terms > maxterms {
			maxterms = terms
			maxtermsnum = i
		}
	}

	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(maxtermsnum, ",", maxterms)
}

/*Problem15 Starting in the top left corner of a 2×2 grid, and only being able
to move to the right and down, there are exactly 6 routes to the bottom
right corner. How many such routes are there through a 20×20 grid?
*/
func Problem15() {
	//This is a Schröder Sequence (2n)!/(n!)^2
	t := time.Now()
	//one := big.NewInt(1)
	//forty := big.NewInt(40)
	nom := big.NewInt(0)
	nom.MulRange(1, 40)
	fmt.Println(nom)
	denom := big.NewInt(0)
	denom.MulRange(1, 20)
	denom.Mul(denom, denom)
	fmt.Println(denom)
	answer := big.NewInt(0)
	answer.Div(nom, denom)
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(answer)
}

/*Problem16 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
What is the sum of the digits of the number 2^1000?
*/
func Problem16() {
	two := big.NewInt(2)
	thousand := big.NewInt(1000)
	powerSum := big.NewInt(0)
	t := time.Now()
	powerSum.Exp(two, thousand, powerSum)
	digits := powerSum.String()
	sum := 0
	for i := 0; i < len(digits); i++ {
		temp := rune(digits[i])
		sum += (int(temp) - '0')
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(sum)
}

/*Problem17 is a pain in the ass
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
My solution is a terrible way to do this.
*/
func Problem17() {
	m := map[int]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven",
		8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve",
		13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen",
		17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
		30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy",
		80: "eighty", 90: "ninety",
		100: "onehundredand", 200: "twohundredand", 300: "threehundredand", 400: "fourhundredand",
		500: "fivehundredand", 600: "sixhundredand", 700: "sevenhundredand", 800: "eighthundredand",
		900: "ninehundredand", 1000: "onethousand"}
	l := 0
	s := ""
	//do problem setup here
	t := time.Now()
	for i := 1; i < 1001; i++ {
		if i == 1000 {
			l = l + len(m[i])
			fmt.Println(m[i])
			continue
		}
		if i > 9 && i < 20 {
			l = l + len(m[i])
			fmt.Println(m[i])
			continue
		}
		//Convert to String
		s = strconv.Itoa(i)
		if len(s) == 3 {
			if i%100 == 0 {
				l = l - 3
			}
			n, _ := strconv.Atoi(string(s[0]))
			l = l + len(m[n*100])
			fmt.Print(m[n*100])
			if i%100 < 20 {
				l = l + len(m[i%100])
				fmt.Println(m[i%100])
				continue
			}
			if string(s[1]) != "0" {
				n, _ = strconv.Atoi(string(s[1]))
				l = l + len(m[n*10])
				fmt.Print(m[n*10])
			}
			if string(s[2]) != "0" {
				n, _ = strconv.Atoi(string(s[2]))
				l = l + len(m[n])
				fmt.Println(m[n])
			} else {
				fmt.Println("")
			}
			continue
		}
		if len(s) == 2 {

			n, _ := strconv.Atoi(string(s[0]))
			l = l + len(m[n*10])
			fmt.Print(m[n*10])
			if string(s[1]) != "0" {
				n, _ = strconv.Atoi(string(s[1]))
				l = l + len(m[n])
				fmt.Println(m[n])
			} else {
				fmt.Println("")
			}
			continue
		}
		n, _ := strconv.Atoi(string(s[0]))
		l = l + len(m[n])
		//fmt.Println(m[n], " = ", len(m[n]), "l = ", l)
		fmt.Println(m[n])
	}
	d := time.Since(t)
	fmt.Println("\n", l)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}

/*Problem18 is defined as...
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
   3
  7 4
 2 4 6
8 5 9 3
That is, 3 + 7 + 4 + 9 = 23.
Find the maximum total from top to bottom of the triangle below:

NOTE: As there are only 16384 routes, it is possible to solve this problem by
trying every route. However, Problem 67, is the same challenge with a triangle
containing one-hundred rows; it cannot be solved by brute force, and requires a
clever method! ;o)
my NOTE: I brute forced this one, screw #67 for now*/
func Problem18() {
	//do problem setup here
	lines := make([][]int, 15)
	c := make(chan int, len(lines))
	inFile, _ := os.Open("problem18data.csv")
	max := 0
	defer inFile.Close()
	csvReader := csv.NewReader(bufio.NewReader(inFile))
	csvReader.FieldsPerRecord = -1 //Allows csv file to have variable record lengths
	i := 0
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(line)
		lines[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			num, err := strconv.Atoi(line[j])
			if err != nil {
				fmt.Println(err)
			}
			lines[i][j] = num
		}
		i++
	}
	t := time.Now()
	for j := 0; j < len(lines[i-1]); j++ {
		problem18Sum(lines, i-1, j, 0, c)
	}
	count := 0
	for {
		num := <-c
		if num > max {
			max = num

		}
		count++
		if count >= 16384 {
			fmt.Println(max)
			break
		}

	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}
func problem18Sum(tree [][]int, i, j, sumSoFar int, c chan int) {
	sumSoFar = sumSoFar + tree[i][j]
	if i == 0 { //made the top of the tree
		c <- sumSoFar
		return
	}
	if i == 1 { //only call once for row 0
		go problem18Sum(tree, i-1, 0, sumSoFar, c)
		return
	}
	if i == j { //we're at the right side, only call row above index -1
		go problem18Sum(tree, i-1, j-1, sumSoFar, c)
		return
	}
	if j == 0 { //we're at the left side, only call row above
		go problem18Sum(tree, i-1, j, sumSoFar, c)
		return
	}
	//somewhere in the middle of the row, call the one above and to the side
	go problem18Sum(tree, i-1, j, sumSoFar, c)
	go problem18Sum(tree, i-1, j-1, sumSoFar, c)
}

/*Problem19 is defined as:
You are given the following information, but you may prefer to do some research for yourself.

    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
    A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

Doing some searching on date math apparently the Gaussian formula for day of the week is easy math.
*/
func Problem19() {
	num := 0
	t := time.Now()
	for year := 1901; year < 2001; year++ {
		for month := 1; month < 13; month++ {
			if problem19DayOfTheWeek(year, month, 1) == 0 {
				num++
			}
		}
	}
	d := time.Since(t)
	fmt.Println(num)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}

func problem19DayOfTheWeek(year, month, day int) int {
	d := day
	m := (month-1)%12 + 1
	Y := 0
	if m > 10 {
		Y = year - 1
	} else {
		Y = year
	}
	y := Y % 100
	c := (Y - (Y % 100)) / 100
	w := (d + int(math.Floor(float64(2.6)*float64(m)-float64(0.2))) + y + int(math.Floor(float64(y/4))) + int(math.Floor(float64(c/4))) - 2*c) % 7
	return w
}

/*Problem20 is described as
n! means n × (n − 1) × ... × 3 × 2 × 1
For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!
ProblemXX Copy and paste the below above here and rename the function
*/
func Problem20() {
	//do problem setup here
	//This is a Schröder Sequence (2n)!/(n!)^2
	bigNum := big.NewInt(0)
	num := 0
	t := time.Now()
	bigNum.MulRange(1, 100)
	s := bigNum.String()
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(string(s[i]))
		num = num + n
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(num)
}

/*Problem21 is defined as


Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and
each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55
and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and
142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/
func Problem21() {
	//do problem setup here
	//list := make([]int, 10000) //assume all numbers are amicable and make
	num := 0
	t := time.Now()
	for i := 1; i < 10001; i++ {
		if problem21IsAmicable(i) {
			num = num + i
		}
	}
	d := time.Since(t)
	fmt.Println(num)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}
func problem21SumDiv(num int) int {
	sum := 1
	for j := 2; j <= int(num/2); j++ {
		if num%j == 0 {
			sum = sum + j
		}
	}
	return sum
}
func problem21IsAmicable(num int) bool {
	sum := problem21SumDiv(num)
	if num != sum {
		aSum := problem21SumDiv(sum)
		if aSum == num {
			fmt.Println(num, "'s divisor sum is ", sum, "'s divisor sum is ", aSum)
			return true
		}
	} else {
		fmt.Println(num, "'s divisor sum is ", sum, "!!!")
	}
	return false
}

/*Problem22 is defined as
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
containing over five-thousand first names, begin by sorting it into alphabetical
order. Then working out the alphabetical value for each name, multiply this
value by its alphabetical position in the list to obtain a name score.
For example, when the list is sorted into alphabetical order, COLIN, which is
worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would
obtain a score of 938 × 53 = 49714.
What is the total of all the name scores in the file?

NOTE:I pre-sorted the file since it's easier to do in another application than
resort it each time.
*/
func Problem22() {
	//do problem setup here
	m := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9,
		"J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19,
		"T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26}
	//	names := make([]string, 5163)
	inFile, _ := os.Open("problem22_names-sorted.csv")
	defer inFile.Close()
	csvReader := csv.NewReader(bufio.NewReader(inFile))
	csvReader.FieldsPerRecord = 1
	sum := 0
	i := 1
	t := time.Now()
	for {
		name, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		nameSum := 0
		for j := 0; j < len(name[0]); j++ {
			nameSum = nameSum + m[string(name[0][j])]
		}
		sum = sum + i*nameSum
		i++
	}
	d := time.Since(t)
	fmt.Println(sum)
	fmt.Println("Completed in ", d.Seconds(), "seconds")

}

/*Problem23 is
A perfect number is a number for which the sum of its proper divisors is exactly
equal to the number. For example, the sum of the proper divisors of 28 would be
1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
A number n is called deficient if the sum of its proper divisors is less than n
and it is called abundant if this sum exceeds n.
As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
number that can be written as the sum of two abundant numbers is 24.
By mathematical analysis, it can be shown that all integers greater than 28123
can be written as the sum of two abundant numbers. However, this upper limit
cannot be reduced any further by analysis even though it is known that the
greatest number that cannot be expressed as the sum of two abundant numbers is
less than this limit.
Find the sum of all the positive integers which cannot be written as the sum of
two abundant numbers.
NOTE: Wolfram Alpha has a value of 20161 as the last number that can't be made
with abundant numbers.
*/
func Problem23() {
	//do problem setup here
	abList := make([]int, 4994)
	numList := make(map[int]bool)
	count := 0
	sum := 0
	//prime the integer list
	for i := 1; i < 20162; i++ {
		numList[i] = true
	}

	t := time.Now()
	for i := 12; i <= 20161; i++ {
		if problem23IsAbundant(i) {
			//	fmt.Println(i)
			abList[count] = i
			count++
		}
	}
	fmt.Println("found ", count)
	//now loop through the abList adding numbers and updating numList to false
	for i := 0; i < len(abList); i++ {
		for j := 0; j < len(abList); j++ {
			n := abList[i] + abList[j]
			if n < 20162 {
				numList[n] = false
			} else {
				break
			}
		}
	}
	fmt.Println("finished marking")
	for i := 1; i < 20162; i++ {
		if numList[i] {
			sum = sum + i
		}
	}

	d := time.Since(t)
	fmt.Println(sum)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}
func problem23IsAbundant(num int) bool {
	sum := problem21SumDiv(num)
	if num < sum {
		return true
	}
	return false
}

/*Problem24 is defined as
A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are
listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:
012   021   102   120   201   210
What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

*/
func Problem24() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := time.Now()
	stop := false
	for i := 1; i < 1000000; i++ {
		nums, stop = problem24NextIteration(nums)
		if !stop {
			fmt.Println("stopped at ", i)
			break
		}
	}
	d := time.Since(t)
	fmt.Println(nums)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}
func problem24NextIteration(list []int) ([]int, bool) {
	// Find longest non-increasing suffix
	nums := list
	i := len(nums) - 1
	for {
		if i == 0 {
			break
		}
		if nums[i-1] < nums[i] {
			break
		}
		i--
	}
	if i <= 0 {
		return list, false
	}
	// Let num[i - 1] be the pivot
	// Find rightmost element that exceeds the pivot
	j := len(nums) - 1
	for {
		if nums[j] > nums[i-1] {
			break
		}
		j--
	}
	// Now the value nums[j] will become the new pivot
	// Assertion: j >= i
	// Swap the pivot with j
	temp := nums[i-1]
	nums[i-1] = nums[j]
	nums[j] = temp
	// Reverse the suffix
	j = len(nums) - 1
	for {
		if i >= j {
			break
		}
		temp = nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
	// Successfully computed the next permutation
	return nums, true
}

/*Problem25 is defined as
The Fibonacci sequence is defined by the recurrence relation:
    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:
    F1 = 1,F2 = 1,F3 = 2,F4 = 3,F5 = 5,F6 = 8,F7 = 13,F8 = 21,F9 = 34,F10 = 55,
    F11 = 89,F12 = 144
The 12th term, F12, is the first term to contain three digits.
What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/
func Problem25() {
	num := big.NewInt(1)
	t1 := big.NewInt(1)
	t2 := big.NewInt(1)
	index := 2
	t := time.Now()
	for {
		if index < 10 {
			fmt.Println("before", t1.String(), "+", t2.String(), " = ", num.String())
		}
		num = num.Add(t1, t2)
		if index < 10 {
			fmt.Println("after", t1.String(), "+", t2.String(), " = ", num.String())
		}
		t1.Set(t2)
		t2.Set(num)
		index++
		if len(num.String()) >= 1000 {
			break
		}
	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
	fmt.Println(index)
}

/*Problem26 is defined as
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:
    1/2	= 	0.5
    1/3	= 	0.(3)
    1/4	= 	0.25
    1/5	= 	0.2
    1/6	= 	0.1(6)
    1/7	= 	0.(142857)
    1/8	= 	0.125
    1/9	= 	0.(1)
    1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.
Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/
func Problem26() {
	//do problem setup here
	const prec = 10000 * 8 //measured in bits
	max := 0
	t := time.Now()
	for i := 1; i < 1000; i = i + 2 { //Start high, will likely be around 987 or something
		one := new(big.Float).SetPrec(prec).SetInt64(1)
		x := new(big.Float).SetPrec(prec).SetFloat64(float64(i))
		num := new(big.Float).SetPrec(prec).SetInt64(0)
		num.Quo(one, x) // t = 1.0 / x_n
		ans := problem26LengthOfRepeats(strings.TrimPrefix(num.Text('f', 10000), "0."))

		if ans > max {
			fmt.Printf("%d, %d\n", i, ans)
			max = ans
		}

	}
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}

//problem26LengthofRepeats finds the longest set of repeating numbers (bytes in this case)
//returns the length when a repeat is found or 0.
func problem26LengthOfRepeats(s string) int {
	for i := 2; i < len(s)/2; i++ {
		if i+i < len(s) {
			if s[0:i] == s[i:i+i] {
				return i
			}
		}

	}
	return 0
}

/*Problem27 Copy and paste the below above here and rename the function
Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000
where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |-4| = 4

Find the product of the coefficients, a and b, for the quadratic expression
that produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/
func Problem27() {
	//Brute force here we come.
	//b has to be a prime to satisfy n=0.   0^2+a*0+b = prime
	t := time.Now()
	for a := -1000; a < 1001; a++ {
		cont := true
		for i := 0; i < len(primesBelow1000); i++ {
			for n := 0; n < 1000; n++ {
				num := n*n + a*n + primesBelow1000[i]
				for j := 0; j < len(primesBelow1000); j++ {
					if n <= primesBelow1000[i] {

					}
				}
			}
		}

	}
	//Do work
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}

/*Problem67 Copy and paste the below above here and rename the function
 */
func Problem67() {
	lines := make([][]int, 100)
	inFile, _ := os.Open("problem67data.csv")
	defer inFile.Close()
	csvReader := csv.NewReader(bufio.NewReader(inFile))
	csvReader.FieldsPerRecord = -1 //Allows csv file to have variable record lengths
	i := 0
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lines[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			num, err := strconv.Atoi(line[j])
			if err != nil {
				fmt.Println(err)
			}
			lines[i][j] = num
		}
		i++
	}
	t := time.Now()

	for i = i - 1; i >= 0; i-- { //work bottom to top
		for j := 0; j < i; j++ {
			if (lines[i][j] + lines[i-1][j]) > (lines[i][j+1] + lines[i-1][j]) {
				lines[i-1][j] = lines[i][j] + lines[i-1][j]
			} else {
				lines[i-1][j] = lines[i][j+1] + lines[i-1][j]
			}
		}
	}
	fmt.Println(lines[0][0])

	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}

/*ProblemXX Copy and paste the below above here and rename the function
 */
func ProblemXX() {
	//do problem setup here
	t := time.Now()
	//Do work
	d := time.Since(t)
	fmt.Println("Completed in ", d.Seconds(), "seconds")
}

//Utility Functions
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//isPrime is my fastest way to find if a number is prime
func isPrime(num int) bool {
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
