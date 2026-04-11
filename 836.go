func isRectangleOverlap(rect1 []int, rect2 []int) bool {
	return min(rect1[3], rect2[3]) > max(rect1[1], rect2[1]) &&
		min(rect1[2], rect2[2]) > max(rect1[0], rect2[0])
}
