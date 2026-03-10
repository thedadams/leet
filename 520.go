func detectCapitalUse(word string) bool {
    allLowerCase := word[0] >= 'a' && word[0] <= 'z'
    var hasLowerCase *bool

    for i := 1; i < len(word); i++ {
        if allLowerCase && word[i] >= 'A' && word[i] <= 'Z' {
            return false
        } else if !allLowerCase {
            if hasLowerCase == nil {
                hasLowerCase = new(bool)
                if word[i] >= 'a' && word[i] <= 'z' {
                    *hasLowerCase = true
                }
            }

            if *hasLowerCase && word[i] >= 'A' && word[i] <= 'Z' {
                return false
            } else if !*hasLowerCase && word[i] >= 'a' && word[i] <= 'z' {
                return false
            }
        }
    }

    return true
}