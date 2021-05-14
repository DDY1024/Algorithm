package main

func firstMissingPositive(nums []int) int {
	n, idx := len(nums), 0
	for idx < n {
		if nums[idx] == idx+1 {
			idx++
			continue
		}
		if nums[idx] <= 0 || nums[idx] > n {
			idx++
			continue
		}
		if nums[idx] == nums[nums[idx]-1] {
			idx++
			continue
		}
		nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
	}

	ans := n + 1
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			ans = i + 1
			break
		}
	}
	return ans
}

func main() {
	firstMissingPositive([]int{1, 1})
}
