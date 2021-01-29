// Package bit はBinary Indexed Tree(Fenwick Tree)のデータ構造を提供します。
//
// 各操作の計算量はO(log n)です。
package bit

// BIT はBinary Indexed Treeの値を保持するint型スライスです。
type BIT []int

// NewBinaryIndexedTree は長さnのBinary Indexed Treeを作成し、ポインタを返します。
func NewBinaryIndexedTree(n int) *BIT {
	bit := make(BIT, n+1)
	return &bit
}

// Add はi番目の要素に値xを加えます。
func (bit BIT) Add(i, x int) {
	if !(0 <= i && i < len(bit)) {
		panic("")
	}

	for i++; i < len(bit); i += i & -i {
		bit[i] += x
	}
}

// Sum はi番目の要素までの和を返します。
func (bit BIT) Sum(i int) int {
	if !(i < len(bit)) {
		panic("")
	}

	s := 0
	for i++; i > 0; i -= i & -i {
		s += bit[i]
	}
	return s
}

// RangeSum は区間[l,r)の要素の和を返します。
func (bit BIT) RangeSum(l, r int) int {
	if !(0 <= l && l <= r && r <= len(bit)) {
		panic("")
	}

	return bit.Sum(r-1) - bit.Sum(l-1)
}

// LowerBound は Sum(i)>=xを満たす最小のインデックスiを返します。存在しない場合配列サイズnを返します。
// 各要素は非負である必要があります。
func (bit BIT) LowerBound(x int) int {
	idx, k := 0, 1
	for k < len(bit) {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k < len(bit) && bit[idx+k] < x {
			x -= bit[idx+k]
			idx += k
		}
	}
	return idx
}
