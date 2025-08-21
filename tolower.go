package piscine

func ToLower(s string) string {

	m := ""
	for _, v := range s {

		if v >= 'A' && v <= 'Z' {

			m = m + (string(v + 32))
		} else {

			m = m + (string(v))
		}

	}
	return m
}
