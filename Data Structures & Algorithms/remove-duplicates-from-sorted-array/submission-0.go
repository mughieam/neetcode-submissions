func removeDuplicates(nums []int) int {
	for i:=0; i+1<len(nums); i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i+1], nums[i+2:]...)
			i -= 1
		}
	}
	return len(nums)
}
