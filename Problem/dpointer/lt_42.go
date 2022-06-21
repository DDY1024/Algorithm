package main


// https://leetcode.com/problems/trapping-rain-water/
// 双指针解法的经典应用

func trap(height []int) int {
	left, right, maxL, maxR, ans := 0, len(height)-1, 0, 0, 0
	for left <= right {
		// maxR >= maxL，因此只需要考虑 maxL
		if height[left] <= height[right] { // 这行代码已经暗示了 maxL、maxR 大小关系
			if maxL < height[left] {
				maxL = height[left]
			} else {
				ans += maxL - height[left]
			}
			left++
		} else { // maxL >= maxR，因此只需要考虑 maxR
			if height[right] > maxR {
				maxR = height[right]
			} else {
				ans += maxR - height[right]
			}
			right--
		}
	}
	return ans
}

// 利用双指针去解决问题的思路很巧妙,