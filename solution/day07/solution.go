package day07

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
)

type cliInteraction struct {
	in  string
	out []string
}

func parseCLIInteractions(scanner *bufio.Scanner) []*cliInteraction {
	var interactions []*cliInteraction

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line[0] == '$' {
			newInteraction := &cliInteraction{
				in: line[2:],
			}
			interactions = append(interactions, newInteraction)
		} else {
			lastInteraction := interactions[len(interactions)-1]
			lastInteraction.out = append(lastInteraction.out, line)
		}
	}

	return interactions
}

type node struct {
	name     string
	parent   *node
	size     int
	children []*node
}

func (n *node) IsDir() bool {
	return len(n.children) != 0
}

func (n *node) Size() int {
	if !n.IsDir() {
		return n.size
	}

	var size int
	for _, chN := range n.children {
		size += chN.Size()
	}

	return size
}

//nolint:cyclop
func buildFileTree(interactions []*cliInteraction) *node {
	root := &node{
		name: "/",
	}
	currentNode := root

	for _, interaction := range interactions {
		var cmdName, cmdArg string
		fmt.Sscanf(interaction.in, "%s %s", &cmdName, &cmdArg)
		switch cmdName {
		case "cd":
			switch cmdArg {
			case "/":
				currentNode = root
			case "..":
				currentNode = currentNode.parent
			default:
				for _, child := range currentNode.children {
					if child.name == cmdArg {
						currentNode = child
					}
				}
			}
		case "ls":
			var sizeOrDir, name string
			for _, outLine := range interaction.out {
				fmt.Sscanf(outLine, "%s %s", &sizeOrDir, &name)
				child := &node{
					name:   name,
					parent: currentNode,
				}
				if sizeOrDir != "dir" {
					size, _ := strconv.Atoi(sizeOrDir)
					child.size = size
				}
				currentNode.children = append(currentNode.children, child)
			}
		}
	}

	return root
}

func calculateWeirdTotalSize(root *node) int {
	var size int

	var walkFn func(curN *node)
	walkFn = func(curN *node) {
		if curN.IsDir() {
			dirSize := curN.Size()
			if dirSize <= 100000 {
				size += dirSize
			}
		}
		for _, chN := range curN.children {
			walkFn(chN)
		}
	}

	walkFn(root)

	return size
}

func SolutionPt1(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)
	interactions := parseCLIInteractions(scanner)
	fileTree := buildFileTree(interactions)
	totalSize := calculateWeirdTotalSize(fileTree)

	return solution.IntResult(totalSize)
}

func findBestCandidate(root *node) int {
	const (
		minRequired = 30000000
		total       = 70000000
	)

	rootSize := root.Size()
	free := total - rootSize
	needToFree := minRequired - free

	var candidate int

	var walkFn func(curN *node)
	walkFn = func(curN *node) {
		if curN.IsDir() {
			dirSize := curN.Size()
			if dirSize >= needToFree && (candidate == 0 || dirSize < candidate) {
				candidate = dirSize
			}
		}
		for _, chN := range curN.children {
			walkFn(chN)
		}
	}

	walkFn(root)

	return candidate
}

func SolutionPt2(input solution.Input) solution.Result {
	scanner := bufio.NewScanner(input)
	interactions := parseCLIInteractions(scanner)
	fileTree := buildFileTree(interactions)
	candidate := findBestCandidate(fileTree)

	return solution.IntResult(candidate)
}
