package day01

import (
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	input, err := os.Open("testdata/calories.input")
	require.NoError(t, err)
	defer input.Close()
	golden, err := os.Open("testdata/calories.golden")
	require.NoError(t, err)
	defer golden.Close()

	require.Equal(t, expextedResult(t, golden), Solution(input))
}

func expextedResult(t *testing.T, golden io.Reader) int {
	t.Helper()

	b, _ := io.ReadAll(golden)
	i, _ := strconv.Atoi(string(b))

	return i
}
