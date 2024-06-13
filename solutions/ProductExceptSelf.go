package solutions

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefix, postfix := 1, 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
