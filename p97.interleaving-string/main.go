package main

func isInterleave(s1 string, s2 string, s3 string) bool {

	bs1, n1 := []byte(s1), len(s1)
	bs2, n2 := []byte(s2), len(s2)
	bs3, n3 := []byte(s3), len(s3)

	if n1+n2 != n3 {
		return false
	}

	interleaved := make([][]bool, n1+1)
	for i := 0; i < n1; i++ {
		interleaved[i] = make([]bool, n2+1)
	}
	interleaved[0][0] = true

	for i := 1; i <= n1; i++ {
		if bs1[i-1] == bs3[i-1] {
			interleaved[i][0] = true
		} else {
			break
		}
	}
	for i := 1; i <= n2; i++ {
		if bs2[i-1] == bs3[i-1] {
			interleaved[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if (interleaved[i-1][j] && bs1[i-1] == bs3[i+j-1]) || (interleaved[i][j-1] && bs2[j-1] == bs3[i+j-1]) {
				interleaved[i][j] = true
			}
		}
	}
	return interleaved[n1][n2]
}
