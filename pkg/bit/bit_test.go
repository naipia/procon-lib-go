package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryIndexedTree(t *testing.T) {
	n := 5
	bit := newBinaryIndexedTree(n)

	for i := 0; i < n; i++ {
		bit.Add(i, i+1)
	}

	// Test Sum
	for i := 0; i < n; i++ {
		expected := (i + 2) * (i + 1) / 2
		actual := bit.Sum(i)
		assert.Equal(t, expected, actual)
	}

	// Test RangeSum
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			expected := (j+1)*j/2 - (i+1)*i/2
			actual := bit.RangeSum(i, j)
			assert.Equal(t, expected, actual)
		}
	}
}

func TestLowerBound(t *testing.T) {
	n := 10
	bit := newBinaryIndexedTree(n)
	for i := 0; i < n; i++ {
		bit.Add(i, 1)
	}

	for x := 1; x <= n; x++ {
		assert.Equal(t, x-1, bit.LowerBound(x))
	}
	assert.Equal(t, 0, bit.LowerBound(0))
	assert.Equal(t, n, bit.LowerBound(n+10))
}
