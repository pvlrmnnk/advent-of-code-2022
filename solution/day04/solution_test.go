package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAssignmentRanges(t *testing.T) {
	t.Parallel()

	ar1, ar2 := parseAssignmentRanges("1-2,3-4")
	assert.Equal(t, 1, ar1.leftBorder)
	assert.Equal(t, 2, ar1.rightBorder)
	assert.Equal(t, 3, ar2.leftBorder)
	assert.Equal(t, 4, ar2.rightBorder)
}

func TestAssignmentRangeOverlapFull(t *testing.T) {
	t.Parallel()

	t.Skip()
}

func TestAssignmentRangeOverlapPartial(t *testing.T) {
	t.Parallel()

	t.Skip()
}
