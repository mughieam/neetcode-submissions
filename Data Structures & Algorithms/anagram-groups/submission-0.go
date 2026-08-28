func sorting(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		s := sorting(str)
		anagrams[s] = append(anagrams[s], str)
	}
	res := make([][]string, len(anagrams))
	i := 0
	for _, anagram := range anagrams {
		res[i] = append(res[i], anagram...)
		i += 1
	}
	return res
}
