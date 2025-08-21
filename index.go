package piscine

func Index(s string, toFind string) int {
for i, v := range s {
	if toFind==(string(v)){
		return i
	}
	
}
return -1

}