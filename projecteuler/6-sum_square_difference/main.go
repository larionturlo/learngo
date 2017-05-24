package main

import (
	"fmt"
	"math"
	"time"
)

var (
	simpleSum    float64
	sumOfSquares float64
)

func main() {
	start := time.Now()
	for i := 1; i <= 100; i++ {
		num := float64(i)
		simpleSum += num
		sumOfSquares += math.Pow(num, 2.0)
	}
	fmt.Println(math.Pow(simpleSum, 2.0) - sumOfSquares) //cast to int
	elapsed := time.Since(start)
	fmt.Println("Time:", elapsed)
}
