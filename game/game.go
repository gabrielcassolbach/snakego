package game 

import (
	"snakego/snake"
	"snakego/board"
	"fmt"
)

type Game struct {
	gameMap *board.GameMap 
	snke *snake.Snake
}

func NewGame() *Game {
	gme := new(Game)
	gme.gameMap = board.CreateMap(30, 70)
	gme.snke = snake.NewSnake()
	return gme
}

func (game *Game) UpdateSnake() {

}


func (game *Game) PrintGame() {
	gmap := game.gameMap.GetMap()
	x_cord := game.snke.GetX()
	y_cord := game.snke.GetY()
	
	for i_x, i := range gmap {
		fmt.Println("\t\t\t\t")
		for i_y, j := range i {
			if(i_x == x_cord && y_cord == i_y){
				fmt.Printf(" %s", "*")
			}else {
				fmt.Printf(" %s", j)
			}
		}
	}

}
