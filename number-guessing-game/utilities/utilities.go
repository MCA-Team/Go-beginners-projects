package utilities

import "fmt"


type DiffInfos struct {
	Difficulty string
	NumberOfAttempts uint
}


func GetDifficulties() map[uint]DiffInfos {
	choices := make(map[uint]DiffInfos)

	easy := DiffInfos {
		Difficulty: "Easy",
		NumberOfAttempts: 10,
	}
	medium := DiffInfos {
		Difficulty: "Medium",
		NumberOfAttempts: 5,
	}
	hard := DiffInfos {
		Difficulty: "Hard",
		NumberOfAttempts: 3,
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
}


func SelectDifficulty(difficulties map[uint]DiffInfos) uint {
	var choice uint

	for {
		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Printf("Enter your choice: ")
		fmt.Scan(&choice)
		if (choice < 1) || (choice > 3) {
			fmt.Println("Choose a correct number")
			continue
		} else {
			fmt.Printf("Great! You have selected the %v difficulty level.\n", difficulties[choice].Difficulty)
			break
		}
	}
	return choice
}