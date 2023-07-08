package find_the_duplicate_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate_1(t *testing.T) {
	assert.Equal(
		t,
		2,
		findDuplicate([]int{1, 3, 4, 2, 2}),
	)
}

func TestFindDuplicate_2(t *testing.T) {
	assert.Equal(
		t,
		3,
		findDuplicate([]int{3, 1, 3, 4, 2}),
	)
}
