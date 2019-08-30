package main

import (
	"fmt"
	"strings"
)

type player struct {
	name  string
	score int
}

type game struct {
	a player
	b player
}

func (p player) incrPoint() int {
	return p.score + 1
}

func convToScore(point int) string {
	var score string
	switch point {
	case 0:
		score = "Love"
	case 1:
		score = "Fifteen"
	case 2:
		score = "Thirty"
	case 3:
		score = "Forty"
	}
	return score
}

func (g game) currentScore() (msg string, gameOn bool) {
	var playerWithHigherPoint, playerWithLowerPoint string
	var pointDiff int

	a := g.a
	b := g.b

	if a.score != b.score {
		if a.score > b.score {
			playerWithHigherPoint = a.name
			playerWithLowerPoint = b.name
			pointDiff = a.score - b.score
		} else {
			playerWithHigherPoint = b.name
			playerWithLowerPoint = a.name
			pointDiff = b.score - a.score
		}

		if (a.score > 3 || b.score > 3) && pointDiff >= 2 {
			msg = "  " + playerWithHigherPoint + " win the game "
			return msg, false
		} else {
			if a.score > 3 || b.score > 3 {
				msg = "Player " + playerWithHigherPoint + " advantage player " + playerWithLowerPoint
			} else {
				msg = convToScore(a.score) + "-" + convToScore(b.score)
			}
		}
	} else {
		if a.score <= 3 {
			msg = convToScore(a.score) + " All"
		}
		if a.score >= 3 {
			msg = "Deuce"
		}
	}
	return msg, true
}

func main() {
	var g game
	var playerA, playerB, msg, gameWinner string
	var gameOn bool

	fmt.Printf("Enter Player A Name: ")
	fmt.Scanln(&playerA)

	fmt.Printf("Enter Player B Name: ")
	fmt.Scanln(&playerB)

	a := player{name: playerA, score: 0}
	b := player{name: playerB, score: 0}

	fmt.Println("================================")
	fmt.Printf("  Welcome %q and %q", a.name, b.name)
	fmt.Println("\n      !!! Game Start !!!")
	fmt.Println("================================")

	g = game{a, b}
	msg, gameOn = g.currentScore()
	fmt.Println(">>> Current score: " + msg)
	fmt.Println("")

	gameOn = true

	for gameOn {
		fmt.Printf("Enter who get the point? (A - %s / B - %s) : ", a.name, b.name)
		fmt.Scanln(&gameWinner)

		switch strings.ToUpper(gameWinner) {
		case "A":
			a.score = a.incrPoint()
		case "B":
			b.score = b.incrPoint()
		default:
			fmt.Println("Error: incorrect input")
		}
		g = game{a, b}
		msg, gameOn = g.currentScore()

		if gameOn {
			fmt.Println(">>> Current score: " + msg)
		} else {
			fmt.Println("================================")
			fmt.Println(msg)
			fmt.Println("      !!! Game Ends !!!")
			fmt.Println("================================")
		}
	}
}
