package letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	digitsMap := map[uint8][]string{
		'0': {"0"},
		'1': {"1"},
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var result []string
	if digits == "" {
		return result
	}

	backtrack("", 0, digits, &result, digitsMap)

	return result
}

func backtrack(current string, index int, digits string, result *[]string, digitMap map[uint8][]string) {
	if len(digits) == index {
		*result = append(*result, current)
		return
	}

	for _, ch := range digitMap[digits[index]] {
		backtrack(current+ch, index+1, digits, result, digitMap)
	}
}
