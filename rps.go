package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var ch int
var aiCh int
var aiChParse string

func main() {
	cls()
	var in string
	print("Rock Paper Scissors")
	print("s. Start Game")
	print("ex. Exit")
	fmt.Scanln(&in)
	switch in {
	case "s":
		gameStart()
	case "ex":
		os.Exit(0)
	}
}
func print(text string) {
	fmt.Println(text)
}
func gameStart() {
	cls()
	var playerSelect string
	print("Your selection?: (r, p, s)")
	fmt.Scanln(&playerSelect)
	if playerSelect == "r" {
		ch = 0
		logic()
	} else if playerSelect == "p" {
		ch = 1
		logic()
	} else if playerSelect == "s" {
		ch = 2
		logic()
	} else {
		gameStart()
	}
}
func ai() {
	max := 3
	min := 0
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(max-min) + min
	switch val {
	case 0:
		aiCh = 0
	case 1:
		aiCh = 1
	case 2:
		aiCh = 2
	}
}
func logic() {
	cls()
	rand.Seed(time.Now().UnixNano())
	var sn string
	ai()
	fmt.Println("AI Selected: ")
	parseAi()
	print(aiChParse)
	if (aiCh == 0 && ch == 2) || (aiCh == 1 && ch == 0) || (aiCh == 2 && ch == 1) {
		print("You Lose!")
	} else if aiCh == ch {
		print("Tie!")
	} else {
		print("You Win!")
	}
	print("Play again? (y,n)")
	fmt.Scanln(&sn)
	switch sn {
	case "y":
		gameStart()
	case "n":
		os.Exit(0)
	}
}
func cls() {
	rand.Seed(time.Now().UnixNano())
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
func parseAi() {
	switch aiCh {
	case 0:
		aiChParse = "Rock"
	case 1:
		aiChParse = "Paper"
	case 2:
		aiChParse = "Scissors"
	}
}

//Rules
// Scissors beat Paper
// Paper beats Rock
// Rock beats scissors
