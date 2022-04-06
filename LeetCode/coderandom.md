# 代码随想录笔记

## Intro

这个文档是《代码随想录——跟着Carl学算法》的学习笔记。

[toc]

## 20220404 3.1-3.2

### 3.1

P35代码

```c++
#include <iostream>

using namespace std;

void test_arr() {
    int array[2][3] = {
            {0, 1, 2},
            {3, 4, 5}
    };
    cout << &array[0][0] << " " << &array[0][1] << " " << &array[0][2] << endl;
    cout << &array[1][0] << " " << &array[1][1] << " " << &array[1][2] << endl;
}

int main() {
    test_arr();
}
```

执行结果
> 0x7ff7b869d5d0 0x7ff7b869d5d4 0x7ff7b869d5d8
> 0x7ff7b869d5dc 0x7ff7b869d5e0 0x7ff7b869d5e4

可以看出每个int占据4字节，所以相邻元素是连续的

### 3.2 [704. Binary Search](https://leetcode-cn.com/problems/binary-search/)

二分查找虽然比较简单，但是也要仔细考虑边界条件。题目中是有序无重复的数组，因此可用二分法。

```c++
class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size() - 1;
        while (left <= right) {
            //这种写法是为了防止溢出，效果等同于(left + right) / 2
            int middle = left + ((right - left) / 2);
            if (nums[middle] > target) {
                right = middle - 1;
            } else if (nums[middle] < target) {
                left = middle + 1;
            } else {
                return middle;
            }
        }
        return -1;
    }
};
```

![image-20220404081853928](coderandom.assets/image-20220404081853928.png)

### 3.3 [27. Remove Element](https://leetcode-cn.com/problems/remove-element/)

这道题首先想到的就是暴力解法，两遍for循环，第一遍找到要删除的数字，嵌套的for负责把后面的元素踢到前面来。但是这个方法有点笨，因此使用双指针法：通过一个快指针和一个慢指针在一个for里完成两个for的工作。

```c++
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int slowIndex = 0;
        for (int fastIndex = 0; fastIndex < nums.size(); fastIndex++) {
            if (val != nums[fastIndex]) {
                nums[slowIndex++] = nums[fastIndex];
            }
        }
        return slowIndex;
    }
};
```

![image-20220405080947351](coderandom.assets/image-20220405080947351.png)

### 3.4 [209. Minimum Size Subarray Sum](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

暴力法使用两个for循环，反复寻找符合条件的子数组。也可以使用滑动窗口，也就是不断调整数组的起始位置和终止位置，来获得想要的结果。滑动窗口可以理解为是双指针法的一种。可以把复杂度从平方级下降到n级别：

```c++
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int result = INT32_MAX;
        int sum = 0;
        int i = 0;
        int subLength = 0;
        for (int j = 0; j < nums.size(); j++) {
            sum += nums[j];
            while (sum >= target) {
                subLength = (j - i + 1);
                result = result < subLength ? result : subLength;
                // 这里是这个滑动窗口的精髓，不断变更i（窗口的起始位置）
                sum -= nums[i++];
            }
        }
        return result == INT32_MAX ? 0 : result;
    }
};
```

![image-20220406072725685](coderandom.assets/image-20220406072725685.png)

### 3.5 [59. Spiral Matrix II](https://leetcode-cn.com/problems/spiral-matrix-ii/)

这道题需要注意的是，要么左闭右开，要么左开右闭。千万不能乱了。

```c++
class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        vector<vector<int>> res(n, vector<int>(n, 0));
        int startx = 0, starty = 0;
        int loop = n / 2;
        int mid = n / 2;
        int count = 1;
        int offset = 1;
        int i, j;
        while (loop--) {
            i = startx;
            j = starty;
            for (j = starty; j < starty + n - offset; j++) {
                res[startx][j] = count++;
            }
            for (i = startx; i < startx + n - offset; i++) {
                res[i][j] = count++;
            }
            for (; j >starty; j--) {
                res[i][j] = count++;
            }
            for(; i > startx; i--) {
                res[i][j] = count++;
            }
            startx++;
            starty++;
          	// offset用于控制每一圈中每一条边遍历的长度
            offset += 2;
        }
        if (n % 2) {
            res[mid][mid] = count;
        }
        return res;
    }
};
```



![image-20220407083529048](coderandom.assets/image-20220407083529048.png)
