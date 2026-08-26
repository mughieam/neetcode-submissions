func countSeniors(details []string) int {
	people := 0
	for _, detail := range details {
		age := detail[11:13]
		if strings.Compare(age, "60") == 1 {
			people += 1
		}
	}
	return people
}