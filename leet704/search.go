package leet704

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if end < start {
		return -1
	}

	middle := start + (end-start)/2
	switch {
	case target == nums[middle]:
		return middle
	case target > nums[middle]:
		return binarySearch(nums, middle+1, end, target)
	default:
		return binarySearch(nums, start, middle-1, target)
	}
}

func search2(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			end = mid - 1
		default:
			start = mid + 1
		}
	}

	return -1
}
