package main

import (
	"fmt"
	"os"
	"strings"
)

type Shape struct {
	Name string
	WinAgainst string
	LoseAgainst string
	Score int
}

func main() {
	rock := Shape{
		Name: "Rock",
		WinAgainst: "Scissors",
		LoseAgainst: "Paper",
		Score: 1,
	}
	
	paper := Shape{
		Name: "Paper",
		WinAgainst: "Rock",
		LoseAgainst: "Scissors",
		Score: 2,
	}

	scissors := Shape{
		Name: "Scissors",
		WinAgainst: "Paper",
		LoseAgainst: "Rock",
		Score: 3,
	}

	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	firstPartTotalScore := 0
	firstPartRoundScore := 0
	secondPartTotalScore := 0
	secondPartRoundScore := 0

	for _, line := range lines {
		move := strings.Split(line, " ")
		if len(move) == 2 {
			opponent := move[0]
			me := move[1]
			var(
				myMove Shape
				opMove Shape
			)
			switch opponent {
			case "A":
				opMove = rock
			case "B":
				opMove = paper
			case "C":
				opMove = scissors
			}
			switch me {
			case "X":
				myMove = rock
			case "Y":
				myMove = paper
			case "Z":
				myMove = scissors
			}
			firstPartRoundScore = calculateFirstPart(opMove, myMove)
			secondPartRoundScore = calculateSecondPart(opMove, myMove)
		}
		firstPartTotalScore += firstPartRoundScore
		firstPartRoundScore = 0
		secondPartTotalScore += secondPartRoundScore
		secondPartRoundScore = 0
	}
	fmt.Println(firstPartTotalScore)
	fmt.Println(secondPartTotalScore)
}

func calculateFirstPart(opMove, myMove Shape) (firstPartRoundScore int) {
	matchScore := 0
	if opMove.LoseAgainst == myMove.Name {
		matchScore = 6
	} else if opMove.Name == myMove.Name {
		matchScore = 3
	} else {
		matchScore = 0
	}
	moveScore := myMove.Score
	firstPartRoundScore = matchScore + moveScore
	return firstPartRoundScore
}

func calculateSecondPart(opMove, myMove Shape) (secondPartRoundScore int) {
	matchScore := 0
	moveScore := 0

	switch myMove.Name {
	case "Rock":
		matchScore = 0
		moveScore = mustLoseScore(opMove)
	case "Paper":
		matchScore = 3
		moveScore = mustDrawScore(opMove)
	case "Scissors":
		matchScore = 6
		moveScore = mustWinScore(opMove)
	}
	secondPartRoundScore = matchScore + moveScore
	return secondPartRoundScore
}

func mustWinScore(opMove Shape) (score int) {
	switch opMove.Name {
	case "Rock":
		score = 2
	case "Paper":
		score = 3
	case "Scissors":
		score = 1
	}
	return score
}

func mustDrawScore(opMove Shape) (score int) {
	switch opMove.Name {
	case "Rock":
		score = 1
	case "Paper":
		score = 2
	case "Scissors":
		score = 3
	}
	return score
}

func mustLoseScore(opMove Shape) (score int) {
	switch opMove.Name {
	case "Rock":
		score = 3
	case "Paper":
		score = 1
	case "Scissors":
		score = 2
	}
	return score
}
