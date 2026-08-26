func lengthOfLastWord(s string) int {
    word := []byte{}
    for i:=len(s)-1; i>=0; i-- {
        if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
            word = append(word, s[i])
        } else if len(word) > 0 {
            break
        }
    }
    return len(word)
}
