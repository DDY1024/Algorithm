1.哥德巴赫猜想：任意一个偶数均可以被两个素数和表示。

猜想应用：
    https://www.geeksforgeeks.org/break-number-sum-maximum-divisors-parts-minimum/
上述问题分两种情况讨论：
    a. x 为偶数。如果 x = 2, 则结果为 1；如果 x > 2, 则 x 肯定能够被分解成两个素数之和，结果为 2
    b. x 为奇数。存在两种策略 x = 2 + (x-2) 和 x = 3 + (x-3)。如果 x - 2 恰好为素数，则结果为 2；对于 x - 3, 则其被分解成素数的个数至多为 2。
        因此，综上所述 x 为奇数时，其最终结果不会超过 3，有可能为 2.

伪代码如下所示：
/*
int minimumSum(int n) 
{ 
    if (isPrime(n)) 
        return 1; 
  
    // If n is an even number (we can 
    // write it as sum of two primes) 
    if (n % 2 == 0) 
        return 2; 
  
    // If n is odd and n-2 is prime. 
    if (isPrime(n - 2)) 
        return 2; 
  
    // If n is odd, n-3 must be even. 
    return 3; 
} 
*/

2. 反素数：给定整数范围 N, 求解最小的数 x <= N, 使得 x 约数个数最多
https://blog.csdn.net/acdreamers/article/details/25049767
https://www.geeksforgeeks.org/number-maximum-number-prime-factors/

反素数求解过程中的一些剪枝策略：
    a. (p1^x1)*(p2^x2)*...*(pk^xk)，其中满足 x1>=x2>=...>=xk
    b. 在约数个数求解的过程中，num 必须满足 n % num == 0
    c. num * (i + 1) > n, 则直接跳出循环
    d. 已经存在解 ans，如果当前状态下求出的解 >= ans, 则直接剪掉

可以参考如下代码：
/*
void dfs(int dept, int limit, LL tmp, int num)
{
    if(num > n) return;
    if(num == n && ans > tmp) ans = tmp;
    for(int i=1;i<=limit;i++)
    {
        if(ans / p[dept] < tmp || num*(i+1) > n) break;  // 剪枝策略
        tmp *= p[dept];
        if(n % (num*(i+1)) == 0) // 剪枝策略
            dfs(dept+1,i,tmp,num*(i+1));
    }
}
*/


素因子分解 --> 约数个数
线段树 --> 区间统计(动态)
RMQ --> 区间统计(静态)


3. 差分数组原理
    参考：https://www.cnblogs.com/COLIN-LIGHTNING/p/8436624.html

4. 二阶常系数齐次线性递推数列: https://zhuanlan.zhihu.com/p/33854447

