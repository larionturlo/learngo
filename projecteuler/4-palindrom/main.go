package main

import (
	"fmt"
)

const start int = 999
const end int = 99

func main() {
findPolindrome:
	for i := start; i > end; i-- {
		palindrome := generatePalindrome(i, 0, 1)
		for firstNum := start; firstNum > end; firstNum-- {
			secondNum := palindrome / firstNum
			if palindrome%firstNum == 0 && secondNum > end && secondNum <= start {
				fmt.Println("Max palindrom :", palindrome, firstNum, secondNum)
				break findPolindrome
			}
		}
	}
}

func generatePalindrome(leftPart, rightPart, leftPosition int) int {
	leftNumeral := (leftPart / leftPosition) % 10
	rightPart = rightPart*10 + leftNumeral
	leftPosition = leftPosition * 10
	if leftPart/leftPosition == 0 {
		return leftPart*leftPosition + rightPart
	}
	return generatePalindrome(leftPart, rightPart, leftPosition)
}
