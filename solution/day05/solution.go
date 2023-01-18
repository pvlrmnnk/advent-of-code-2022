package day05

import (
	"bufio"
	"container/list"
	"fmt"
	"strings"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
)

type stack struct {
	l *list.List
}

func (s *stack) Push(crate string) {
	s.l.PushBack(crate)
}

func (s *stack) Pop() string {
	back := s.l.Back()
	if back == nil {
		return ""
	}

	return s.l.Remove(back).(string)
}

func (s *stack) Peek() string {
	back := s.l.Back()
	if back == nil {
		return ""
	}

	return back.Value.(string)
}

func newStack() *stack {
	return &stack{
		list.New(),
	}
}

type procedure struct {
	cnt       int
	fromIndex int
	toIndex   int
}

func parseStacks(scanner *bufio.Scanner) []*stack {
	var stacks []*stack

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lineParts := strings.Split(line, " ")
		rawCrates := lineParts[1:]
		crateStack := newStack()
		for _, rawCrate := range rawCrates {
			crate := strings.TrimFunc(rawCrate, func(r rune) bool {
				return r == '[' || r == ']'
			})
			crateStack.Push(crate)
		}

		stacks = append(stacks, crateStack)
	}

	return stacks
}

func parseProcedures(scanner *bufio.Scanner) []*procedure {
	var procedures []*procedure

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var proc procedure
		fmt.Sscanf(line, "move %d from %d to %d", &proc.cnt, &proc.fromIndex, &proc.toIndex)
		// fix index offset
		proc.fromIndex--
		proc.toIndex--

		procedures = append(procedures, &proc)
	}

	return procedures
}

func applyProcedure(proc *procedure, stacks []*stack) {
	for i := 0; i < proc.cnt; i++ {
		crate := stacks[proc.fromIndex].Pop()
		stacks[proc.toIndex].Push(crate)
	}
}

func onTopOfStacks(stacks []*stack) string {
	var b strings.Builder
	for _, s := range stacks {
		b.WriteString(s.Peek())
	}

	return b.String()
}

func SolutionPt1(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)
	stacks := parseStacks(scanner)
	procedures := parseProcedures(scanner)

	for _, proc := range procedures {
		applyProcedure(proc, stacks)
	}

	return solution.Result(onTopOfStacks(stacks))
}

func applyProcedure2(proc *procedure, stacks []*stack) {
	crates := make([]string, 0, proc.cnt)
	for i := 0; i < proc.cnt; i++ {
		crate := stacks[proc.fromIndex].Pop()
		crates = append(crates, crate)
	}
	for i := len(crates) - 1; i >= 0; i-- {
		stacks[proc.toIndex].Push(crates[i])
	}
}

func SolutionPt2(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)
	stacks := parseStacks(scanner)
	procedures := parseProcedures(scanner)

	for _, proc := range procedures {
		applyProcedure2(proc, stacks)
	}

	return solution.Result(onTopOfStacks(stacks))
}
