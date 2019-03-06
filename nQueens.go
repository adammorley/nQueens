// solve the nQueens problem
package nQueens

import (
	"container/list"
	"strconv"
	"strings"
)

const boardSize = 5

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func queenAllowed(r, c int, s [boardSize]int) bool {
	for i := 0; i < r; i++ {
		if s[i] == c || abs(i-r) == abs(c-s[i]) {
			return false
		}
	}
	return true
}
func placeQueen(r int, s [boardSize]int, l *list.List) {
	if r == boardSize {
		l.PushBack(s)
	} else {
		for c := 0; c < boardSize; c++ {
			if queenAllowed(r, c, s) {
				s[r] = c
				placeQueen(r+1, s, l)
			}
		}
	}
}
func solve() *list.List {
	l := list.New()
	var s [boardSize]int
	placeQueen(0, s, l)
	return l
}
func printSolutions(l *list.List) string {
	var sb strings.Builder
	sb.WriteString("board with size ")
	sb.WriteString(strconv.Itoa(boardSize))
	sb.WriteString("\n")
	e := l.Front()
	for e != nil {
		s := e.Value.([boardSize]int)
		for i := range s {
			for j := 0; j < s[i]; j++ {
				sb.WriteString(" ")
			}
			sb.WriteString("Q\n")
		}
		sb.WriteString("\n")
		e = e.Next()
	}
	return sb.String()
}
func printValues(l *list.List) string {
	var sb strings.Builder
	sb.WriteString("board with size ")
	sb.WriteString(strconv.Itoa(boardSize))
	sb.WriteString("\n")
	e := l.Front()
	for e != nil {
		s := e.Value.([boardSize]int)
		for i := range s {
            sb.WriteString("board row ")
            sb.WriteString(strconv.Itoa(i))
            sb.WriteString(" is column ")
            sb.WriteString(strconv.Itoa(s[i]))
            sb.WriteString("\n")
		}
		sb.WriteString("\n")
		e = e.Next()
	}
	return sb.String()
}
