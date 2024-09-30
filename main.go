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

func initialScreen() {
	fmt.Println("\t SNAKEGO");
	fmt.Println("\t\npress a, w, d or s to start the game")
}

func validKey(k string) bool {
	fmt.Println(k)
	return k == "a" || k == "w" || k == "s" || k == "d"
}

func gameLoop (gme *game.Game, delay time.Duration) {
	running := false
	for !running {
		initialScreen()
		key := gme.InitGame()		
		running = validKey(key)
		time.Sleep(delay)
		clearScreen()	
	}
	
	for running {
		running = gme.PrintGame()
		time.Sleep(delay)
		clearScreen()	
	}
}

func main() {
	setBuffer()
	gme := game.NewGame()
	go gme.InitKeyLoop()
	go gameLoop(gme,  fps * time.Millisecond)
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	for gme.IsPlaying() { }
	fmt.Println("\n\tGAME OVER")
}


