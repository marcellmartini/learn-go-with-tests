package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPrime(9))
	fmt.Println(primeBetween(1, 100))
}

func isPrime(i int) bool {
	if i < 2 {
		return false
	}

	if i != 2 && i%2 == 0 {
		return false
	}

	for j := 3; j <= int(math.Sqrt(float64(i))); j += 2 {
		if i%j == 0 {
			return false
		}
	}

	return true
}

func primeBetween(num1, num2 int) []int {
	xi := []int{}

	for i := num1; i <= num2; i++ {
		if isPrime(i) {
			xi = append(xi, i)
		}
	}

	return xi
}
