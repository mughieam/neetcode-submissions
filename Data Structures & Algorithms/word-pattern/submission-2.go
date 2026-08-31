func wordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)
	for i:=0; i<len(words); i++ {
		char := pattern[i]
		word := words[i]
		
		if val, exists := charToWord[char]; exists && val != word {
			return false
		}
		if val, exists := wordToChar[word]; exists && val != char {
			return false
		}

		wordToChar[word] = char
		charToWord[char] = word
	}
	return true
}