func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := []int{}
    for _, num := range nums1 {
        j := 0
		for i, _ := range nums2 {
			if num == nums2[i] {
				j = i
				break
			}
		}
		for j < len(nums2) {
			if num < nums2[j] {
				res = append(res, nums2[j])
				break
			}
			j++
		}
		if j == len(nums2) {
			res = append(res, -1)
		}
    }
    return res
}
