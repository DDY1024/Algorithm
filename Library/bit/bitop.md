## 进制转换
- 十进制整数 --> 二进制整数
    - 不断除 2 取余，逆序输出 
- 十进制小数 --> 二进制小数
    - 不断乘 2 取整，正序输出

## 位运算
### go 运算符优先级
- https://golang.org/ref/spec#Operators

### 常见公式
- a|b = (a^b) + (a&b)
- a^b = (a|b) - (a&b)
- a+b = (a|b) + (a&b) = (a&b)*2 + (a^b) = (a|b)*2 - (a^b)
- (a&b)^(a&c) = a&(b^c)

### 常见操作
#### 32 位 or 64 位
```go
const bitSize = 32 << (^uint(0) >> 32 & 1)  
```

#### 求绝对值
```cpp
int Abs(int n) {
  return (n ^ (n >> 31)) - (n >> 31);
  /* n>>31 取得 n 的符号，若 n 为正数，n>>31 等于 0，若 n 为负数，n>>31 等于 -1
    若 n 为正数 n^0=n, 数不变，若 n 为负数有 n^(-1)
    需要计算 n 和 -1 的补码，然后进行异或运算，
    结果 n 变号并且为 n 的绝对值减 1，再减去 -1 就是绝对值 */
}
```
#### 最大值/最小值
```cpp
// 如果 a >= b, (a - b) >> 31 为 0，否则为 -1
int max(int a, int b) { return (b & ((a - b) >> 31)) | (a & (~(a - b) >> 31)); }
int min(int a, int b) { return (a & ((a - b) >> 31)) | (b & (~(a - b) >> 31)); }
```
#### 两个数符号是否相同
```go
func isSameSign(x, y int) bool {
    return (x^y) >= 0 // 符号相同：最高位异或结果为 0，>= 0；符号不同：最高位异或结果为 1，< 0
}
```

#### 交换两个数
```cpp
void swap(int &a, int &b) { a ^= b ^= a ^= b; }
```

#### 获取二进制某一位
```cpp
int getBit(int a, int b) { return (a >> b) & 1; }  
```

#### 清除某一位
- go: `a&^(1<<b)`
- cpp: `a & ~(1<<b)`

#### 设置某一位
- a | (1<<b)

#### 某一位取反
- a ^ (1<<b)

#### 统计 1 的个数
```cpp
// 求 x 的汉明权重
int popcount(int x) {
    int cnt = 0;
    while (x) {
        cnt += x & 1;
        x >>= 1;
    }
    return cnt;
}

// 求 x 的汉明权重
int popcount(int x) {
    int cnt = 0;
    while (x) {
        cnt++;
        x -= x & -x;
    }
    return cnt;
}

func countBit(x int) int {
	cnt := 0
	for x > 0 {
		x &= (x - 1)
		cnt++
	}
	return cnt
}
```

```go
// 快速统计 32 位整数中 1 的个数
// O(log log n)
func countBitTwo(n uint32) uint32 {
	// 01,01,01,...,01
	// 0011,...,0011
	// 00001111,...,00001111
	// 0000000011111111,...,0000000011111111
	// 00000000000000001111111111111111
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	return n
}
```

#### 汉明权重递增的排列
```cpp
// 1. 最右边可以向左移动一位的 1 左移一下
// 2. 其右边的 1 全部移动到最右面
int t = x + (x & -x);
x = t | ((((t&-t)/(x&-x))>>1)-1);
```
```go
// 枚举所有汉明权重排列
for (int i = 0; (1<<i)-1 <= n; i++) { // 按照数字 1 个数从小到大作为起点进行枚举
    for (int x = (1<<i)-1, t; x <= n; t = x+(x&-x), x = x ? (t|((((t&-t)/(x&-x))>>1)-1)) : (n+1)) {
        // 写下需要完成的操作
    }
}
```

#### 子集枚举
```go
    var x int 
    for i := x; i > 0; i = (i-1)&x

    // 注意: 需要额外处理单独的 0（空集）
```

## 参考资料
- https://oi-wiki.org/math/base/
- https://oi-wiki.org/math/bit/
- https://oi-wiki.org/math/binary-set/