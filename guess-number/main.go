package main

import (
	"fmt"
	"math/rand/v2"
)

func absoluteDiff(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func main() {
	randomNumber := rand.IntN(100) + 1
	var guessedNumber int

	fmt.Println("Try to guess the number between 1 and 100!")

	for {
		fmt.Scanln(&guessedNumber)

		if guessedNumber == randomNumber {
			fmt.Printf("You've guessed the number correctly!. It was %v\n", randomNumber)
			break
		}
		diff := absoluteDiff(guessedNumber, randomNumber)

		if diff < 5 {
			fmt.Println("Burning")
		} else if diff < 10 {
			fmt.Println("Hot")
		} else if diff < 20 {
			fmt.Println("Warm")
		} else if diff < 30 {
			fmt.Println("Cold")
		} else {
			fmt.Println("Freezing")
		}
	}
}
