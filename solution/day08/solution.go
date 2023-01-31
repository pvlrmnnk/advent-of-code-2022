package day08

import (
	"bufio"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
)

type coords struct {
	x, y int
}

func parseForest(scanner *bufio.Scanner) map[coords]int {
	const (
		asciiZeroOffset = 48
	)

	forest := make(map[coords]int)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for x, heightRune := range line {
			treeCoords := coords{
				x: x,
				y: y,
			}
			treeHeight := int(heightRune) - asciiZeroOffset
			forest[treeCoords] = treeHeight
		}

		y++
	}

	return forest
}

func isVisible(forest map[coords]int, treeCoords coords) bool {
	return false
}

func countVisible(forest map[coords]int) int {
	cnt := 0

	for treeCoords, _ := range forest {
		if isVisible(forest, treeCoords) {
			cnt++
		}
	}

	return cnt
}

func SolutionPt1(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)
	forest := parseForest(scanner)
	visibleCnt := countVisible(forest)

	return solution.IntResult(visibleCnt)
}

func SolutionPt2(input solution.Input) solution.Result {
	panic("not_implemented")
}
