package main

import (
	"fmt"
	"number-guessing-game/utilities"
)

func main() {
	utilities.GreetUser()
	fmt.Println(utilities.GetDifficulties())
	fmt.Println(utilities.SelectDifficulty(utilities.GetDifficulties()))
}


