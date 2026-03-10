var keyInRow = map[rune]byte {
    'a': 2,
    'b': 3,
    'c': 3,
    'd': 2,
    'e': 1,
    'f': 2,
    'g': 2,
    'h': 2,
    'i': 1,
    'j': 2,
    'k': 2,
    'l': 2,
    'm': 3,
    'n': 3,
    'o': 1,
    'p': 1,
    'q': 1,
    'r': 1,
    's': 2,
    't': 1,
    'u': 1,
    'v': 3,
    'w': 1,
    'x': 3,
    'y': 1,
    'z': 3,
    'A': 2,
    'B': 3,
    'C': 3,
    'D': 2,
    'E': 1,
    'F': 2,
    'G': 2,
    'H': 2,
    'I': 1,
    'J': 2,
    'K': 2,
    'L': 2,
    'M': 3,
    'N': 3,
    'O': 1,
    'P': 1,
    'Q': 1,
    'R': 1,
    'S': 2,
    'T': 1,
    'U': 1,
    'V': 3,
    'W': 1,
    'X': 3,
    'Y': 1,
    'Z': 3,
}

func findWords(words []string) []string {
    var (
        result []string
        row byte
    )
    for _, w := range words {
        row = 4
        for _, r := range w {
            if row == 4 {
                row = keyInRow[r]
            } else if row != keyInRow[r] {
                row = 4
                break
            }
        }
        if row != 4 {
            result = append(result, w)
        }
    }

    return result
}