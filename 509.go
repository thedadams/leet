func fib(n int) int {
    return fibMem(n, make(map[int]int, n))
}

func fibMem(n int, mem map[int]int) int {
    if n < 2 {
        return n
    }

    if f := mem[n]; f != 0 {
        return f
    }

    mem[n] = fibMem(n - 1, mem) + fibMem(n - 2, mem)
    return mem[n]
}