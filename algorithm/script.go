package algorithm

import (
	"bufio"
	"fmt"
	"os"
)

type Nucleus struct {	
	Player *Positions
}

func (n *Nucleus) ReadInput() string {		
	r := bufio.NewReader(os.Stdin) 
	char, _, err := r.ReadRune()

	if err != nil {
		fmt.Println(err)
	}	
	if char == 'W' || char == 'w' { return up}
	if char == 'D' || char == 'd' { return right}
	if char == 'S' || char == 's' { return down}
	if char == 'A' || char == 'a' { return left}
	return ""
}


func (n *Nucleus) Show() {
	//8x6 grid 
	//# Obstacle 
	//. clear path / has been visited 
	//X player posistion 
	//$ probably true location --> only generate when check distance is close 
	//O treasure found 
	var (	
		whatToShow string
	)

	//printing maps from left - right - bottom
	for y := 6; y >= 1; y-- {
		for x := 1; x <= 8; x++{
			whatToShow = ""

			if len(n.Player.ClearPath) > 0 {
				for _, itemClearPath := range n.Player.ClearPath {
					if itemClearPath[0] == x && itemClearPath[1] == y {
						whatToShow = " . "
					}
				}
			}
			
			if len(n.Player.TreasureHint) > 0 {
				for _, itemTreasureHint := range n.Player.TreasureHint {
					if itemTreasureHint[0] == x && itemTreasureHint[1] == y {
						whatToShow = " $ "
					}
				}
			}
			
			if x == n.Player.TreasurePosition[0] && y == n.Player.TreasurePosition[1] {
				whatToShow = " O " //treasure
			}
			
			if x == n.Player.UserPosition[0] && y == n.Player.UserPosition[1] {
				
				if whatToShow == " O " {
					whatToShow = " X$" //showing user position
				}else{
					whatToShow = " X " //showing user position
				}							
			} 
			
			if whatToShow == "" {
				whatToShow = " # " //showing obstacles				
			}

			fmt.Printf(whatToShow)			
		}
		fmt.Println("")
	}
	
	if int(n.Player.Distance) <= 2 {
		fmt.Println("hint: You're close!")
		n.Player.GenerateProbabiltyTreasure()
	}
	if int(n.Player.Distance) == 0 {
		fmt.Println("Congratulation you're being hired!")		
	}
}



