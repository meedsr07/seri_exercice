package piscine

func AlphaCount(s string) int {
	c := 0
	for _, v := range s {
		if v >= 'A' && v <= 'Z'|| v>='a'&& v<='z' {
			c++
		}
	}
	return c

}
