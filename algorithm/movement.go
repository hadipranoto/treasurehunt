package algorithm

import "math"

var (
	up, down, right, left string = "up_arrow", "down_arrow", "right_arrow", "left_arrow"
)


type Positions struct {
	UserPosition 		[]int
	TreasurePosition 	[]int	
	ClearPath 	 		[][]int 
	SuggestionPosition 	[]int
	Distance 			int
	TreasureHint 	 	[][]int 
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
		newPosition = fnIncrementY(currentPosition)
	}
	if toWhere == down {
		newPosition = fnDecrementY(currentPosition)
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
	p.Distance = p.CalculateMyDistance()
}

func (p *Positions) CalculateMyDistance () int {
	nowX 	:= float64(p.UserPosition[0])
	trueX 	:= float64(p.TreasurePosition[0])
	nowY 	:= float64(p.UserPosition[1])
	trueY 	:= float64(p.TreasurePosition[1])

	nowDistance := math.Abs(nowX - trueX) + math.Abs(nowY - trueY)		
	return int(nowDistance)	
}

func (p *Positions) CalculateDistance (a, b []int) int {
	aX 	:= float64(a[0])
	bX 	:= float64(b[0])
	aY 	:= float64(a[1])
	bY 	:= float64(b[1])

	nowDistance := math.Abs(aX - bX) + math.Abs(aY - bY)
	return int(nowDistance)	
}

func (p *Positions) GenerateProbabiltyTreasure (){
	var (
		allProbabilities [][]int			
	)

	for y := 6; y >= 1; y-- {
		for x := 1; x <= 8; x++{
			if p.CalculateDistance([]int{x,y}, []int{p.TreasurePosition[0], p.TreasurePosition[1]}) <= 2 {
				allProbabilities = append(allProbabilities, []int{x,y})
			}
		}
	}
	
	for i := 0; i <= len(allProbabilities)/2; i++ {
		p.TreasureHint = append(p.TreasureHint, allProbabilities[i])
	}   
}