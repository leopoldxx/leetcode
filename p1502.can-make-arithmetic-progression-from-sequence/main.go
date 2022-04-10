package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func sort2(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func min2from3(a, b, c int) (min1, min2 int) {
	// max is a
	if a > b && a > c {
		return sort2(b, c)
	} else if b > a && b > c {
		return sort2(a, c)
	}
	return sort2(a, b)
}

func sameValueArray(v int, arr []int) bool {
	for _, a := range arr {
		if v != a {
			return false
		}
	}
	return true
}

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}
	max1 := max(arr[0], arr[1])
	min1, min2 := sort2(arr[0], arr[1])
	cache := map[int]int{
		min1: 1,
		min2: 1,
	}

	for i := 2; i < len(arr); i++ {
		cache[arr[i]]++
		min1, min2 = min2from3(arr[i], min1, min2)
		max1 = max(max1, arr[i])
	}

	diff := min2 - min1
	if diff == 0 {
		return sameValueArray(min1, arr)
	}
	j := min1
	for j <= max1 {
		if cache[j] != 1 {
			return false
		}
		delete(cache, j)
		j += diff
	}
	if len(cache) > 0 {
		return false
	}
	return j >= max1
}
