package letter_combinations_of_a_phone_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations_1(t *testing.T) {
	assert.Equal(
		t,
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		letterCombinations("23"),
	)
}

func Test_letterCombinations_2(t *testing.T) {
	assert.Equal(
		t,
		[]string{"a", "b", "c"},
		letterCombinations("2"),
	)
}
