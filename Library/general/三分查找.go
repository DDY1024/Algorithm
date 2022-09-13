package main

// 应用场景：求解凸函数（单峰函数）极值点（该凸函数可以是连续，也可以是离散的）

// 参考资料：https://oi-wiki.org/basic/binary/

// 优化点：三分法每次会舍弃掉左边或右边的一个区间，为了减少三分法的次数，在选择划分点时，尽量选在区间中点附近

// 1. 连续函数（极小值）
/*
while (r - l > eps) {
  mid = (lmid + rmid) / 2;
  lmid = mid - eps;
  rmid = mid + eps;
  if (f(lmid) < f(rmid))
    r = mid;
  else
    l = mid;
}
*/

// 2. 连续函数（极大值）
/*
while (r - l > eps) {
  mid = (lmid + rmid) / 2;
  lmid = mid - eps;
  rmid = mid + eps;
  if (f(lmid) < f(rmid))
    l = mid;
  else
    r = mid;
}
*/

// 3. 离散函数（极大值）--> 极值点唯一
/*
for left < right {
  lmid = left + (right - left)/2;
  rmid = lmid + (right - lmid)/2;
  if f(lmid) < f(rmid)
      left = lmid
  else
      right = rmid
}
*/
