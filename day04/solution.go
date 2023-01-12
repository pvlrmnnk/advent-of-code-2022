package day04

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type assignmentRange struct {
	left, right int
}

func (ar *assignmentRange) From(s string) {
	parts := strings.Split(s, "-")
	ar.left, _ = strconv.Atoi(parts[0])
	ar.right, _ = strconv.Atoi(parts[1])
}

func (ar assignmentRange) OverlapFull(ar2 assignmentRange) bool {
	return ar.left <= ar2.left && ar.right >= ar2.right
}

func ranges(s string) (assignmentRange, assignmentRange) {
	var ar1, ar2 assignmentRange
	parts := strings.Split(s, ",")
	ar1.From(parts[0])
	ar2.From(parts[1])

	return ar1, ar2
}

func SolutionPt1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	var overlapCnt int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ar1, ar2 := ranges(line)
		if ar1.OverlapFull(ar2) || ar2.OverlapFull(ar1) {
			overlapCnt++
		}
	}

	return overlapCnt
}

func (ar assignmentRange) OverlapPartial(ar2 assignmentRange) bool {
	leftBorderOverlap := ar.left <= ar2.left && ar.right >= ar2.left
	rightBorderOverlap := ar.right >= ar2.right && ar.left <= ar2.right

	return leftBorderOverlap || rightBorderOverlap
}

func SolutionPt2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	var overlapCnt int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ar1, ar2 := ranges(line)
		if ar1.OverlapPartial(ar2) || ar2.OverlapPartial(ar1) {
			overlapCnt++
		}
	}

	return overlapCnt
}
