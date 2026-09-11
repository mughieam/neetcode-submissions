func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		row := []int{1}
		val := 1
		for j:=1; j<=i; j++ {
			val = val * (i - j + 1) / j
			row = append(row, val)
		}
		res = append(res, row)
	}
	return res
}
