package main

import (
	"fmt"
	"math"
)

// 1. 快速幂计算
// x^y % n --> (x%n)^y % n
func powMod(x, y, n int) int {
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

// 2. 快速乘法取模
// a * b % c
func mulMod(a, b, c int) int {
	res := 0
	a %= c
	for b > 0 {
		if b&1 > 0 {
			res = (res + a) % c
		}
		a = a * 2 % c
		b >>= 1
	}
	return res
}

// 3. 给定 n、p，其中 p 为素数，判断是否存在 x 使得 (x * x) % p = n % p
// 判断方法：n^[(p-1)/2]%p = 1
// https://www.geeksforgeeks.org/eulers-criterion-check-if-square-root-under-modulo-p-exists/

// 5. big_number % x
// big_num 字符串表示
// https://www.geeksforgeeks.org/how-to-compute-mod-of-a-big-number/
/*
int mod(string num, int a)
{
    // Initialize result
    int res = 0;

    // One by one process all digits of 'num'
    for (int i = 0; i < num.length(); i++)
         res = (res*10 + (int)num[i] - '0') %a;

    return res;
}
*/

// 6. 费马小定理
// a^p = a (mod p)，其中 p 为素数
// (1) 如果 a 为 p 的倍数，a^p = a (mod p) 为 0
// (2) 如果 a 不为 p 的倍数，a^(p-1) = 1 (mod p)

// 7. 离散对数问题
// a^k = b (mod m)，求解 k 使得上述同余方程成立, 其中 gcd(a, m) = 1
// 求解该问题通常采用 Baby-step giant-step algorithm 算法
// 算法原理参考：https://www.geeksforgeeks.org/discrete-logarithm-find-integer-k-ak-congruent-modulo-b/
// 如果 k 在 [0, m) 区间内不存在解，则该同余方程不存在解。离散对数的求解算法便是将 k 换一种表示方式，通过 i * n - j
// 其中 i 在 [1, n) 区间内， j 在 [0, n) 区间内, n = ceil(sqrt(m))
// 求解 a^k = b (mod m)
// 注意: a 和 m 必须互素，即 gcd(a, m) = 1
func discreteLog(a, b, m int) int {
	// n := int(math.Sqrt(float64(m))) + 1
	n := int(math.Ceil(math.Sqrt(float64(m))))
	mark := map[int]int{}
	for i := 1; i <= n; i++ { // LHS: a^(i*n) % m
		x := powMod(a, i*n, m)
		if _, ok := mark[x]; !ok { // 注意: 此处要进行判重操作
			mark[x] = i
		}
	}
	for i := 0; i < n; i++ {
		cur := powMod(a, i, m) * b % m // a^j * b % m , 如果此处发生整数溢出的情况，则乘法操作的取余要单独处理
		if _, ok := mark[cur]; ok {
			res := mark[cur]*n - i
			// fmt.Println(mark[cur], n, i)
			if res < m { // < m 判断必须存在
				return res
			}
		}
	}
	return -1
}

// 8. 给定一个正整数数组 arr[0], arr[1], ... , arr[n-1], 找出所有的正整数 k ，使得
// arr[0] % k = arr[1] % k = ... = arr[n-1] % k
// https://www.geeksforgeeks.org/finding-k-modulus-array-element/
// 解题思路
// 方法一：
// 枚举所有可能的约数进行验证，首先我们容易知道 a = b (mod c)，则 (a-b) 是 c 的倍数
// 即 c 是 (a-b) 的约数。因此寻找数组的 最小元素(minElem) 和 最大元素 (maxElem), 计算
// 两者差值的所有约数，进行枚举验证即可。
// 方法二：
// a = b (mod c) --> (a-b) % c == 0
// (1) 枚举所有的 (i, j) 对，求解两者差值的绝对值，构造相应序列，求解该序列的最大公约数
// (2) 求解该 gcd 的所有约数即可

// 9. 正整数 n 分解成若干正整数的和且使得这些正整数的乘积最大
// https://www.geeksforgeeks.org/breaking-integer-to-get-maximum-product/
// 引理：正整数 n 分解成两个数之和且使得其乘积最大。容易知道，如果两个数十分接近，则获取到的乘积是最大的。即 x = n/2
// 在该问题中，同样如果分解成的若干正整数 x 比较相近的话，则最后的乘积则是比较大。问题转化成求解 x^(n/x) 的最大值，易
// 其导数为 0 点，即为最大值点。 x^(n/x) 求导方法可以参考 https://zhidao.baidu.com/question/1800647590078701987.html
// y = x^(n/x) 极值点为 e，其中 2 < e < 3 。
/*
int breakInteger(int N)
{
    //  base case 2 = 1 + 1
    if (N == 2)
        return 1;

    //  base case 3 = 2 + 1
    if (N == 3)
        return 2;

    int maxProduct;

    //  breaking based on mod with 3
    switch (N % 3)
    {
        // If divides evenly, then break into all 3
        case 0:
            maxProduct = power(3, N/3);
            break;

        // If division gives mod as 1, then break as
        // 4 + power of 3 for remaining part
        case 1:
            maxProduct = 2 * 2 * power(3, (N/3) - 1);
            break;

        // If division gives mod as 2, then break as
        // 2 + power of 3 for remaining part
        case 2:
            maxProduct = 2 * power(3, N/3);
            break;
    }
    return maxProduct;
}
*/

// 10. https://www.geeksforgeeks.org/numbers-of-length-n-having-digits-a-and-b-and-whose-sum-of-digits-contain-only-digits-a-and-b/
// N 位 digit，每位只可能为 A 或 B，求解所有的组合方案中 SUM(digit) 只包含 A 或 B 的方案数，总的方案数对 10^9+7 取模
// a. 首先枚举 A 可能出现的位数 x，则 B 对应出现的位数为 y (N-x)。判断该情况下 SUM(digit) 是否也只包含 A 或 B。如果符合条件，则对应的方案数加
// 					ANS += N!/(X!Y!) 可重集全排列公式
// b. 由于总的方案数需要对 10^9 + 7 取模，因此需要一种方法能够快速计算 N!/(X!Y!) % (10^9 +7)
// 费马小定理: a ^ (p-1) = 1 mod(p)，因此对于 N!/(X!Y!) % p --> N!*(X!)^(-1)*(Y!)^(-1) % p，其中 N! % p 和 (X!)^(-1) 均可以预处理出来
// (X!)^(-1) % p = X^(p-2) * (X-1)^(p-2) * ... * 2^(p-2) % p
/*
ll countNumbers(int n, int a, int b)
{
    ll fact[MAX], inv[MAX];
    ll ans = 0;

	// Generating factorials of all numbers
	// Tips: N! % p 预处理
    fact[0] = 1;
    for (int i = 1; i < MAX; i++) {
        fact[i] = (1LL * fact[i - 1] * i);
        fact[i] %= MOD;
    }

    // Generating inverse of factorials modulo
	// MOD of all numbers
	// 模逆元预处理
	// 费马小定理: a^(p-1) = 1 (mod p)
	// 1/(X!) % p = (X)^(p-2) * (X-1)^(p-2) * ... * 2^(p-2) % p = (X*(X-1)*...*2)^(p-2) % p
    inv[MAX - 1] = modInverse(fact[MAX - 1]);
    for (int i = MAX - 2; i >= 0; i--) {
        inv[i] = (inv[i + 1] * (i + 1));
        inv[i] %= MOD;
    }

    // Keeping a as largest number
    if (a < b)
        swap(a, b);

    // Iterate over all possible values of s and
    // if it is a valid S then proceed further
    for (int s = n; s <= 9 * n; s++) {
        if (!check(s, a, b))
            continue;

        // Check for invalid cases in the equation
        if (s < n * b || (s - n * b) % (a - b) != 0)
            continue;
        int numDig = (s - n * b) / (a - b);
        if (numDig > n)
            continue;

        // Find answer using combinatorics
        ll curr = fact[n];
        curr = (curr * inv[numDig]) % MOD;
        curr = (curr * inv[n - numDig]) % MOD;

        // Add this result to final answer
        ans = (ans + curr) % MOD;
    }
    return ans;
}
*/

func main() {
	// fmt.Println("Yes")
	// 离散对数测试
	fmt.Println(discreteLog(2, 3, 5))
	fmt.Println(discreteLog(3, 7, 11))
	fmt.Println(discreteLog(3, 7, 10))
}
