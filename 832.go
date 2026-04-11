func flipAndInvertImage(image [][]int) [][]int {
	l := len(image[0])
	for _, row := range image {
		for i := range (l + 1) / 2 {
			row[i], row[l-i-1] = (row[l-i-1]+1)%2, (row[i]+1)%2
		}
	}

	return image
}
