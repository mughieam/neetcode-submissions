func calPoints(operations []string) int {
	scores := []int{}
	for _, op := range operations {
		if op == "C" {
			scores = scores[:len(scores)-1]
		} else if op == "D" {
			double := scores[len(scores)-1] * 2
			scores = append(scores, double)
		} else if op == "+" {
			prev := scores[len(scores)-2] + scores[len(scores)-1]
			scores = append(scores, prev)
		} else {
            num, _ := strconv.Atoi(op)
			scores = append(scores, num)
        }
	}
	sum := 0
	for _, score := range scores {
        sum += score
    }
	return sum
}
