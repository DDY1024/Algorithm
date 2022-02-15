# 题目链接：https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/

class Solution:
    def luckyNumbers(self, matrix: List[List[int]]) -> List[int]:
        minRow = [min(row) for row in matrix]
        maxCol = [max(col) for col in zip(*matrix)]  # 通过 zip 方法将 行元组 --> 列元组
        ans = []
        for i, row in enumerate(matrix):  #
            for j, x in enumerate(row): #
                if x == minRow[i] == maxCol[j]:
                    ans.append(x)
        return ans

# zip：https://www.runoob.com/python/python-func-zip.html
# enumerate: https://www.runoob.com/python/python-func-enumerate.html --> (下标索引，数据项)