func validWordAbbreviation(word string, abbr string) bool {
    n,m := len(word),len(abbr)
    i,j := 0,0
    for i < n && j < m {
        if abbr[j] == '0' {
            return false
        }
        if 'a' <= abbr[j] && abbr[j] <= 'z' {
            if word[i] == abbr[j] {
                i += 1
                j += 1
            } else {
                return false
            }
        } else {
            sub := 0
            for j < m && '0' <= abbr[j] && abbr[j] <= '9' {
                sub = sub*10 + int(abbr[j]-'0')
                j += 1
            }
            i += sub
        }
    }
    return i == n && j == m
}