package main

func sortColors(nums []int) {
	red := 0
	white := 0
	blue := 0

	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}
	counts := []int{red, white, blue}
	k := 0
	for i := 0; i < len(counts); i++ {
		for j := 0; j < counts[i]; j++ {
			nums[k] = i
			k++
		}
	}
}
