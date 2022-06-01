package main

const (
	MASK11 = 0x80
	MASK21 = 0xE0
	MASK31 = 0xF0
	MASK41 = 0xF8
	MASKN  = 0xC0

	FLAG1 = 0x0
	FLAG2 = 0xC0
	FLAG3 = 0xE0
	FLAG4 = 0xF0
	FLAGN = 0x80
)

func nextN(data []int, index *int, n int) bool {
	for i := 0; i < n; i++ {
		if *index == len(data) {
			return false
		}
		b := byte(data[*index])
		if b&MASKN != FLAGN {
			return false
		}
		*index = *index + 1
	}
	return true
}

func validUtf8(data []int) bool {
	var index int

	for index < len(data) {
		b1 := byte(data[index])
		switch {
		case b1&MASK11 == FLAG1:
			index++
		case b1&MASK21 == FLAG2:
			index++
			if !nextN(data, &index, 1) {
				return false
			}
		case b1&MASK31 == FLAG3:
			index++
			if !nextN(data, &index, 2) {
				return false
			}
		case b1&MASK41 == FLAG4:
			index++
			if !nextN(data, &index, 3) {
				return false
			}
		default:
			return false
		}

	}
	return true
}
