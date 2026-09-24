func plusOne(digits []int) []int {
    if len(digits) == 0 {
		return []int{1}
	}
	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1] += 1
		return digits
	} else {
		return append(plusOne(digits[:len(digits)-1]), 0)
	}
}
