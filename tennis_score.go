package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func main() {
	var a, b, point string
	var playerWithHigherPoint, playerWithLowerPoint string
	var playerApoint string
	var playerBpoint string
	var pointDiff int
	var gameOn bool

	fmt.Printf("Enter Player A Name: ")
	fmt.Scanln(&a)

	fmt.Printf("Enter Player B Name: ")
	fmt.Scanln(&b)

	fmt.Println("================================")
	fmt.Printf("Welcome %q and %q", a, b)
	fmt.Println("\n!!! Game Start !!!")
	fmt.Println("================================")
	gameOn = true

	for gameOn {
		fmt.Printf("Enter game point (0-0) : ")
		fmt.Scanln(&point)

		p := strings.Split(point, "-")

		pointA, err := strconv.Atoi(p[0])
		if err != nil {
			fmt.Println(err)
		}
		pointB, err := strconv.Atoi(p[1])
		if err != nil {
			fmt.Println(err)
		}

		if pointA != pointB {
			if pointA > pointB {
				playerWithHigherPoint = a
				playerWithLowerPoint = b
				pointDiff = pointA - pointB
			} else {
				playerWithHigherPoint = b
				playerWithLowerPoint = a
				pointDiff = pointB - pointA
			}

			if (pointA > 3 || pointB > 3) && pointDiff >= 2 {
				fmt.Println("================================")
				fmt.Println(playerWithHigherPoint + " win the game")
				fmt.Println(" !!! Game End !!! ")
				fmt.Println("================================")
				gameOn = false
			} else {
				if pointA > 3 || pointB > 3 {
					fmt.Println("Player " + playerWithHigherPoint + " advantage player " + playerWithLowerPoint)
				} else {
					playerApoint = convToScore(pointA)
					playerBpoint = convToScore(pointB)
					fmt.Println(playerApoint + "-" + playerBpoint)
				}
			}
		} else {
			if pointA <= 3 {
				playerApoint = convToScore(pointA)
				fmt.Println("All-" + playerApoint)
			}
			if pointA >= 3 {
				fmt.Println("Deuce")
			}
			pointDiff = 0
		}
	}
}
