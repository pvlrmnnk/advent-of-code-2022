package day01

import (
	"bufio"
	"container/heap"
	"io"
	"strconv"
)

// get max amount of calories per elf
func SolutionPt1(input io.Reader) int {
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

	return max
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
func SolutionPt2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	h := &IntHeap{}
	heap.Init(h)
	var totalPerElf int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// empty line is elf separator
			// that means we successfully calculate total calories per elf
			heap.Push(h, totalPerElf)
			totalPerElf = 0
			continue
		}
		calories, _ := strconv.Atoi(line)
		totalPerElf += calories
	}

	return heap.Pop(h).(int) + heap.Pop(h).(int) + heap.Pop(h).(int)
}
