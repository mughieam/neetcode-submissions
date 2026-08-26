func scoreOfString(s string) int {
	l,r := 0,1
	sum := 0
	for r <= len(s)-1 {
		if int(s[l]) > int(s[r]) {
			sum += (int(s[l]) - int(s[r]))
		} else {
			sum += (int(s[r]) - int(s[l]))
		}
		l += 1
		r += 1
	}
	return sum
}
