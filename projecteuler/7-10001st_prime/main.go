package main

import (
	"fmt"
	"time"
)

var (
	primes = []int{}
	count  = 1
)

func main() {
	start := time.Now()
	num := 2
	primes = append(primes, num)
	for count < 10001 {
		num++
		if isPrime(num) {
			primes = append(primes, num)
			count++
			if count == 10001 {
				fmt.Println(num)
			}
		}
	}
	//fmt.Println(primes[10000:])
	elapsed := time.Since(start)
	fmt.Println("Time:", elapsed)
}

func isPrime(num int) bool {
	for _, prime := range primes {
		if num%prime == 0 {
			return false
		}
	}
	return true
}
