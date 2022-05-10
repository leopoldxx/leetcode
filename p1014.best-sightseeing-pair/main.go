package main

func maxScoreSightseeingPair(values []int) int {
	preMaxSpotValue := values[0] + 0
	maxScore := -1

	for i := 1; i < len(values); i++ {
		value1 := values[i] + i // act as the first spot
		value2 := values[i] - i // act as the second spot

		if preMaxSpotValue+value2 > maxScore {
			maxScore = preMaxSpotValue + value2
		}

		if value1 > preMaxSpotValue {
			preMaxSpotValue = value1
		}

	}
	return maxScore

}
