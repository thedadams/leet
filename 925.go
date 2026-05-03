func isLongPressedName(name, typed string) bool {
	if len(name) > len(typed) {
		return false
	}

	var i int
	for _, b := range typed {
		b := byte(b)
		if i < len(name) && name[i] == b {
			i++
		} else if i == 0 && b != name[i] || i > 0 && b != name[i-1] {
			return false
		}
	}

	return i == len(name)
}
