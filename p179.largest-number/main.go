package main

import (
	"sort"
	"strconv"
	"strings"
)

type stringnum []string

func (sn stringnum) Len() int {
	return len(sn)
}
func (sn stringnum) Less(i, j int) bool {
	ai := sn[i]
	aj := sn[j]

	return ai+aj > aj+ai
}
func (sn stringnum) Swap(i, j int) {
	sn[i], sn[j] = sn[j], sn[i]
}

func largestNumber(nums []int) string {
	snums := make([]string, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		snums = append(snums, strconv.Itoa(nums[i]))
	}

	sort.Sort(stringnum(snums))
	if snums[0][0] == '0' {
		return "0"
	}
	sb := strings.Builder{}
	for i := 0; i < len(nums); i++ {
		sb.WriteString(snums[i])
	}
	return sb.String()

}
