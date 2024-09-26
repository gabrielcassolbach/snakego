package main

import (
	"os/exec"
	"os"
	"time"
	"snakego/game"
	"fmt"
)

const fps = 80

func setBuffer() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func gameLoop (gme *game.Game, delay time.Duration) {
	running := true
	for running {
		running = gme.PrintGame()
		time.Sleep(delay)
		clearScreen()	
	}
}

func main() {
	setBuffer()
	gme := game.NewGame()
	go gameLoop(gme,  fps * time.Millisecond)
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	for gme.IsPlaying() { }
	fmt.Println("\n\tGAME OVER")
}


