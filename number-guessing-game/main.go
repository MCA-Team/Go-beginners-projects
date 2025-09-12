// Packge main orders and implement the game logic by using the functions and struct types defined in the utilities package
package main

import (
	"fmt"
	"number-guessing-game/utilities"
)

func main() {
	var userNumber uint
	continueGame := "yes"
	difficultyLevels := utilities.GetDifficulties()
	
	for {
		if continueGame=="yes" {
			numberToGuess := utilities.GenerateRandomNumber()
			difficultyID, numberofAttemps := utilities.SelectDifficulty(difficultyLevels)
			utilities.GreetUser()
			utilities.GameSession(numberToGuess, userNumber, numberofAttemps, difficultyID, difficultyLevels)
			fmt.Print("\n")
		} else if continueGame=="no"{
			break
		}
		fmt.Println("Another game ? (yes/no)")
		fmt.Scan(&continueGame)
	}
	
}