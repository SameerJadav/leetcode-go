package solutions

func TwoSum(nums []int, target int) []int {
	numIdxMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numIdxMap[target-num]; ok {
			return []int{j, i}
		}
		numIdxMap[num] = i
	}
	return nil
}
