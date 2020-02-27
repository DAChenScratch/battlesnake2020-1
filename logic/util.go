package logic

// Stolen from StackOverflow:
// https://stackoverflow.com/questions/29002724/implement-ruby-style-cartesian-product-in-go
func cartesianProduct(moves [][]Move) [][]Move {
	n := 1
	for _, m := range moves {
		if len(m) != 0 {
			n = n * len(m)
		}
	}
	ret := make([][]Move, n)
	k := 0
	lengths := func(i int) int { return len(moves[i]) }
	for tmp := make([]int, len(moves)); tmp[0] < lengths(0); nextIndex(tmp, lengths) {
		var m []Move
		for i, j := range tmp {
			if len(moves[i]) != 0 {
				m = append(m, moves[i][j])
			}
		}
		ret[k] = m
		k++
	}
	return ret
}

func nextIndex(tmp []int, lengths func(i int) int) {
	for i := len(tmp) - 1; i >= 0; i-- {
		tmp[i]++
		if i == 0 || tmp[i] < lengths(i) {
			return
		}
		tmp[i] = 0
	}
}
