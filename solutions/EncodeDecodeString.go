package solutions

import (
	"fmt"
	"strconv"
)

func Encode(strs []string) string {
	var res string
	for _, str := range strs {
		res += fmt.Sprint(len(str)) + "#" + str
	}
	return res
}

func Decode(str string) []string {
	var res []string
	for i := 0; i < len(str); {
		j := i
		if str[j] != '#' {
			j++
		}
		wordLen, _ := strconv.Atoi(str[i:j])
		res = append(res, str[j+1:j+1+wordLen])
		i = j + 1 + wordLen
	}
	return res
}
