package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if !(IsLower(string(v))) && !IsUpper(string(v)) && !IsNumeric(string(v)) {
			return false
		}

	}
	return true

}

//func IsAlpha(s string) bool {
//	for _, v := range s {
//		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') && (v < '0' || v > '9') {
//			return false
//		}
//
//	}
//	return true
//
//}
