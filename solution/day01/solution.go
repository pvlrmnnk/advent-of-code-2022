package day01

import (
	"bufio"
	"container/heap"
	"strconv"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
)

// get max amount of calories per elf
func SolutionPt1(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)

	var max, totalPerElf int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// empty line is elf separator
			// that means we successfully calculate total calories per elf
			if totalPerElf > max {
				max = totalPerElf
			}
			totalPerElf = 0

			continue
		}
		calories, _ := strconv.Atoi(line)
		totalPerElf += calories
	}

	return solution.IntResult(max)
}

// heap implementation from https://pkg.go.dev/container/heap#example-package-IntHeap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // for max heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// get max amount of calories by top3
// this solution is based on max heap
func SolutionPt2(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)

	caloriesMaxHeap := &IntHeap{}
	heap.Init(caloriesMaxHeap)
	var totalPerElf int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// empty line is elf separator
			// that means we successfully calculate total calories per elf
			heap.Push(caloriesMaxHeap, totalPerElf)
			totalPerElf = 0

			continue
		}
		calories, _ := strconv.Atoi(line)
		totalPerElf += calories
	}

	top3sum := heap.Pop(caloriesMaxHeap).(int) + heap.Pop(caloriesMaxHeap).(int) + heap.Pop(caloriesMaxHeap).(int)

	return solution.IntResult(top3sum)
}
