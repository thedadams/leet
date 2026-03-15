func calPoints(operations []string) int {
	var records []int
	for _, s := range operations {
		switch s {
		case "D":
			records = append(records, 2*records[len(records)-1])
		case "C":
			records = records[:len(records)-1]
		case "+":
			records = append(records, records[len(records)-1]+records[len(records)-2])
		default:
			i, _ := strconv.Atoi(s)
			records = append(records, i)
		}
	}

	var result int
	for _, i := range records {
		result += i
	}

	return result
}
