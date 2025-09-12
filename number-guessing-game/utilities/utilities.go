// Package utilities provides a bunch of functions and struct type necessecary for the game logic implementation.
package utilities

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// difficultyInfos is a struct type which modelizes what a difficulty level is; whith the difficulty name, its number of attempts and the high score related to the difficulty level.
type difficultyInfos struct {
	difficulty string
	numberOfAttempts int
	highScore int
}


// GetDifficulties return a map where the key is the difficulty ID and the value is a struct type which defines all informations for a difficulty level.
// ID: 1, Difficulty: Easy
// ID: 2, Difficulty: Medium
// ID: 3, Difficulty: Hard
func GetDifficulties() map[uint]difficultyInfos {
	choices := make(map[uint]difficultyInfos)

	easy := difficultyInfos {
		difficulty: "Easy",
		numberOfAttempts: 10,
		highScore: 10,
	}
	medium := difficultyInfos {
		difficulty: "Medium",
		numberOfAttempts: 5,
		highScore: 5,
	}
	hard := difficultyInfos {
		difficulty: "Hard",
		numberOfAttempts: 3,
		highScore: 3,
	}

	choices[1] = easy
	choices[2] = medium
	choices[3] = hard

	return choices
}


// GreetUser welcomes the user in the game interface and explains him/her the game logic.
func GreetUser() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I am thinking of a number between 1 and 100.")
	fmt.Println("You have a defined number of chances to guess the correct number, based on the difficulty level selected.")
	fmt.Print("\n")
}


// SelectDifficulty takes the map of all available difficulties as a parameter, drives the user to make a choice, and returns the selected difficulty level's ID and the related number of attempts.
func SelectDifficulty(difficulties map[uint]difficultyInfos) (uint, int) {
	var choice uint

	for {
		fmt.Println("Please select the difficulty level:")
		for difficultyID, difficultyInfos := range difficulties {
			fmt.Printf("%v. %v (%v chances)\n", difficultyID, difficultyInfos.difficulty ,difficultyInfos.numberOfAttempts)
		}

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
	return choice, difficulties[choice].numberOfAttempts
}


// GenerateRandomNumber returns a random number between 1 and 100, based on math/rand/v2 package.
func GenerateRandomNumber() uint {
	return uint(rand.IntN(100)+1)
}

// GameSession implements the game session logic by considering the random number to guess, the number guessed by the user, the number of attempts, the difficulty ID and the difficulty levels map.
// While the number to guess is not equal to the guessed number by the user and while the maximum number of attempts is not reached, the game continues.
// The game stops when the correct number is guessed by the user or when the maximum number of attempt is reached. 
func GameSession(numberToGuess uint, userNumber uint, numberofAttemps int, difficultyID uint, difficultyLevels map[uint]difficultyInfos) {
	start := time.Now()
	// fmt.Printf("Start time: %v\n", start)

	for i := 0; i < numberofAttemps; i++ {
		// fmt.Println(numberToGuess)
		fmt.Print("\nEnter your guess: ")
		fmt.Scan(&userNumber)
		if userNumber == numberToGuess {
			if i+1 < difficultyLevels[difficultyID].highScore {
				temp := difficultyLevels[difficultyID]
				temp.highScore = i+1
				difficultyLevels[difficultyID] = temp
			}
			fmt.Print("\n##############################\n")
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.\n", i+1)
			fmt.Print("##############################\n")
			break
		} else if userNumber > numberToGuess {
			fmt.Printf("Incorrect! The number is less than %v.\n", userNumber)
		} else if userNumber < numberToGuess {
			fmt.Printf("Incorrect! The number is greater than %v.\n", userNumber)
		}
	}
	elapsedTime := time.Since(start)
	fmt.Print("\n******************************\n")
	fmt.Printf("%v level's high score is: %v\n", difficultyLevels[difficultyID].difficulty, difficultyLevels[difficultyID].highScore)
	fmt.Printf("Game elapsed time: %v\n", elapsedTime)
	fmt.Print("******************************\n")

	if numberToGuess != userNumber {
		fmt.Print("\n##############################\n")
		fmt.Printf("Sorry, you've used all your attempts. The correct number was %v.\n", numberToGuess)
		fmt.Print("##############################\n")
	}
}