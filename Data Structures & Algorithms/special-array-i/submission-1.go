func isArraySpecial(nums []int) bool {
    i := 0
	for i < len(nums)-1 {
		if nums[i] % 2 == nums[i+1] % 2 {
			return false
		}
		i += 1
	}
	return true
}