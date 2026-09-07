func minOperations(logs []string) int {
	dir := []string{}
	for _, log := range logs {
		if log == "../" && len(dir) > 0 {
			dir = dir[:len(dir)-1]
		} else if log == "./" {
			continue
		} else if log != "../" {
			dir = append(dir, log)
        }
	}
	return len(dir)
}
