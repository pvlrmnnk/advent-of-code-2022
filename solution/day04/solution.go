package day04

import (
	"bufio"
	"fmt"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
)

type assignmentRange struct {
	leftBorder, rightBorder int
}

func (ar assignmentRange) OverlapFull(ar2 assignmentRange) bool {
	return ar.leftBorder <= ar2.leftBorder && ar.rightBorder >= ar2.rightBorder
}

func parseAssignmentRanges(line string) (assignmentRange, assignmentRange) {
	var ar1, ar2 assignmentRange
	fmt.Sscanf(
		line,
		"%d-%d,%d-%d",
		&ar1.leftBorder,
		&ar1.rightBorder,
		&ar2.leftBorder,
		&ar2.rightBorder,
	)

	return ar1, ar2
}

func SolutionPt1(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)

	var overlapCnt int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ar1, ar2 := parseAssignmentRanges(line)
		if ar1.OverlapFull(ar2) || ar2.OverlapFull(ar1) {
			overlapCnt++
		}
	}

	return solution.IntResult(overlapCnt)
}

func (ar assignmentRange) OverlapPartial(ar2 assignmentRange) bool {
	leftBorderOverlap := ar.leftBorder <= ar2.leftBorder && ar.rightBorder >= ar2.leftBorder
	rightBorderOverlap := ar.rightBorder >= ar2.rightBorder && ar.leftBorder <= ar2.rightBorder

	return leftBorderOverlap || rightBorderOverlap
}

func SolutionPt2(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)

	var overlapCnt int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ar1, ar2 := parseAssignmentRanges(line)
		if ar1.OverlapPartial(ar2) || ar2.OverlapPartial(ar1) {
			overlapCnt++
		}
	}

	return solution.IntResult(overlapCnt)
}
