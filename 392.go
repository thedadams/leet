func isSubsequence(s, t string) bool {
    if s == "" {
        return true
    }

    var ids, idt int

    for ids < len(s) {
        for idt < len(t) && s[ids] != t[idt] {
            idt++
        }


        idt++
        if idt > len(t) {
            break
        }

        ids++
    }

    return ids == len(s) && idt <= len(t)
}