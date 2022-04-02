package main

func average(salary []int) float64 {
	count := len(salary)
	low := salary[0]
	hight := salary[0]
	total := 0
	for i := 0; i < len(salary); i++ {
		if salary[i] < low {
			low = salary[i]
		} else if salary[i] > hight {
			hight = salary[i]
		}
		total += salary[i]
	}
	return float64(total-low-hight) / float64(count-2)
}
