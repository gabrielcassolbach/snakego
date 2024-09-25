package game 

import (
	"snakego/snake"
	"snakego/board"
	"fmt"
)

type Game struct {
	board *board.GameMap 
	snke *snake.Snake
}

func NewGame() *Game {
	gme := new(Game)
	gme.board = board.CreateMap(50, 50)
	gme.snke = snake.NewSnake()
	return gme
}

func (game *Game) PrintGame() {
	fmt.Print(game.board)	
}
