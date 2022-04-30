package main

func combine(cand int, l [][]int, r [][]int) [][]int {
	s := make([][]int, len(l)+len(r))
	for i := 0; i < len(l); i++ {
		s[i] = append([]int{cand}, l[i]...)
	}

	for i := 0; i < len(r); i++ {
		s[len(l)+i] = r[i]
	}
	return s
}

func combinationSum(candidates []int, target int) [][]int {
	sum := make([][][]int, target+1)
	newsum := make([][][]int, target+1)
	sum[0] = [][]int{{}}

	for _, cand := range candidates {
		for j := 0; j <= target; j++ {
			if cand > j {
				newsum[j] = sum[j]
			} else {
				newsum[j] = combine(cand, newsum[j-cand], sum[j])
			}
		}
		sum, newsum = newsum, sum
	}
	return sum[target]
}
