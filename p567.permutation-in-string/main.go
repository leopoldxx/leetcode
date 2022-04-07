package main

type container struct {
	hash    int
	letters [26]int
}

func (c *container) push(b byte) {
	c.hash += int(b)
	c.letters[int(b-'a')]++
}
func (c *container) pop(b byte) {
	c.hash -= int(b)
	c.letters[int(b-'a')]--
}
func (c *container) equal(cc *container) bool {
	if c.hash != cc.hash {
		return false
	}
	for i := 0; i < 26; i++ {
		if c.letters[i] != cc.letters[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	bs1 := []byte(s1)
	bs2 := []byte(s2)

	c1 := &container{}
	c2 := &container{}

	for i := 0; i < len(bs1); i++ {
		c1.push(bs1[i])
		c2.push(bs2[i])
	}
	if c1.equal(c2) {
		return true
	}

	j := 0
	jj := len(bs1)

	for jj < len(bs2) {
		c2.pop(bs2[j])
		c2.push(bs2[jj])
		if c1.equal(c2) {
			return true
		}
		j++
		jj++
	}
	return false
}
