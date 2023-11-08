package main

func calcFactor(x int) []int {
	factor := make([]int, 0)
	for i := 2; i*i <= x; i += 2 {
		if x%i == 0 {
			for x%i == 0 {
				factor = append(factor, i)
				x /= i
			}
		}
		if i == 2 {
			i--
		}
	}
	if x > 1 {
		factor = append(factor, x)
	}
	return factor
}

// 利用素数筛法求解【最小素因子】
func solveSPF(n int) {
	spf := make([]int, n+1)
	for i := 1; i <= n; i++ {
		spf[i] = i
	}
	for i := 2; i <= n; i += 2 {
		spf[i] = 2
	}
	// 最小 i * i <= n
	for i := 3; i*i <= n; i += 2 {
		if spf[i] == i {
			for j := i * i; j <= n; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}
}

// 求解 [a,b] 范围内约数个数为 n 的数的个数
//
// 方法一
// a. 循环遍历 [a,b] 区间内的每个元素并计算其素因子分解 x = p1^a1 * ... pk^ak
// b. 统计约数个数，如果等于 n，则 cnt++
//
// 方法二
// a. 打表预处理的方法，利用素数筛法的思想求解出每个数的约数个数
//
// 结论
// a. 完全平方数，其约数个数为奇数个
// b. 非完全平方数，其约数个数为偶数个 i, n/i
func calcFactorNum(n int) {
	f := make([]int, n+1)
	for i := 1; i*i <= n; i++ {
		f[i*i]++
		for j := i * (i + 1); j <= n; j += i {
			f[j] += 2
		}
	}
}

// 利用素数筛法求解【最大素因子】
func calcMaxPrimeFactor(n int) {
	f := make([]int, n+1)
	// 素数筛法的思想很实用
	// 最大: i*2 <= n
	for i := 2; i*2 <= n; i++ {
		if f[i] == 0 {
			for j := i; j <= n; j += i {
				f[j] = i
			}
		}
	}
}

// 假设 x 的素因子分解为: x = p1^x1 * p2^x2 * ... * pk^xk
// 约数个数: (1+x1)*(1+x2)*...*(1+xk)
// 约数和: (1+p1^1+...+p1^x1)*(1+p2^1+...+p2^x2) * ... * (1+pk^1+...+pk^xk)
//
// 偶数约数和
// a. 如果 n 的素因子分解不存在 2, 则其所有的约数均为奇数, 偶数约数和为 0
// b. 如果 n 的素因子分解存在 2, 则 2 必然出现所有偶数约数的素因子分解中，因此 sum = (2 + ... + 2^x1) * (1+..+p2^x2) * ... * (1+...+pk^xk)
//
// 奇数约数和
// a. 如果 n 的素因子分解不存在 2, 则其所有约数的约数和即为奇数约数和
// b. 如果 n 的素因子分解存在 2, 则其奇数约数和 = (1+...+p2^x2) * ... * (1+...+pk^xk)，其中 p2, ..., pk 均为奇素数
//
// 正整数 n 的因子分解使得 n = a1 * a2 * ... * ak 且 SUM(a1, a2, ..., ak) 最小
// 解题思路:
// 将正整数 n 进行素因子分解 n = p1^x1 * ... * pk^xk, result = x1*p1 + x2*p2 ... + xk*pk

// 题目: 求解正整数 x 的所有约数的乘积
// https://www.geeksforgeeks.org/product-factors-number/
// 假设 x = p1^x1 * p2^x2 * ... * pk^xk
// num(x) = (1+x1) * (1+x2) * ... * (1+xk)
// f(p1) = p1^[(1+x2)*...(1+xk)] * (p1^2)^[(1+x2)*...(1+xk)] * ... (p1^x1)^[(1+x2)*...(1+xk)] = p1^(1+2+...+x1)^[(1+x2)*...*(1+xk)]
// result = f(p1) * f(p2) * ... * f(pk) % mod
// 总之，求解每个素因子的指数幂在所有约数乘积中出现的次数，从而进一步求解出所有约数的乘积
// 如果牵扯到取余运算，则直接快速幂求解即可

// 题目：求解 n! 约数个数
// 解题思路
// a. 素数筛法求解出 [2, n] 范围内的所有约数
// b. 计算 n! 中含有某一个素因子的幂次 calcNum(n!, p)
// c. res *= (num + 1)

// Smith Number
// sum(number digit) = sum(prime factor digit)
// https://www.geeksforgeeks.org/smith-number/

// Sphenic Number
// 性质: 三个不同素因子的乘积，每个素因子恰好出现一次
// https://www.geeksforgeeks.org/sphenic-number/

// Hoax Number
// sum(number digit) = sum(different prime factor digit)
// https://www.geeksforgeeks.org/hoax-number/

// 大整数素因子分解算法
// 算法原理
// a. Pollard's Rho 算法快速找到大整数 n 的一个因子
// b. 米勒拉宾算法快速判断 n 的因子 x 是否为素数
// 使用情况: a. a large composite number b. small prime factors
/*
// 注: 参数 c 的取值是任意的, 通常取一个随机的整数即可. 例如, 201
// https://blog.csdn.net/Sunshine_cfbsl/article/details/52512706
// https://www.geeksforgeeks.org/pollards-rho-algorithm-prime-factorization/

const int C = 201;
int pcnt;
int res[maxn];

int64 random(int64 n) {
    return (int64)((double)rand() / RAND_MAX * n + 0.5);
}

int64 mul_mod(int64 a, int64 b, int64 c) {
    int64 ans = 0;
    while(b) {
        if(b & 1) ans = (ans + a) % c;
        a = 2 * a % c;
        b >>= 1;
    }
    return ans;
}

int64 pow_mod(int64 a, int64 b, int64 c) {
    int64 ans = 1;
    while(b) {
        if(b & 1) ans = mul_mod(ans, a, c);
        a = mul_mod(a, a, c);
        b >>= 1;
    }
    return ans;
}

int64 pollard_rho(int64 n, int64 c) {
    int64 x, y, d, i = 1, k = 2;
    x = random(n - 1) + 1;
    y = x;
    while(1) {
        ++i;
        x = (mul_mod(x, x, n) + c) % n;
        d = gcd(y - x + n, n);
        if(d > 1 && d < n) return d;
        if(y == x) return n;
        if(i == k) {
            y = x;
            k *= 2;
        }
    }
    return n;
}

void find_prime(int64 n, int k) {
    if(n == 1) return ;
    if(miller(n)) {
        res[pcnt++] = n;  // 素因子 + 1
        return ;
    }
    int64 p = n;
    while(p >= n) {
        p = pollard_rho(p, k - 1);
	}
	// a. 快速求解小于 n 的一个约数, 然后递归求解 p, n/p
	// b. 递归的边界条件: n 为一个素数（米勒拉宾素数测试）
    find_prime(p, k);
    find_prime(n / p, k);
    return ;
}
*/

// Find politeness of a number
// https://www.geeksforgeeks.org/find-politeness-number/
// 求解正整数 n 可以被连续整数 x + (x+1) + ... + (x+k) 表示的方案数
//
// 方法一
// 枚举 [0, n/2] 的起始整数，通过计算等差数列的和，判断是否存在和等于 n。
// 由于等差数列的和是单调递增的，因此此处我们采用二分的方法进行判定，时间复杂度为 O(NlogN)
//
// 方法二
// 存在一个证明：上述方案数等于 n 的所有奇约数的个数，其中不包括 1。
// 证明参考: https://en.wikipedia.org/wiki/Polite_number#Construction_of_polite_representations_from_odd_divisors
// 对于 n 的奇约数个数的求解，我们通过素因子分解的方法进行求解
// 假设 n 的素因子分解为 n = p1^x1 * ... p2^x2 * ... * pk^xk
// 不包括 1 的奇约数的个数 = (x1 + 1) * ... * (xk + 1) - 1，其中满足 p1, ... , pk 均为奇素因子(蠢话，大于 2 的素数均为奇数)
// 奇数 * 偶数 = 偶数
// 偶数 * 偶数 = 偶数
// 奇数 * 奇数 = 奇数

// 求解 n! 的素因子分解中素数 p 出现的次数
// x1*p^1 + x2*p^2 + ... + xk*p^k = p * (x1 + x2*p + ... + xk*p^(k-1))
// https://www.geeksforgeeks.org/finding-power-prime-number-p-n/
//
// 扩展: 求解 n! 包含 x 的幂次，其中 x 不一定是素数
// a. 将 x 进行素因子分解 x = p1^a1 * ... * pk^ak
// b. 针对每个素因子 p1, p2, ... , pk 进行计算
// c. ans = min(ans, calc(pi)/ai)
// 最终求解出 x 在 n! 中出现的幂次数
func powerOfPrimeInFactorial(n, p int) int {
	ans := 0
	for n > 0 {
		ans += n / p
		n /= p
	}
	return ans
}

// 朴素素数筛法：O(n * logn * logn)
func doPrimeList(n int) {
	mark := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !mark[i] {
			for j := i * i; j <= n; j += i {
				mark[j] = true
			}
		}
	}
}

// 线性筛法：每一个合数只被其“最小质因数”筛到
func doLinearPrimeList(n int) {
	mark := make([]bool, n+1)
	plist := make([]int, 0, n)
	for i := 2; i <= n; i++ {
		if !mark[i] {
			plist = append(plist, i)
		}

		for _, p := range plist {
			if p*i > n {
				break
			}
			mark[p*i] = true
			// i % p == 0
			// i * next_p 的最小素因子是 p，而不是 next_p，不满足我们只能被最小素因子 p 筛到的要求
			if i%p == 0 {
				break
			}
		}
	}
}
