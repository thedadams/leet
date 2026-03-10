func checkRecord(s string) bool {
    var late, absent int
    for _, c := range s {
        switch c {
        case 'L':
            late++
            if late == 3 {
                return false
            }

            continue

            case 'A':
            absent++
            if absent == 2 {
                return false
            }
        }

        late = 0
    }

    return true
}