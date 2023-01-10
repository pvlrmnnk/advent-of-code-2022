package day02

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolutionPt1(t *testing.T) {
	const Score = 13809
	input, err := os.Open("testdata/input.txt")
	require.NoError(t, err)
	defer input.Close()

	require.Equal(t, Score, SolutionPt1(input))
}

func TestSolutionPt2(t *testing.T) {
	const Score = 12316
	input, err := os.Open("testdata/input.txt")
	require.NoError(t, err)
	defer input.Close()

	require.Equal(t, Score, SolutionPt2(input))
}
