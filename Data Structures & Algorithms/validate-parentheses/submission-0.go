func isOpen(s rune) bool {
    switch s {
        case '(', '{', '[':
            return true
        default:
            return false
    }
}

func isValidClosure(o rune, c rune) bool {
    switch o {
        case '(':
            return c == ')'
        case '{':
            return c == '}'
        case '[':
            return c == ']'
        default:
            return false
    }
}

func isValid(s string) bool {
    runes := []rune(s)
    bracket := []rune{}
    for _, r := range runes {
        if isOpen(r) {
            bracket = append(bracket, r)
            continue
        }

        if len(bracket) > 0 && isValidClosure(bracket[len(bracket)-1], r) {
            bracket = bracket[:len(bracket)-1]
        } else {
            return false
        }
    }

    return len(bracket) == 0
}