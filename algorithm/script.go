package algorithm

import (
	"bufio"
	"fmt"
	"os"
)
var (
	up, down, right, left string = "up_arrow", "down_arrow", "right_arrow", "left_arrow"
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

type Positions struct {
	UserPosition 		[]int
	TreasurePosition 	[]int	
	ClearPath 	 		[][]int 
	SuggestionPosition 	[]int
}

func (p *Positions) PushClearPath(position []int){
	var isExists bool = false	

	//if already exists dont append 
	for _, item := range p.ClearPath {
		if item[0] == position[0] && item[1] == position[1]{
			isExists = true
		}
	}
	if !isExists {
		p.ClearPath = append(p.ClearPath, position)
	}			
}

func (p *Positions) Movement (toWhere string) {
	var (
		currentPosition, newPosition []int
		fnIncrementX = func(pos []int) []int {
			x,y := pos[0],pos[1]
			if x < 8 {x = x + 1} 
			return []int{x,y}
		}
		fnIncrementY = func(pos []int) []int {
			x,y := pos[0],pos[1]
			if y < 6 {y = y + 1} 			
			return []int{x,y}
		}
		fnDecrementX = func(pos []int) []int {
			x,y := pos[0],pos[1]
			if x > 1 {x = x - 1} 
			return []int{x,y}
		}
		fnDecrementY = func(pos []int) []int {
			x,y := pos[0],pos[1]
			if y > 1 {y = y - 1} 			
			return []int{x,y}
		}
	)
	//please note that we write [0,0] from top left not cartesian diagram!
	currentPosition = p.UserPosition
	if toWhere == up {
		newPosition = fnDecrementY(currentPosition)
	}
	if toWhere == down {
		newPosition = fnIncrementY(currentPosition)
	}
	if toWhere == right {
		newPosition = fnIncrementX(currentPosition)
	}
	if toWhere == left {
		newPosition = fnDecrementX(currentPosition)
	}
	if toWhere == "" {
		newPosition = []int{1, 1}
	}
		
	p.UserPosition = newPosition
	p.PushClearPath(currentPosition)
	//check distance to add probability treasure is nearby 
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
	for y := 1; y <= 6; y++ {
		for x := 1; x <= 8; x++{
			whatToShow = ""

			if len(n.Player.ClearPath) > 0 {
				for _, item := range n.Player.ClearPath {
					if item[0] == x && item[1] == y {
						whatToShow = " . "
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
	
	
}



