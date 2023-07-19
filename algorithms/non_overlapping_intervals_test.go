package non_overlapping_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_eraseOverlapIntervals_1(t *testing.T) {
	assert.Equal(
		t,
		1,
		eraseOverlapIntervals([][]int{
			{1, 2}, {2, 3}, {3, 4}, {1, 3},
		}),
	)
}

func Test_eraseOverlapIntervals_2(t *testing.T) {
	assert.Equal(
		t,
		2,
		eraseOverlapIntervals([][]int{
			{1, 2}, {1, 2}, {1, 2},
		}),
	)
}

func Test_eraseOverlapIntervals_3(t *testing.T) {
	assert.Equal(
		t,
		0,
		eraseOverlapIntervals([][]int{
			{1, 2}, {2, 3},
		}),
	)
}
