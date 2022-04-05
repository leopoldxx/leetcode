package main

func findTarget(src []byte, target byte) int {
	left, right := 0, len(src)-1
	for {
		if left > right {
			return left
		}
		mid := left + (right-left)/2
		if src[mid] == target {
			left = mid + 1
		} else if src[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func nextGreatestLetter(letters []byte, target byte) byte {
	pos := findTarget(letters, target)
	if pos == len(letters) {
		return letters[0]
	}
	if letters[pos] != target {
		return letters[pos]
	}
	if pos == len(letters)-1 {
		return letters[0]
	}
	return letters[pos+1]
}
