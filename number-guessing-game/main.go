package main

import (
	"fmt"
	"number-guessing-game/utilities"
)

func main() {
	var userNumber uint
	continueGame := "yes"

	utilities.GreetUser()
	
	for {
		if continueGame=="yes" {
			numberToGuess := utilities.GenerateRandomNumber()
			choices := utilities.GetDifficulties()
			diff := utilities.SelectDifficulty(choices)

			utilities.GameSession(numberToGuess, userNumber, diff)
			fmt.Print("\n")
		} else if continueGame=="no"{
			break
		}
		fmt.Println("Another game ? (yes/no)")
		fmt.Scan(&continueGame)
	}
	
}