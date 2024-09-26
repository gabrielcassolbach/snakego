package game 

import (
	"snakego/snake"
	"snakego/board"
	"snakego/io"
	"fmt"
)

type Game struct {
	gameMap *board.GameMap 
	snke *snake.Snake
	kboard *io.Keyboard
}

func NewGame() *Game {
	gme := new(Game)
	gme.gameMap = board.CreateMap(30, 70)
	gme.snke = snake.NewSnake()
	gme.kboard = io.NewKeyboard()
	return gme
}

func (game *Game) InitKeyLoop() {
	game.kboard.AwaitKeyLoop()
} 

func (game *Game) UpdateVariables() {
	key := game.kboard.GetKey()

	x := game.snke.GetX()
	y := game.snke.GetY()

	if (key == "a") {
		game.snke.SetY(y - 1)	
	}
	if (key == "w") {
		game.snke.SetX(x - 1)
	}	
	if (key == "d") {
		game.snke.SetY(y + 1)
	}
	if (key == "s") {
		game.snke.SetX(x + 1)
	}
}

func (game *Game) GameOver(x int, y int) bool {
	limit_x := game.gameMap.GetSizeX()
	limit_y := game.gameMap.GetSizeY()

	if((x <= 0) || (x >= limit_x) || (y <= 0) || (y >= limit_y)) {	
		return true
	}	

	return false	
}

func (game *Game) PrintGame() bool {
	gmap := game.gameMap.GetMap()
	go game.InitKeyLoop()	
	
	game.UpdateVariables()
	
	x_cord := game.snke.GetX()
	y_cord := game.snke.GetY()

	if(game.GameOver(x_cord, y_cord)){
		return false;
	}
			 
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

	return true
}


