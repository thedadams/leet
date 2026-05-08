func validMountainArray(arr []int) bool {
	var decreasing bool
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			if decreasing {
				return false
			}
		} else if arr[i-1] > arr[i] {
			if i == 1 {
				return false
			}
			decreasing = true
		} else {
			return false
		}
	}

	return decreasing
}
