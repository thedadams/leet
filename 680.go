func validPalindrome(s string) bool {
	offIdx := -1
	for i := range len(s) / 2 {
		if s[i] != s[len(s)-i-1] {
			offIdx = i
			break
		}
	}

	// The string is already a palindrome.
	if offIdx == -1 {
		return true
	}

    palindrome := true
    // Check if the string is a palindrome by skipping the character at offIdx.
    for i := offIdx + 1; i < len(s) / 2 + 1; i++ {
        if s[i] != s[len(s)-i] {
            palindrome = false
			break
		}
    }

    if palindrome {
        return true
    }

    palindrome = true
    // Check if the string is a palindrome by skipping the character at len(s) - offIdx - 1.
    for i := offIdx; i < len(s) / 2; i++ {
        if s[i] != s[len(s)-i-2] {
            palindrome = false
			break
		}
    }
    return palindrome
}
