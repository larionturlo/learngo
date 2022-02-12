package main

import "fmt"

func main() {
	var str string

	for {
		fmt.Println("Enter a word:")

		fmt.Scan(&str)

		fmt.Println("==========")
		if checkChars([]rune(str)) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
		fmt.Println("==========")

		fmt.Println("For exit press Ctrl+C")
	}
}

func checkChars(chars []rune) bool {
	if len(chars) == 0 {
		return false
	}
	if chars[0] != 'i' && chars[0] != 'I' {
		return false
	}

	lastIndex := len(chars) - 1

	if chars[lastIndex] != 'n' && chars[lastIndex] != 'N' {
		return false
	}

	for i := 1; i < lastIndex; i++ {
		if chars[i] == 'a' || chars[i] == 'A' {
			return true
		}
	}

	return false
}
