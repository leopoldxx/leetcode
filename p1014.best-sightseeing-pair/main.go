package main

func maxScoreSightseeingPair(values []int) int {
	maxPre := values[0]
	maxTotal := values[0]

	for i := 1; i < len(values); i++ {
		if maxPre+values[i]-i > maxTotal {
			maxTotal = maxPre + values[i] - i
		}

		if values[i]+i > maxPre {
			maxPre = values[i] + i
		}

	}
	return maxTotal

}
