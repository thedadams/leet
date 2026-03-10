func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }

    if prices[1] < prices[0] {
        prices[0] = prices[1]
        prices[1] = 0
    } else {
        prices[1] = prices[1] - prices[0]
    }

    for i := 2; i < len(prices); i++ {
        if prices[i] < prices[0] {
            prices[0] = prices[i]
            continue
        }

        prices[1] = max(prices[i] - prices[0], prices[1])
    }

    return prices[1]
}