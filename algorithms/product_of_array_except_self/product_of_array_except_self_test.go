package product_of_array_except_self

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf_1(t *testing.T) {
	assert.Equal(
		t,
		[]int{24, 12, 8, 6},
		productExceptSelf([]int{1, 2, 3, 4}),
	)
}

func Test_productExceptSelf_2(t *testing.T) {
	assert.Equal(
		t,
		[]int{0, 0, 9, 0, 0},
		productExceptSelf([]int{-1, 1, 0, -3, 3}),
	)
}
