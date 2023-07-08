package find_the_duplicate_number

func findDuplicate(nums []int) int {
	flags := make([]int, len(nums))

	for _, n := range nums {
		if flags[n] == 1 {
			return n
		}

		flags[n] = 1
	}

	return 0
}
