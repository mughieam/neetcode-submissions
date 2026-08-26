func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, exist := seen[diff]; exist {
			return []int{j,i}
		}
		seen[num] = i
	}
	return []int{0,0}
}
