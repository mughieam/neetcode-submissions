func moveZeroes(nums []int) {
	s,w := 0,0
	for s < len(nums) {
		if nums[s] != 0 {
			nums[s], nums[w] = nums[w], nums[s]
			w += 1
		}
		s += 1
	}
}
