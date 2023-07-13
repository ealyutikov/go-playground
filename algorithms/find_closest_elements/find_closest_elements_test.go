package find_closest_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findClosestElements_1(t *testing.T) {
	assert.Equal(
		t,
		[]int{1, 2, 3, 4},
		findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3),
	)
}

func Test_findClosestElements_2(t *testing.T) {
	assert.Equal(
		t,
		[]int{1, 2, 3, 4},
		findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1),
	)
}
