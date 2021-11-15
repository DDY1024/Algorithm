package main

import (
	"fmt"
)

func main() {
	fmt.Println(^-1)
	fmt.Println(^0)
	fmt.Println("Hello, World!")
	// for i := 0; ; i-- {

	// }
}

// C++ Version
// 权值树状数组查询第k小
// int kth(int k) {
// 	int cnt = 0, ret = 0;
// 	for (int i = log2(n); ~i; --i) {      // i 与上文 depth 含义相同
// 	  ret += 1 << i;                      // 尝试扩展
// 	  if (ret >= n || cnt + t[ret] >= k)  // 如果扩展失败
// 		ret -= 1 << i;
// 	  else
// 		cnt += t[ret];  // 扩展成功后 要更新之前求和的值
// 	}
// 	return ret + 1;
// }
