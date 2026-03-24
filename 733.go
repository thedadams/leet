func floodFill(image [][]int, sr, sc, color int) [][]int {
    if original := image[sr][sc]; original != color {
        fill(image, sr, sc, original, color)
    }
    return image
}

func fill(image [][]int, sr, sc, original, color int) {
    image[sr][sc] = color
    for _, d := range [][2]int{{sr, sc - 1}, {sr - 1, sc}, {sr, sc + 1}, {sr + 1, sc}} {
        if d[0] < 0 || d[1] < 0 || d[0] >= len(image) || d[1] >= len(image[d[0]]) || image[d[0]][d[1]] != original {
            continue
        }

        fill(image, d[0], d[1], original, color)
    }
}
