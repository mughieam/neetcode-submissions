func maxDifference(s string) int {
	chars := make(map[rune]int)
	for _, char := range s {
		chars[char] += 1
	}
	odd, even := 0, math.MaxInt
	for _, val := range chars {
		if val % 2 != 0 {
			odd = max(odd, val)
		} else {
			even = min(even, val)
		}
	}
	return odd - even
}
