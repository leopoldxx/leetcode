package main

func fillNewColor(image [][]int, sr int, sc int, oldColor int, newColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = newColor
	fillNewColor(image, sr-1, sc, oldColor, newColor)
	fillNewColor(image, sr+1, sc, oldColor, newColor)
	fillNewColor(image, sr, sc-1, oldColor, newColor)
	fillNewColor(image, sr, sc+1, oldColor, newColor)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	fillNewColor(image, sr, sc, image[sr][sc], newColor)
	return image
}
