func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    
    var prevHasFlower bool
    for i, spot := range flowerbed {
        if spot == 1 {
            prevHasFlower = true
            continue
        }

        if !prevHasFlower && (i == len(flowerbed) - 1 || flowerbed[i+1] == 0) {
            n--
            if n == 0 {
                return true
            }
            prevHasFlower = true
        } else {
            prevHasFlower = false
        }
    }

    return false
}