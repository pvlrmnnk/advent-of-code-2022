package day01

import (
	"bufio"
	"io"
	"strconv"
)

func Solution(caloriesInpit io.Reader) int {
	scanner := bufio.NewScanner(caloriesInpit)

	var max, localMax int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if localMax > max {
				max = localMax
			}
			localMax = 0
			continue
		}
		calories, _ := strconv.Atoi(line)
		localMax += calories
	}

	return max
}
