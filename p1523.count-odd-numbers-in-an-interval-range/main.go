package main

func countOdds(low int, high int) int {
	rang := high - low + 1
	if rang%2 == 0 {
		return rang / 2
	}
	return rang/2 + low%2
}
