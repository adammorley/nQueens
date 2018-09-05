package nQueens

// avoid having to cast to float and back with math library
func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// row, column
type square struct {
	r, c int
}

// calculate whether two squares are on the same diagonal
func (s0 square) sameDiagonal(s1 square) bool {
	if abs(s0.r-s1.r) == abs(s0.c-s1.c) {
		return true
	}
	return false
}

// type to allow sorting (for printing)
type byRow []square

func (r byRow) Len() int {
	return len(r)
}
func (r byRow) Less(i, j int) bool {
	return r[i].r < r[j].r
}
func (r byRow) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
