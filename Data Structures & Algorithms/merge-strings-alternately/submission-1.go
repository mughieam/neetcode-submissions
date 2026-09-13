func mergeAlternately(word1 string, word2 string) string {
    i, res := 0, []byte{}
    for i < len(word1) && i < len(word2) {
        res = append(res, word1[i])
        res = append(res, word2[i])
		i += 1
    }
	if i < len(word1) {
		res = append(res, word1[i:]...)
	}
	if i < len(word2) {
		res = append(res, word2[i:]...)
	}
	return string(res)
}
