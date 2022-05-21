package main

func convertToTitle(columnNumber int) string {
	res := []byte{}
	for columnNumber > 0 {
		div := columnNumber / 26
		diff := columnNumber - div*26
		columnNumber = div
		if diff == 0 {
			res = append([]byte{'Z'}, res...)
			columnNumber--
		} else {
			res = append([]byte{'A' + byte(diff-1)}, res...)
		}
	}
	return string(res)
}
