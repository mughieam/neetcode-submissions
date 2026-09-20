func canConstruct(ransomNote string, magazine string) bool {
	for _, val := range ransomNote {
        char := string(val)
        if strings.Count(magazine, char) >= strings.Count(ransomNote, char) {
            continue
        } else {
            return false
        }
    }
    return true
}
