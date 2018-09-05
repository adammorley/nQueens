package nQueens

import "math/rand"

/*
   a representation of the board.
   a board has a given row or column blocked if a queen is in that row or column
   the size is arbitrary (normally boards are 8x8)
   the queens present on the board
*/
type board struct {
	blockedR map[int]bool
	blockedC map[int]bool
	size     int
	queens   map[square]bool
}
func newBoard(size int) (b *board) {
	b = new(board)
	b.size = size
	b.queens = make(map[square]bool)
	b.blockedR = make(map[int]bool)
	b.blockedC = make(map[int]bool)
	return
}
func (b *board) placeQueen(q square) {
	b.updateBlocked(q)
	b.queens[q] = true
}
func (b *board) queenAllowed(sq square) bool {
	if b.blockedR[sq.r] || b.blockedC[sq.c] {
		return false
	}
	for q, _ := range b.queens {
		if sq.sameDiagonal(q) {
			return false
		}
	}
	return true
}
func (b *board) randomColumn(r int) square {
	return square{r, rand.Intn(b.size)}
}
func (b *board) updateBlocked(sq square) {
	b.blockedR[sq.r] = true
	b.blockedC[sq.c] = true
}
