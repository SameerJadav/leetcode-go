package solutions

func GroupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, char := range str {
			count[char-'a']++
		}
		anagramsMap[count] = append(anagramsMap[count], str)
	}
	res := make([][]string, len(anagramsMap))
	var i int
	for _, str := range anagramsMap {
		res[i] = str
		i++
	}
	return res
}
