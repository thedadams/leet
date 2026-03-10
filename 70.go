func climbStairs(n int) int {
    return climbStairsMem(n, make(map[int]int, n))
}

func climbStairsMem(n int, mem map[int]int) int {
    if n <= 0 {
        return 0
    }
    if n <= 3 {
        return n
    }
    
    if ans := mem[n]; ans != 0 {
        return ans
    }

    mem[n] = climbStairsMem(n-1, mem) + climbStairsMem(n-2, mem)
    return mem[n]
}