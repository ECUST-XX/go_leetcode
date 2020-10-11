package leet1

func TwoSum(nums []int, target int) []int {
	for k := range nums {
		for i := k + 1; i < len(nums); i++ {
			if nums[k]+nums[i] == target {
				return []int{k, i}
			}
		}

	}
	return []int{0, 0}
}
