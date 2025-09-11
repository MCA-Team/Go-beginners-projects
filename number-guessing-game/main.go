package main

import (
	"fmt"
	"number-guessing-game/utilities"
)

func main() {
	var userNumber uint
	continueGame := "yes"

	utilities.GreetUser()
	choices := utilities.GetDifficulties()
	
	for {
		if continueGame=="yes" {
			numberToGuess := utilities.GenerateRandomNumber()
			t, diff := utilities.SelectDifficulty(choices)

			utilities.GameSession(numberToGuess, userNumber, diff, t, choices)
			fmt.Print("\n")
		} else if continueGame=="no"{
			break
		}
		fmt.Println("Another game ? (yes/no)")
		fmt.Scan(&continueGame)
	}
	
}