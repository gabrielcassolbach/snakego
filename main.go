package main

import (
	"fmt"
	"os/exec"
	"os"
	"snakego/io"
	//"time"
	"snakego/board"
	//"snakego/snake"
)

const fps = 35

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// generate the map in the screen.	
	fmt.Println("\t\t\t Snake Go")
	mp := board.CreateMap(30, 30) 	

	key := io.GetPressedKey()
	for key != "e" {
		go io.GetPressedKey()
		clearScreen()
		mp.PrintMap()			
	}

}


