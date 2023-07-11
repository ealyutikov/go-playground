package interval_list_Intersections

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var i, j int = 0, 0

	result := [][]int{}
	for i < len(firstList) && j < len(secondList) {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		if start <= end {
			result = append(result, []int{start, end})
		}

		if firstList[i][1] < secondList[j][1] {
			i += 1
		} else {
			j += 1
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
