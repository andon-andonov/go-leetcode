func distributeCandies(candyType []int) int {
	set := make(map[int]bool)
	for i := range candyType {
		set[candyType[i]] = true
	}
	return min(len(set), len(candyType)/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
