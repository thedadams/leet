var match = map[rune]rune{
    ']': '[',
    ')': '(',
    '}': '{',
}

func isValid(s string) bool {
    st := new(stack)

    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            st.push(c)
        } else if st.pop() != match[c] {
            return false
        }
    }

    return len(st.runes) == 0
}

type stack struct {
    runes []rune
}

func (s *stack) push(b rune) {
    s.runes = append(s.runes, b)
}

func (s *stack) pop() rune {
    if len(s.runes) == 0 {
        return ' '
    }
    b := s.runes[len(s.runes) - 1]
    s.runes = s.runes[:len(s.runes)-1]
    return b
}