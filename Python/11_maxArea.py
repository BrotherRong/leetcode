# -*- coding: utf-8 -*-
# @Time     : 2020/9/4 16:46
# @Author   : Rong
# @email    : mrrong0720@163.com
from typing import List
# 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
# 说明：你不能倾斜容器，且 n 的值至少为 2。
# 示例：
# 输入：[1,8,6,2,5,4,8,3,7]
# 输出：49

class Solution:
    def maxArea(self, height: List[int]) -> int:
        i , j , res = 0 , len(height)-1 , 0
        while i < j :
            if height[i] < height[j]  :
                res = max(res , height[i]*(j - i))
                i += 1
            else :
                res = max(res , height[j]*(j - i))
                j -= 1
        return res


s = Solution()
print(s.maxArea([1,8,6,2,5,4,8,3,7]))
