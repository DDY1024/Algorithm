package main

// 指数循环节问题总结
// 参考: https://blog.csdn.net/ACdreamers/article/details/8236942

// 指数取模运算降幂公式: a^b % c = a^(b%phi(c)+phi(c)) % c，其中 b >= phi(c)

// 1. http://acm.fzu.edu.cn/problem.php?pid=1759
func sloveMod(a, b, c int) int {

	var powMod = func(x, y, n int) int {
		x = x % n
		ans := 1
		for y > 0 {
			if y&1 > 0 {
				ans = ans * x % n
			}
			x = x * x % n
			y >>= 1
		}
		return ans
	}

	// var mulMod = func(a, b, c int) int {
	// 	res := 0
	// 	a %= c
	// 	for b > 0 {
	// 		if b&1 > 0 {
	// 			res = (res + a) % c
	// 		}
	// 		a = a * 2 % c
	// 		b >>= 1
	// 	}
	// 	return res
	// }

	var calcPhi = func(x int) int {
		res := x
		for i := 2; i*i <= x; i += 2 {
			if x%i == 0 {
				res = res / i * (i - 1)
			}
			for x%i == 0 {
				x /= i
			}
			if i == 2 {
				i--
			}
		}
		if x > 1 {
			res = res / x * (x - 1)
		}
		return res
	}

	// 注意: 防止 calcPhi 重复计算，此处只是为了方便表示
	return powMod(a, b%calcPhi(c)+calcPhi(c), c)
}
