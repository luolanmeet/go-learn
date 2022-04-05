package main

import (
	"fmt"
	"math/bits"
)

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeSetBits(left, right int) (ans int) {
	for i := left; i <= right; i++ {
		if isPrime(bits.OnesCount(uint(i))) {
			ans++
		}
	}
	return
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
}
