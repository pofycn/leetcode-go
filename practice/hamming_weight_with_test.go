package practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func HammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			count++
		}
	}
	return count
}

func TestHammingWeight(t *testing.T) {
	assert := assert.New(t)

	var x uint32
	x = 0b00000000000000000100000000001011

	assert.Equal(HammingWeight(x), 4, "Should be equal to 4")
}
