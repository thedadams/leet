import "slices"

func separateDigits(nums []int) []int {
	result := make([]int, 0, len(nums))
	var single []int
	for _, n := range nums {
		for n > 0 {
			single = append(single, n%10)
			n /= 10
		}

		for _, n := range slices.Backward(single) {
			result = append(result, n)
		}

		single = single[:0]
	}

	return result
}
