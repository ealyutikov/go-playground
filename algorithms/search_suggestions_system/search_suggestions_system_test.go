package search_suggestions_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_suggestedProducts_1(t *testing.T) {
	assert.Equal(
		t,
		[][]string{
			{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"},
		},
		suggestedProducts(
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
		),
	)
}

func Test_suggestedProducts_2(t *testing.T) {
	assert.Equal(
		t,
		[][]string{
			{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"},
		},
		suggestedProducts(
			[]string{"havana"},
			"havana",
		),
	)
}
