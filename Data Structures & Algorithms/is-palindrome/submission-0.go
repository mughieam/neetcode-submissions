func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    l,r := 0, len(s)-1
    for l<r {
        if !(('a' <= s[l] && s[l] <= 'z') || ('0' <= s[l] && s[l] <= '9')) {
            l += 1
        } else if !(('a' <= s[r] && s[r] <= 'z') || ('0' <= s[r] && s[r] <= '9')){
            r -= 1
        } else {
            if string(s[l]) != string(s[r]) {
                return false
            }
            l += 1
            r -= 1
        }
    }
    return true
}