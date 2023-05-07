package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("Slayer Game")
	fmt.Println("Starting a new game....")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("--------------------------")

	fmt.Println("(1). Attack Enemy")
	fmt.Println("(2). Heal")

	if specialAttackIsAvailable == true {
		fmt.Println("(3). Special Attack")
	}
}
