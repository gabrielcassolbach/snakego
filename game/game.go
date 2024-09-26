package game 

import (
	"snakego/snake"
	"snakego/board"
	"snakego/io"
	"fmt"
)

const height = 30
const lenght = 70

type Game struct {
	gameMap *board.GameMap 
	snke *snake.Snake
	kboard *io.Keyboard
	playing bool
}

func NewGame() *Game {
	gme := new(Game)
	gme.playing = true
	gme.gameMap = board.CreateMap(height, lenght)
	gme.snke = snake.NewSnake(height/2, lenght/2)
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

	return ((x <= 0) || (x >= limit_x-1) || (y <= 0) || (y >= limit_y-1))
}

func (game *Game) IsPlaying() bool {
	return game.playing
}

func (game *Game) InsertIntoQueue() {
	x := game.snke.GetX()
	y := game.snke.GetY()
	
	game.snke.Enqueue(snake.Point{x, y})			
}

func (game *Game) PrintGame() bool {
	gmap := game.gameMap.GetMap()
	go game.InitKeyLoop()	
	
	game.InsertIntoQueue()
	game.UpdateVariables()
	
	x_cord := game.snke.GetX()
	y_cord := game.snke.GetY()

	if(game.GameOver(x_cord, y_cord)){
		game.playing = false
		return false;
	}
	
	for game.snke.GetQueueSize() > 6 {
		pt := game.snke.QTop()
		gmap[pt.Px][pt.Py] = " "
	}
		
	gmap[x_cord][y_cord] = "*"

	for _, i := range gmap {
		fmt.Println("\t")
		for _, j := range i {
			fmt.Printf(" %s", j)
		}
	}

	return true
}


