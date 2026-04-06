import "math"

func largestTriangleArea(points [][]int) float64 {
	var maxArea float64
	for i := range points {
		for j := range len(points) - i {
			for k := range len(points) - j {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j+i][0], points[j+i][1]
				x3, y3 := points[k+j][0], points[k+j][1]
				maxArea = max(maxArea, 0.5*math.Abs(float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))))
			}
		}
	}
	return maxArea
}
