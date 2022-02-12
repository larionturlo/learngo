package main

import (
	"fmt"
	"strconv"
)

func main() {
	var enter string
	var needNumber bool = true

	for needNumber {

		fmt.Println("Enter your floating point number:")

		fmt.Scan(&enter)

		trunk, errorParsing := strconv.ParseFloat(enter, 64)

		if errorParsing != nil {
			fmt.Println("Entered wrong number!")
			continue
		}

		needNumber = false

		fmt.Printf("Your truncated number: %d \n", int(trunk))
	}
}
