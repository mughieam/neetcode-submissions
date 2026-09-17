func maxNumberOfBalloons(text string) int {
    balloons := map[string]int{
		"b": 0,
		"a": 0,
		"l": 0,
		"o": 0,
		"n": 0,
	}
	for _, val := range text {
		char := string(val)
		if _, ok := balloons[char]; ok {
			balloons[char] += 1
		}
	}
	return min(balloons["b"], balloons["a"], int(balloons["l"]/2), int(balloons["o"]/2), balloons["n"])
}