package solution_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day01"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day02"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day03"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day04"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day05"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day06"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day07"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day08"
)

func input(t *testing.T, name string) io.ReadCloser {
	t.Helper()

	input, err := os.Open(name)
	require.NoError(t, err)

	return input
}

func golden(t *testing.T, name string) solution.Result {
	t.Helper()

	golden, err := os.Open(name)
	require.NoError(t, err)
	defer golden.Close()

	b, err := io.ReadAll(golden)
	require.NoError(t, err)

	return solution.Result(string(b))
}

func TestSolution(t *testing.T) {
	tcs := []struct {
		name       string
		solutionFn solution.Solution
		inputName  string
		goldenName string
	}{
		// day1
		{
			"day01 pt1 my input",
			day01.SolutionPt1,
			"day01/testdata/my.input.txt",
			"day01/testdata/my_pt1.golden.txt",
		},
		{
			"day01 pt2 my input",
			day01.SolutionPt2,
			"day01/testdata/my.input.txt",
			"day01/testdata/my_pt2.golden.txt",
		},

		// day2
		{
			"day02 pt1 my input",
			day02.SolutionPt1,
			"day02/testdata/my.input.txt",
			"day02/testdata/my_pt1.golden.txt",
		},
		{
			"day02 pt2 my input",
			day02.SolutionPt2,
			"day02/testdata/my.input.txt",
			"day02/testdata/my_pt2.golden.txt",
		},

		// day3
		{
			"day03 pt1 example input",
			day03.SolutionPt1,
			"day03/testdata/example.input.txt",
			"day03/testdata/example_pt1.golden.txt",
		},
		{
			"day03 pt1 my input",
			day03.SolutionPt1,
			"day03/testdata/my.input.txt",
			"day03/testdata/my_pt1.golden.txt",
		},
		{
			"day03 pt2 example input",
			day03.SolutionPt2,
			"day03/testdata/example.input.txt",
			"day03/testdata/example_pt2.golden.txt",
		},
		{
			"day03 pt2 my input",
			day03.SolutionPt2,
			"day03/testdata/my.input.txt",
			"day03/testdata/my_pt2.golden.txt",
		},

		// day4
		{
			"day04 pt1 example input",
			day04.SolutionPt1,
			"day04/testdata/example.input.txt",
			"day04/testdata/example_pt1.golden.txt",
		},
		{
			"day04 pt1 my input",
			day04.SolutionPt1,
			"day04/testdata/my.input.txt",
			"day04/testdata/my_pt1.golden.txt",
		},
		{
			"day04 pt2 example input",
			day04.SolutionPt2,
			"day04/testdata/example.input.txt",
			"day04/testdata/example_pt2.golden.txt",
		},
		{
			"day04 pt2 my input",
			day04.SolutionPt2,
			"day04/testdata/my.input.txt",
			"day04/testdata/my_pt2.golden.txt",
		},

		// day05
		{
			"day05 pt1 example input",
			day05.SolutionPt1,
			"day05/testdata/example.input.txt",
			"day05/testdata/example_pt1.golden.txt",
		},
		{
			"day05 pt1 my input",
			day05.SolutionPt1,
			"day05/testdata/my.input.txt",
			"day05/testdata/my_pt1.golden.txt",
		},
		{
			"day05 pt2 example input",
			day05.SolutionPt2,
			"day05/testdata/example.input.txt",
			"day05/testdata/example_pt2.golden.txt",
		},
		{
			"day05 pt2 my input",
			day05.SolutionPt2,
			"day05/testdata/my.input.txt",
			"day05/testdata/my_pt2.golden.txt",
		},

		// day06
		{
			"day06 pt1 example input",
			day06.SolutionPt1,
			"day06/testdata/example.input.txt",
			"day06/testdata/example_pt1.golden.txt",
		},
		{
			"day06 pt1 my input",
			day06.SolutionPt1,
			"day06/testdata/my.input.txt",
			"day06/testdata/my_pt1.golden.txt",
		},
		{
			"day06 pt2 example input",
			day06.SolutionPt2,
			"day06/testdata/example.input.txt",
			"day06/testdata/example_pt2.golden.txt",
		},
		{
			"day06 pt2 my input",
			day06.SolutionPt2,
			"day06/testdata/my.input.txt",
			"day06/testdata/my_pt2.golden.txt",
		},

		// day07
		{
			"day07 pt1 example input",
			day07.SolutionPt1,
			"day07/testdata/example.input.txt",
			"day07/testdata/example_pt1.golden.txt",
		},
		{
			"day07 pt1 my input",
			day07.SolutionPt1,
			"day07/testdata/my.input.txt",
			"day07/testdata/my_pt1.golden.txt",
		},
		{
			"day07 pt2 example input",
			day07.SolutionPt2,
			"day07/testdata/example.input.txt",
			"day07/testdata/example_pt2.golden.txt",
		},
		{
			"day07 pt2 my input",
			day07.SolutionPt2,
			"day07/testdata/my.input.txt",
			"day07/testdata/my_pt2.golden.txt",
		},

		// day08
		{
			"day08 pt1 example input",
			day08.SolutionPt1,
			"day08/testdata/example.input.txt",
			"day08/testdata/example_pt1.golden.txt",
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			in := input(t, tc.inputName)
			defer in.Close()
			r := tc.solutionFn(in)
			er := golden(t, tc.goldenName)
			assert.Equal(t, er, r)
		})
	}
}
