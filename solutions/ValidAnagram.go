package solutions

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countS := make(map[byte]int)
	countT := make(map[byte]int)
	for i := range len(s) {
		countS[s[i]]++
		countT[t[i]]++
	}
	for char := range countS {
		if countS[char] != countT[char] {
			return false
		}
	}
	return true
}
