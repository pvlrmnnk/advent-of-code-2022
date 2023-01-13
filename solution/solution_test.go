package solution_test

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
	"github.com/pvlrmnnk/advent-of-code-2022/solution/day04"
)

func input(t *testing.T, name string) io.ReadCloser {
	t.Helper()

	input, err := os.Open(name)
	require.NoError(t, err)

	return input
}

func goldenInt(t *testing.T, name string) int {
	t.Helper()

	golden, err := os.Open(name)
	require.NoError(t, err)
	defer golden.Close()

	b, err := io.ReadAll(golden)
	require.NoError(t, err)

	i, err := strconv.Atoi(string(b))
	require.NoError(t, err)

	return i
}

func TestSolution(t *testing.T) {
	ss := []struct {
		name       string
		solutionFn solution.Solution
		inputName  string
		goldenName string
	}{
		{
			"day4 pt1 example input",
			day04.SolutionPt1,
			"day04/testdata/example.input.txt",
			"day04/testdata/example_pt1.golden.txt",
		},
		{
			"day4 pt1 my input",
			day04.SolutionPt1,
			"day04/testdata/my.input.txt",
			"day04/testdata/my_pt1.golden.txt",
		},
		{
			"day4 pt2 example input",
			day04.SolutionPt2,
			"day04/testdata/example.input.txt",
			"day04/testdata/example_pt2.golden.txt",
		},
		{
			"day4 pt2 my input",
			day04.SolutionPt2,
			"day04/testdata/my.input.txt",
			"day04/testdata/my_pt2.golden.txt",
		},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			in := input(t, s.inputName)
			defer in.Close()
			r := s.solutionFn(in)
			switch r := r.(type) { //nolint:gocritic
			case solution.IntResult:
				er := goldenInt(t, s.goldenName)
				assert.Equal(t, er, r.Int())
			}
		})
	}
}