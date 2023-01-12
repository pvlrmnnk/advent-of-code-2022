package day04

import (
	"testing"

	"github.com/pvlrmnnk/advent-of-code-2022/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolutionPt1(t *testing.T) {
	tcs := []struct {
		input, golden string
	}{
		{
			"testdata/example.input.txt",
			"testdata/example.golden.txt",
		},
		{
			"testdata/my.input.txt",
			"testdata/my.golden.txt",
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			in, err := common.Input(tc.input)
			require.NoError(t, err)
			defer in.Close()
			g, err := common.GoldenInt(tc.golden)
			require.NoError(t, err)
			r := SolutionPt1(in)
			assert.Equal(t, g, r)
		})
	}
}

func TestSolutionPt2(t *testing.T) {
	tcs := []struct {
		input, golden string
	}{
		{
			"testdata/example.input.txt",
			"testdata/example2.golden.txt",
		},
		{
			"testdata/my.input.txt",
			"testdata/my2.golden.txt",
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			in, err := common.Input(tc.input)
			require.NoError(t, err)
			defer in.Close()
			g, err := common.GoldenInt(tc.golden)
			require.NoError(t, err)
			r := SolutionPt2(in)
			assert.Equal(t, g, r)
		})
	}
}
