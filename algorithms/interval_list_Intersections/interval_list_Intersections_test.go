package interval_list_Intersections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalIntersection(t *testing.T) {
	assert.Equal(
		t,
		[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		intervalIntersection(
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}),
	)
}
