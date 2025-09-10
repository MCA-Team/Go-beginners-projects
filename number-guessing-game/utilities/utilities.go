package utilities

import (
	"fmt"
	"math/rand/v2"
	"time"
)


type diffInfos struct {
	difficulty string
	numberOfAttempts int
	HighScore int
}


func GetDifficulties() map[uint]diffInfos {
	choices := make(map[uint]diffInfos)

	easy := diffInfos {
		difficulty: "Easy",
		numberOfAttempts: 10,
		HighScore: 0,
	}
	medium := diffInfos {
		difficulty: "Medium",
		numberOfAttempts: 5,
		HighScore: 0,
	}
	hard := diffInfos {
		difficulty: "Hard",
		numberOfAttempts: 3,
		HighScore: 0,
	}

	choices[1] = easy
	choices[2] = medium
	choices[3] = hard

	return choices
}


func GreetUser() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I am thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Print("\n")
}


func SelectDifficulty(difficulties map[uint]diffInfos) int {
	var choice uint

	for {
		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)
		if (choice < 1) || (choice > 3) {
			fmt.Println("Choose a correct number")
			fmt.Print("\n")
			continue
		} else {
			fmt.Printf("\nGreat! You have selected the %v difficulty level.\n", difficulties[choice].difficulty)
			fmt.Println("Let's start the game!")
			break
		}
	}
	return difficulties[choice].numberOfAttempts
}

func GenerateRandomNumber() uint {
	return uint(rand.IntN(100)+1)
}

func GameSession(numberToGuess uint, userNumber uint, diff int) {
	start := time.Now()
	// fmt.Printf("Start time: %v\n", start)

	for i := 0; i < diff; i++ {
		// fmt.Println(numberToGuess)
		fmt.Print("\nEnter your guess: ")
		fmt.Scan(&userNumber)
		if userNumber == numberToGuess {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.\n", i+1)
			break
		} else if userNumber > numberToGuess {
			fmt.Printf("Incorrect! The number is less than %v.\n", userNumber)
		} else if userNumber < numberToGuess {
			fmt.Printf("Incorrect! The number is greater than %v.\n", userNumber)
		}
	}
	elapsedTime := time.Since(start)
	fmt.Print("\n******************************\n")
	fmt.Printf("Game elapsed time: %v\n", elapsedTime)
	fmt.Print("******************************\n")

	if numberToGuess != userNumber {
		fmt.Print("\n##############################\n")
		fmt.Printf("Sorry, you've used all your attempts. The correct number was %v.\n", numberToGuess)
		fmt.Print("##############################\n")
	}
}