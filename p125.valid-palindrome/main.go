package main

func isPalindrome(s string) bool {
	bs := []byte(s)
	nbs := []byte{}
	diff := byte('a' - 'A')
	for i := range bs {
		bc := bs[i]
		if bc >= 'A' && bc <= 'Z' {
			nbs = append(nbs, bc+diff)
		} else if (bc >= 'a' && bc <= 'z') || (bc >= '0' && bc <= '9') {
			nbs = append(nbs, bc)
		}
	}

	rnbs := []byte{}
	for i := len(nbs) - 1; i >= 0; i-- {
		rnbs = append(rnbs, nbs[i])
	}
	return string(nbs) == string(rnbs)
}
