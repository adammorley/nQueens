package nQueens

import (
	"sort"
	"strconv"
	"strings"
)

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
