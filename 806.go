func numberOfLines(widths []int, s string) []int {
    lines, length := 1, 0
    for _, r := range s {
        length += widths[r - 'a']
        if length > 100 {
            lines++
            length = widths[r - 'a']
        }
    }

    return []int{lines, length}
}
