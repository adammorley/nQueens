package nQueens

//import "sort"
import "testing"

func TestSolutions(t *testing.T) {
	for i := 4; i <= 8; i++ {
		generateTest(t, i)
	}
}
func generateTest(t *testing.T, i int) {
	boards := generateSolutions(i)
	if len(boards) < 1 {
		t.Error("no solutions found for", i)
	}
}

// test that the solutions are different
/*func TestDifferent(t *testing.T) {
	boards := generateSolutions(2)
    t.Error("foo")
	var queens map[int]byRow = make(map[int]byRow)
	for i := range boards {
		for k, _ := range boards[i].queens {
			queens[i] = append(queens[i], k)
		}
		sort.Sort(queens[i])
	}
	for i := range queens {
		for j := range queens {
			if i == j {
				continue
			}
            var same = true
			for k := 0; k < len(queens[i]); k++ {
				if queens[i][k].r != queens[j][k].r || queens[i][k].c != queens[j][k].c {
					same = false
				}
			}
            if same {
                t.Error("two are same:", i, j, boards[i],boards[j])
            }
		}
	}
}*/
