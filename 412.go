func fizzBuzz(n int) []string {
    result := make([]string, n)
    var ans string
    for i := 1; i <= n; i++ {
        if i % 15 == 0 {
            ans = "FizzBuzz"
        } else if i % 3 == 0 {
            ans = "Fizz"
        } else if i % 5 == 0 {
            ans = "Buzz"
        } else {
            ans = fmt.Sprintf("%d", i)
        }

        result[i - 1] = ans
    }

    return result
}