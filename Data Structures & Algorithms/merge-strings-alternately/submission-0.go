func mergeAlternately(word1 string, word2 string) string {
	maxLen := max(len(word1), len(word2))
    res := make([]byte, 0, maxLen)
    for i:=0; i<maxLen; i++ {
        if i < len(word1) {
            res = append(res, word1[i])
        }
        if i < len(word2) {
            res = append(res, word2[i])
        }
    }
	return string(res)
}
