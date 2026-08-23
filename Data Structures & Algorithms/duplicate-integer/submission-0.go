func hasDuplicate(nums []int) bool {
    hash := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := hash[num]; !exist {
			hash[num] = struct{}{} 
		} else {
			return true
		}
	}
	return false
}
