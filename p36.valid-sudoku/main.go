package main

func zIndex(x, y int8) int8 {
	switch {
	case x < 3:
		switch {
		case y < 3:
			return 0
		case y >= 3 && y < 6:
			return 1
		default:
			return 2
		}
	case x >= 3 && x < 6:
		switch {
		case y < 3:
			return 3
		case y >= 3 && y < 6:
			return 4
		default:
			return 5
		}
	default:
		switch {
		case y < 3:
			return 6
		case y >= 3 && y < 6:
			return 7
		default:
			return 8
		}
	}
}

func isValidSudoku(board [][]byte) bool {
	if board == nil {
		return false
	}
	var blank byte = byte('.')
	type coordinate struct {
		x int8
		y int8
		z int8
	}
	var nums [9][]coordinate
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == blank {
				continue
			}
			num := board[x][y] - '1'
			nums[num] = append(nums[num], coordinate{x: int8(x), y: int8(y), z: int8(zIndex(int8(x), int8(y)))})
		}
	}
	for i := range nums {
		var (
			xHash [9]int8
			yHash [9]int8
			zHash [9]int8
		)
		for _, v := range nums[i] {
			xHash[v.x]++
			if xHash[v.x] > 1 {
				return false
			}
			yHash[v.y]++
			if yHash[v.y] > 1 {
				return false
			}
			zHash[v.z]++
			if zHash[v.z] > 1 {
				return false
			}
		}
	}
	return true
}
