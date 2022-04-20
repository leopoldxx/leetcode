package main

// https://en.wikipedia.org/wiki/Longest_palindromic_substring

func expand(s string) []byte {
	ns := make([]byte, 2*len(s)+1)
	for i, b := range []byte(s) {
		ii := i << 1
		ns[ii] = '|'
		ns[ii+1] = b
	}
	ns[len(s)*2] = '|'
	return ns
}
func shrink(bs []byte) string {
	if len(bs) == 0 {
		return ""
	}
	res := []byte{}
	i := 1
	for i < len(bs)-1 {
		res = append(res, bs[i])
		i = i + 2
	}
	return string(res)
}
func findMax(maxPalindrome []int) (index int, radius int) {
	radius = -1
	for i, r := range maxPalindrome {
		if radius < r {
			index = i
			radius = r
		}
	}
	return index, radius
}

func longestPalindrome(s string) string {
	ns := expand(s)
	maxPalindrome := make([]int, len(ns))

	center := 0
	radius := 0

	for center < len(ns) {
		expandRadius := radius + 1
		left := center - expandRadius
		right := center + expandRadius
		for left >= 0 && right < len(ns) && ns[left] == ns[right] {
			radius = expandRadius

			expandRadius = radius + 1 // expand
			left = center - expandRadius
			right = center + expandRadius
		}

		maxPalindrome[center] = radius

		radius = 0
		nextCenter := center + 1
		for nextCenter < right {
			nextCenterMirror := center - (nextCenter - center)
			mirrorRestrictRadius := nextCenterMirror - (left + 1)
			if maxPalindrome[nextCenterMirror] < mirrorRestrictRadius {
				maxPalindrome[nextCenter] = maxPalindrome[nextCenterMirror]
				nextCenter++
			} else if maxPalindrome[nextCenterMirror] > mirrorRestrictRadius {
				maxPalindrome[nextCenter] = mirrorRestrictRadius
				nextCenter++
			} else {
				center = nextCenter
				radius = mirrorRestrictRadius
				break
			}
		}
		// to the range end, reset
		center = nextCenter
	}
	index, maxRadius := findMax(maxPalindrome)
	return shrink(ns[index-maxRadius : index+maxRadius+1])
}
