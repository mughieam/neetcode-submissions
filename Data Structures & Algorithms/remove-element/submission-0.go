func removeElement(nums []int, val int) int {
    s,w := 0,0
	for s < len(nums) {
		if nums[s] != val {
			nums[w] = nums[s]
			w += 1
		}
		s += 1
	}
	return w
}
