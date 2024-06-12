package solutions

func LongestConsecutive(nums []int) int {
	var res int
	numSet := make(map[int]bool, len(nums))
	for _, num := range nums {
		numSet[num] = true
	}
	for num := range numSet {
		if numSet[num-1] {
			continue
		}
		sequence := 1
		for numSet[num+sequence] {
			sequence++
		}
		if sequence > res {
			res = sequence
		}
	}
	return res
}
