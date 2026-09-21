func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	
	l,r := 0,0
	res := 0
	for l < len(g) && r < len(s) {
        if g[l] <= s[r] {
            res += 1
            l += 1
        }
        r += 1
    }
    return res;
}
