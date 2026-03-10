func dividePlayers(skill []int) int64 {
    skills := make(map[int]int, 0)
    max, min := 0, skill[0]
    var chem int64
    for _, i := range skill {
        if max < i {
            max = i
        }
        if min > i {
            min = i
        }
    }

    for _, i := range skill {
        if _, ok := skills[i]; !ok {
            skills[min + max - i]++
        } else {
            if skills[i] == 1 {
                delete(skills, i)
            } else {
                skills[i]--
            }
            chem += int64(i * (min + max - i))
        }
    }

    if len(skills) != 0 {
        return -1
    }
    return chem
}