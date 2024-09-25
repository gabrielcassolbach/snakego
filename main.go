package main

import (
	"os/exec"
	"os"
	"time"
	"snakego/game"
)

const fps = 100

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func gameLoop(gme *game.Game, delay time.Duration) {
	for {
		gme.PrintGame()
		time.Sleep(delay)
		clearScreen()	
	}
}

func main() {
	gme := game.NewGame()
	go gameLoop(gme,  fps * time.Millisecond)
	for { }
}
