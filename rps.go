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
var chParse string

func main() {
	cls()
	var in string
	var opt int
	print("Rock Paper Scissors")
	print("s. Start Game")
	print("ex. Exit")
	fmt.Scanln(&in)
	switch in {
	case "s":
		print("One or two player? (1, 2): ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			gameStart()
		case 2:
			twoPlayerGame()
		}

	case "ex":
		os.Exit(0)
	}
}
func twoPlayerGame() {
	cls()
	var p1 string
	var p2 string
	print("Player 1: Your selection?: (r, p, s)")
	fmt.Scanln(&p1)
	cls()
	print("Player 2: Your selection?: (r, p, s)")
	fmt.Scanln(&p2)
	if (p1 == "r" && p2 == "p") || (p1 == "p" && p2 == "s") || (p1 == "s" && p2 == "r") {
		print("You Lose!")
	} else if p1 == p2 {
		print("Tie!")
	} else {
		print("You Win!")
	}
	print("*****************")
	print("Play again? (y,n)")
	var sn string
	fmt.Scanln(&sn)
	switch sn {
	case "y":
		gameStart()
	case "n":
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
	parseAi()
	parsePlayer()
	fmt.Println("AI Selected: ", aiChParse)
	fmt.Println("You Selected: ", chParse)
	if (aiCh == 0 && ch == 2) || (aiCh == 1 && ch == 0) || (aiCh == 2 && ch == 1) {
		print("You Lose!")
	} else if aiCh == ch {
		print("Tie!")
	} else {
		print("You Win!")
	}
	print("*****************")
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
func parsePlayer() {
	switch ch {
	case 0:
		chParse = "Rock"
	case 1:
		chParse = "Paper"
	case 2:
		chParse = "Scissors"
	}
}

//Rules
// Scissors beat Paper
// Paper beats Rock
// Rock beats scissors
