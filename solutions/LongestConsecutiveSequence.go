package solutions

func LongestConsecutive(nums []int) int {
	var longest int
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	for num := range numSet {
		if _, ok := numSet[num-1]; !ok {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}
