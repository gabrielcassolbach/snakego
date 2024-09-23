package main

import (
	"fmt"
	"board"
)

func main() {
	// generate the map in the screen.	
	fmt.Println("start of the snake game")
	board.gameMap := CreateMap(30, 30)
	board.gameMap.PrintMap()
}






