func isOneBitCharacter(bits []int) bool {
    for i := 0; i < len(bits); i++ {
        if bits[i] == 1 {
            i++
        } else if i == len(bits) - 1 {
            return true
        }
    }

    return false
}
