// solve the nQueens problem
package nQueens

import (
)

/*
   create a new solution board by randomly choosing a
   column for the given row, and then picking
   potential locations for the "pq" (proposed queen)
   until a spot is found, breaking out early if so
*/
func generateSolution(r, sz int) *board {
	if sz < 4 {
		panic("small")
	}
	var b *board = newBoard(sz)
	iq := b.randomColumn(r)
	b.placeQueen(iq)
	for i := 0; i < b.size; i++ {
		if i == r {
			continue
		}
		for j := 0; j < b.size; j++ {
			if b.blockedC[j] {
				continue
			}
			pq := square{i, j}
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
	m, n := 0, 100
	for m < n {
		for i := 0; i < size; i++ {
			b := generateSolution(i, size)
			if len(b.queens) == size {
				boards = append(boards, b)
			}
		}
        if len(boards) > 0 {
            break
        }
		m++
	}
	return boards
}
