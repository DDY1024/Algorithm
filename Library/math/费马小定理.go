package main

// 参考：https://zh.wikipedia.org/wiki/%E8%B4%B9%E9%A9%AC%E5%B0%8F%E5%AE%9A%E7%90%86
// 假设 p 是一个素数，则 a^(p-1) = 1 (mod) p

// 费马小定理应用
// 1.除法取模操作(直接求解模逆元)
// (a/b)%p = a*(b^(p-2))%p = (a/b)*(b^(p-1))%p
