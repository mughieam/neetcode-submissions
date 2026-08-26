func lengthOfLastWord(s string) int {
    c := 0
    space := 0
    for i:=len(s)-1; i >= 0; i-- {
        if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
            if space == 0 {
                space += 1
            }
            c += 1
        } else if space > 0 {
            break
        }
    }

    return c
}
