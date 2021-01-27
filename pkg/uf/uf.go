// Package uf はUnion-Findのデータ構造を提供します。
// 各操作のならし計算量はO(α(n))です（α(n)はアッカーマン関数の逆関数）。
//
package uf

// UnionFind はUnion-Find用のint型配列です。
type UnionFind []int

// newUnionFind はn要素のグループで初期化したUnion-Findデータ構造を作成し、そのポインタを返します。
func newUnionFind(n int) *UnionFind {
	uf := make(UnionFind, n)
	for i := 0; i < n; i++ {
		uf[i] = -1
	}
	return &uf
}

// Find は要素xの所属するグループを返します。
func (uf UnionFind) Find(x int) int {
	if uf[x] < 0 {
		return x
	}
	uf[x] = uf.Find(uf[x])
	return uf[x]
}

// Unite は要素xと要素yのグループを併合します。
// 併合された場合はtrue、同じグループであったときはfalseを返します。
func (uf UnionFind) Unite(x, y int) bool {
	x, y = uf.Find(x), uf.Find(y)
	if x != y {
		if uf[x] > uf[y] {
			x, y = y, x
		}
		uf[x] += uf[y]
		uf[y] = x
		return true
	}
	return false
}

// Same は要素xと要素yが同じグループであればtrue、そうでなければfalseを返します。
func (uf UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Size は要素xの所属するグループの要素数を返します。
func (uf UnionFind) Size(x int) int {
	return -uf[uf.Find(x)]
}
