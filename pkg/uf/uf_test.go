package uf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://atcoder.jp/contests/practice2/tasks/practice2_a
func TestAtcoderPractice2Sample1(t *testing.T) {
	n, q := 4, 7
	query := [][]int{
		{1, 0, 1},
		{0, 0, 1},
		{0, 2, 3},
		{1, 0, 1},
		{1, 1, 2},
		{0, 0, 2},
		{1, 1, 3},
	}
	expected := []int{0, 1, 0, 1}

	actual := []int{}
	uf := NewUnionFind(n)
	for i := 0; i < q; i++ {
		t, u, v := query[i][0], query[i][1], query[i][2]
		if t == 0 {
			uf.Unite(u, v)
		} else {
			if uf.Same(u, v) {
				actual = append(actual, 1)
			} else {
				actual = append(actual, 0)
			}
		}
	}

	assert.Equal(t, expected, actual)
}
