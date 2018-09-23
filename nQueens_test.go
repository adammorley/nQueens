package nQueens

import "testing"

var knownSolutions = map[int]int{4: 2, 5: 10, 6: 4, 7: 40, 8: 92}

func TestSolve(t *testing.T) {
	l := solve()
	if l.Len() != knownSolutions[boardSize] {
		t.Error("wrong number of solutions")
	}
}
