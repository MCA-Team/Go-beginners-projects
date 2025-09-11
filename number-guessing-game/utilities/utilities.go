package utilities

import (
	"fmt"
	"math/rand/v2"
	"time"
)


type difficultyInfos struct {
	difficulty string
	numberOfAttempts int
	highScore int
}


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


func GreetUser() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I am thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Print("\n")
}


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

func GenerateRandomNumber() uint {
	return uint(rand.IntN(100)+1)
}

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