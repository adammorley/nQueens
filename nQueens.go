// solve the nQueens problem
package nQueens

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

/*
   create a new solution board by randomly choosing a
   column for the given row, and then randomly picking
   potential locations for the "pq" (proposed queen)
   until a spot is found, breaking out early if so
*/
func generateSolution(r, sz int) *board {
	var b *board = newBoard(sz)
	iq := b.randomColumn(r)
	b.placeQueen(iq)
	for i := 0; i < b.size; i++ {
		if i == r {
			continue
		}
		for {
			pq := b.randomColumn(i)
			if b.queenAllowed(pq) {
				b.placeQueen(pq)
				break
			}
		}
	}
	return b
}

/*
   generate all the solutions by generating a solution
   for each initial starting queen location (in a given
   row, across all the rows of the board)
*/
func generateSolutions(size int) []*board {
	var boards []*board = make([]*board, 0)
	for i := 0; i < size; i++ {
		b := generateSolution(i, size)
		boards = append(boards, b)
	}
	return boards
}

// functions related to squares (eq: the spots on the board)
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
	if abs(s0.r-s1.r) == abs(s1.c-s1.c) {
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
func (b *board) updateBlocked(sq square) {
	b.blockedR[sq.r] = true
	b.blockedC[sq.c] = true
}
func (b *board) createAllProposedQueens() (pq map[square]bool) {
	pq = make(map[square]bool)
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			pq[square{r, c}] = true
		}
	}
	return
}
func (b *board) randomColumn(r int) square {
	return square{r, rand.Intn(b.size)}
}
func (b *board) String() string {
	var sb strings.Builder
	sb.WriteString("board with size ")
	sb.WriteString(strconv.Itoa(b.size))
	sb.WriteString("\n")
	var queens byRow
	for k, _ := range b.queens {
		queens = append(queens, k)
	}
	sort.Sort(queens)
	for i := range queens {
		for j := 0; j < queens[i].c; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString("Q")
		sb.WriteString("\n")
	}
	return sb.String()
}
