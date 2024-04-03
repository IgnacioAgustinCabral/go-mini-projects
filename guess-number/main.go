package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	randomNumber := rand.IntN(100) + 1
	counter := 0
	fmt.Println("Try to guess the number between 1 and 100!")
	for {
		var input string
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&input)
		if containsDecimalPoint(input) || containsSpecialCharacters(input) || containsSpace(input) {
			fmt.Println("Please enter a whole number. No letters, special characters or floating numbers allowed")
			continue
		}

		guessedNumber, _ := strconv.Atoi(input)

		if guessedNumber < 1 || guessedNumber > 100 {
			fmt.Println("Entered number must be between 1 and 100")
			continue
		}

		counter++

		if guessedNumber == randomNumber {
			fmt.Printf("You've guessed the number correctly!.It was %v\n", randomNumber)
			fmt.Printf("It took you %v tries\n", counter)
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
func absoluteDiff(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func containsDecimalPoint(input string) bool {
	for _, char := range input {
		if char == '.' {
			return true
		}
	}
	return false
}

func containsSpecialCharacters(input string) bool {
	for _, char := range input {
		if char < '0' || char > '9' {
			return true
		}
	}
	return false
}

func containsSpace(input string) bool {
	for _, char := range input {
		if char == ' ' {
			return true
		}
	}
	return false
}
