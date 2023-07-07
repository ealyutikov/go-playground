package product_of_array_except_self

// https://leetcode.com/problems/product-of-array-except-self
func productExceptSelf(nums []int) []int {
	product := 1
	left := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		left[i] = product * nums[i]
		product *= nums[i]
	}

	product = 1
	rigth := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		rigth[i] = product * nums[i]
		product *= nums[i]
	}

	result := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		var l, r int = 1, 1

		if i-1 >= 0 {
			l = left[i-1]
		}

		if i+1 < len(nums) {
			r = rigth[i+1]
		}

		result[i] = l * r
	}

	return result
}
