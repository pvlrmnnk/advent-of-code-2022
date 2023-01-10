package day01

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolutionPt1(t *testing.T) {
	const CaloriesTOP1 = 75501
	input, err := os.Open("testdata/input.txt")
	require.NoError(t, err)
	defer input.Close()

	require.Equal(t, CaloriesTOP1, SolutionPt1(input))
}

func TestSolutionPt2(t *testing.T) {
	const CaloriesTOP3 = 215594
	input, err := os.Open("testdata/input.txt")
	require.NoError(t, err)
	defer input.Close()

	require.Equal(t, CaloriesTOP3, SolutionPt2(input))
}
