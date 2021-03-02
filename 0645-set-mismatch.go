func findErrorNums(nums []int) []int {
	var dupl, lost int
	for i := range nums {
		if nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		} else {
			dupl = abs(nums[i])
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			lost = i + 1
			break
		}
	}
	return []int{dupl, lost}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
