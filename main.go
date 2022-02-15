package main

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if mid&1 > 0 {
			if nums[mid] != nums[mid-1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if mid-1 >= 0 && nums[mid-1] == nums[mid] {
				r = mid - 2
			} else if mid+1 < n && nums[mid+1] == nums[mid] {
				l = mid + 2
			} else {
				return nums[mid]
			}
		}
	}
	return -1
}
