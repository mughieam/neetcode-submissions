func asteroidCollision(asteroids []int) []int {
	res := []int{}
	for _, ast := range asteroids {
		explode := false
		for !explode && len(res) > 0 && res[len(res)-1] > 0 && ast < 0 {
			if res[len(res)-1] == -ast {
				res = res[:len(res)-1]
				explode = true
			} else if res[len(res)-1] < -ast {
				res = res[:len(res)-1]
			} else {
				explode = true
			}
		}
		if !explode {
			res = append(res, ast)
		}
	}
	return res
}