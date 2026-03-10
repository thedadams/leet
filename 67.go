func addBinary(a string, b string) string {
	if len(a) < len(b) {
		s := a
		a = b
		b = s
	}

	var (
        carry byte
        result = make([]byte, len(a))
    )
	aOffset := len(a) - len(b)
	for i := len(b) - 1; i >= 0; i-- {
		if a[i+aOffset] == b[i] {
			result[i+aOffset] = '0' + carry
			if b[i] == '1' {
                carry = 1
            } else {
                carry = 0
            }
		} else if carry == 0 {
            result[i+aOffset] = '1'
        } else {
            result[i+aOffset] = '0'
        }
	}

	for i := aOffset - 1; i >= 0; i-- {
        result[i] = a[i] + carry
        if result[i] == '2' {
            result[i] = '0'
        } else {
            carry = 0
        }
	}

	if carry == 1 {
        return string(append([]byte{'1'}, result...))
	}

	return string(result)
}