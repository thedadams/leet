func findComplement(num int) int {
    powOf2 := 1
    for powOf2 <= num {
        powOf2 <<= 1
    }

    return num ^ (powOf2 - 1)
}