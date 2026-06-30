func twoSum(nums []int, target int) []int {
	sum := make(map[int]int)
	for i, n := range nums {
		if j, ok := sum[target-n]; ok {
			return []int{j, i}
		}
		sum[n] = i
	}
	return []int{}
}
