package main

import (
	"fmt"

	algo "github.com/hadipranoto/treasurehunt/algorithm"
)


func main (){
	
	
	player := algo.Positions{
		UserPosition: []int{5, 5},
		TreasurePosition: []int{2, 3},				
	}
	nucleus := algo.Nucleus{Player: &player}

	
	for i := 0; i < 100; i++ {			
		message := "#######################################################\n" +
		"# A treasure has been hidden at a location in a 8x6   #\n" +
		"# Guess where it is using [W, S, A, D] as arrow       #\n" +
		"# Author : Hadi 								       #\n" +
		"#######################################################";
		fmt.Println(message)

		nucleus.Show()
		direction := nucleus.ReadInput()		
		nucleus.Player.Movement(direction)
		nucleus.Show()
	}

	
	
}