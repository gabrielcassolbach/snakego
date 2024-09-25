package main

import (
	"os/exec"
	"os"
	"time"
	"snakego/game"
	"fmt"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func gameLoop(gme *game.Game, delay time.Duration) {
	for {
		fmt.Println("\t\t\t Snake Go")
		gme.PrintGame()
		time.Sleep(delay)
		clearScreen()	
	}
}

func main() {
	gme := game.NewGame()
	go gameLoop(gme,  50 * time.Millisecond)
	for { }
}
