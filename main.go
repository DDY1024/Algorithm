package main

import (
	"strconv"
	"strings"
)

func main() {

}

// 格里高利历闰年计算方法: https://zh.wikipedia.org/wiki/%E9%97%B0%E5%B9%B4
func dayOfYear(date string) int {
	days := [][]int{
		{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
		{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	}

	// 闰年的判断不止这么简单
	var isLeapYear = func(x int) int {
		if (x%100 == 0 && x%400 == 0) || x%4 == 0 {
			return 1
		}
		return 0
	}

	arr := strings.Split(date, "-")
	year, _ := strconv.ParseInt(arr[0], 10, 64)
	month, _ := strconv.ParseInt(arr[1], 10, 64)
	day, _ := strconv.ParseInt(arr[2], 10, 64)

	ans := 0
	for i := 1; i < int(month); i++ {
		ans += days[isLeapYear(int(year))][i]
	}
	ans += int(day)
	return ans
}

// 闰年 2 月 29 天
// 平年 2 月 28 天
