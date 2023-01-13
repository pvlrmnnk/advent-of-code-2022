package solution_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day02"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day03"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day04"
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
