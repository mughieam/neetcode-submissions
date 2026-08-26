func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	ps,pt := 0,0
	for ps < len(s) && pt < len(t) {
		if s[ps] == t[pt] {
			ps += 1
			pt += 1
		} else {
			pt += 1
		}
	}
	if ps == len(s) {
		return true
	}
	return false
}
