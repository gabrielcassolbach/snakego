package main

import (
	"fmt"
	"os/exec"
	"os"
	//"snakego/io"
	"time"
	"snakego/board"
	"snakego/snake"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func gameLoop(mp *board.GameMap, snke *snake.Snake,  delay time.Duration) {
	for {
		fmt.Println("\t\t\t Snake Go")
		mp.PrintMap()
		time.Sleep(delay)
		clearScreen()	
	}
}

func main() {
	snke := snake.NewSnake()  	
	mp := board.CreateMap(30, 30) 	
	go gameLoop(mp, snke,  100 * time.Millisecond)
	for { }
}
