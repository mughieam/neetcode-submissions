func isPalindrome(s string) bool {
	l,r := 0,len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func validPalindrome(s string) bool {
	l,r := 0,len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s[l:r]) || isPalindrome(s[l+1:r+1])
		}
		l += 1
		r -= 1
	}
	return true
}
