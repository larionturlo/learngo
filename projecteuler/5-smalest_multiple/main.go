package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	var uniquePrimeFactors = make(map[int]int)
	for i := 20; i > 1; i-- {
		uniquePrimes := uniquePrimeFactorization(i, 2, make(map[int]int))
		for prime, count := range uniquePrimes {
			if uniquePrimeFactors[prime] < count {
				uniquePrimeFactors[prime] = count
			}
		}
	}
	var lcm = 1.0
	for prime, power := range uniquePrimeFactors {
		lcm = lcm * math.Pow(float64(prime), float64(power))
	}
	elapsed := time.Since(start)
	fmt.Println("Smallest multiple:", int(lcm))
	fmt.Println("Time:", elapsed) // max duration 4 ms, min 1.7 ms
}

func uniquePrimeFactorization(number, factor int, uniquePrime map[int]int) map[int]int {
	mod := number % factor
	if mod > 0 {
		factor++
		return uniquePrimeFactorization(number, factor, uniquePrime)
	}

	uniquePrime[factor]++
	newNumber := number / factor
	if newNumber == 1 {
		return uniquePrime
	}

	return uniquePrimeFactorization(newNumber, factor, uniquePrime)
}
