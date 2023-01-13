package day02

import (
	"bufio"
	"strings"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
)

const (
	OpponentRock     = "A"
	OpponentPaper    = "B"
	OpponentScissors = "C"

	MyRock     = "X"
	MyPaper    = "Y"
	MyScissors = "Z"

	Loss1 = OpponentRock + " " + MyScissors
	Loss2 = OpponentPaper + " " + MyRock
	Loss3 = OpponentScissors + " " + MyPaper

	Draw1 = OpponentRock + " " + MyRock
	Draw2 = OpponentPaper + " " + MyPaper
	Draw3 = OpponentScissors + " " + MyScissors

	Win1 = OpponentRock + " " + MyPaper
	Win2 = OpponentPaper + " " + MyScissors
	Win3 = OpponentScissors + " " + MyRock
)

//nolint:revive
func SolutionPt1(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)

	var score int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		switch line {
		case Loss1, Loss2, Loss3:
			score += 0
		case Draw1, Draw2, Draw3:
			score += 3
		case Win1, Win2, Win3:
			score += 6
		}
		switch line {
		case Loss2, Draw1, Win3:
			score += 1
		case Loss3, Draw2, Win1:
			score += 2
		case Loss1, Draw3, Win2:
			score += 3
		}
	}

	return solution.IntResult(score)
}

const (
	Rock     = OpponentRock
	Paper    = OpponentPaper
	Scissors = OpponentScissors

	NeedLoss = "X"
	NeedDraw = "Y"
	NeedWin  = "Z"
)

func win(figure string) string {
	switch figure {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		panic("unexcpected")
	}
}

func draw(figure string) string {
	return figure
}

func loss(figure string) string {
	switch figure {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		panic("unexcpected")
	}
}

//nolint:revive
func SolutionPt2(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)

	var score int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		tokens := strings.Split(line, " ")
		oppenentFigure := tokens[0]
		myAction := tokens[1]
		var myFigure string
		switch myAction {
		case NeedLoss:
			myFigure = loss(oppenentFigure)
			score += 0
		case NeedDraw:
			myFigure = draw(oppenentFigure)
			score += 3
		case NeedWin:
			myFigure = win(oppenentFigure)
			score += 6
		}
		switch myFigure {
		case Rock:
			score += 1
		case Paper:
			score += 2
		case Scissors:
			score += 3
		}
	}

	return solution.IntResult(score)
}
