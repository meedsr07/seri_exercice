package piscine

func ToUpper(s string) string {
	m := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			m = m + (string(v - 32))
			// m = m + (string(40))
		} else {
			m = m + (string(v))

		}

	}
	return m

}
