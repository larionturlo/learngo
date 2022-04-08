package main

import (
	"fmt"
	"strconv"
)

func sliceFromInput() {
	inputData := new([]byte)

	exit := false
	numbers := make([]int64, 3)
	index := 0

	for !exit {
		fmt.Println("Enter a number or `x` for exit: ")
		fmt.Scan(inputData)

		inputString := string(*inputData)

		if inputString == "x" || inputString == "X" {
			exit = true
			continue
		}

		num, err := strconv.ParseInt(inputString, 0, 0)
		if err != nil {
			fmt.Println("Wrong input data!")
			continue
		}

		if index > len(numbers)-1 {
			numbers = append(numbers, num)
		} else {
			numbers[index] = num
		}

		for i := index; i > 0; i-- {
			if numbers[i] > numbers[i-1] {
				break
			}
			numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
		}
		index++
		fmt.Println(numbers)
	}
}
