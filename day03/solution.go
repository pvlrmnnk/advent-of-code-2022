package day03

import (
	"bufio"
	"io"
)

// ascii codes from https://www.ascii-code.com/
func priority(r rune) int {
	switch {
	case r >= 97 && r <= 122:
		// 1-26 for a-z
		return int(r) - 96
	case r >= 65 && r <= 90:
		// 27-52 for A-Z
		return int(r) - 64 + 26
	default:
		return 0
	}
}

func SolutionPt1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		runes := []rune(line)
		seen := make(map[rune]struct{}, len(runes))
		for i := 0; i < len(runes)/2; i++ {
			seen[runes[i]] = struct{}{}
		}
		for i := len(runes) / 2; i < len(runes); i++ {
			if _, ok := seen[runes[i]]; ok {
				sum += priority(runes[i])

				break
			}
		}
	}

	return sum
}

func offset(r rune) int {
	return priority(r) - 1
}

func offset2(i uint64) int {
	cnt := 1
	for i != 1 {
		cnt++
		i >>= 1
	}

	return cnt
}

func priority2(mask uint64) int {
	return offset2(mask)
}

func SolutionPt2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	var sum int
	var groupMember int
	var mask1, mask2, mask3 uint64
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for _, r := range line {
			switch groupMember {
			case 0:
				mask1 |= 1 << offset(r)
			case 1:
				mask2 |= 1 << offset(r)
			case 2:
				mask3 |= 1 << offset(r)
			}
		}

		if groupMember == 2 {
			commonMask := mask1 & mask2 & mask3
			priority := priority2(commonMask)
			sum += priority
			// reset
			groupMember = 0
			mask1, mask2, mask3 = 0, 0, 0
		} else {
			groupMember++
		}
	}

	return sum
}
