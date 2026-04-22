func fairCandySwap(aliceSizes, bobSizes []int) []int {
	var diff int
	aliceSet := make(map[int]struct{}, len(aliceSizes))
	for _, i := range aliceSizes {
		aliceSet[i] = struct{}{}
		diff += i
	}
	for _, i := range bobSizes {
		diff -= i
	}

	for _, i := range bobSizes {
		newDiff := diff + 2*i
		if _, ok := aliceSet[newDiff/2]; ok {
			return []int{newDiff / 2, i}
		}
	}

	return nil
}
