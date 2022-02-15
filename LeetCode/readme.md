# LeetCode Everyday

## 简介

这个文件夹用来存放每天做的Leetcode学到的内容。也许会把这个内容同步到博客上。

[toc]

## 2022年1月

### 20220124 [448. Find All Numbers Disappeared in an Array](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/)

```c++
class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        vector<int> ans;
        for (const int & num: nums) {
            int pos = abs(num) - 1;
            if (nums[pos] > 0) {
                nums[pos] = -nums[pos];
            }
        }
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] > 0) {
                ans.push_back(i + 1);
            }
        }
        return ans;
    }
};
```

![448. Find All Numbers Disappeared in an Array](README.assets/image-20220124085812012.png)

今天这道题，利用hash表，把所有重复出现的位置进行标记，再遍历一遍就可以找到没有出现过的数字。  
进阶题目要求原地hash，可以直接对原数组进行标记，把重复出现的数字标为负数。最后遍历，仍是正数的就是没有出现过的数字。

### 20220125 [48. Rotate Image](https://leetcode-cn.com/problems/rotate-image/)

```c++
class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int temp = 0, n = matrix.size() - 1;
        for (int i = 0; i <= n / 2; ++i) {
            for (int j = i; j < n - i; ++j) {
                temp = matrix[j][n-i];
                matrix[j][n-i] = matrix[i][j];
                matrix[i][j] = matrix[n-j][i];
                matrix[n-j][i] = matrix[n-i][n-j];
                matrix[n-i][n-j] = temp;
            }
        }
    }
};
```

![image-20220125073437226](README.assets/image-20220125073437226.png)

利用很巧妙的方法，可以一次循环内转移点对称的四个值，空间复杂度可以做到常数。

 [如下图题解所示](https://leetcode-cn.com/problems/rotate-image/solution/48-xuan-zhuan-tu-xiang-fu-zhu-ju-zhen-yu-jobi/)



![image-20220125073646549](README.assets/image-20220125073646549.png)

### 20220126 [240. Search a 2D Matrix II](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)

```c++
class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int m = matrix.size();
        if (m == 0) {
            return false;
        }
        int n = matrix[0].size();
        int i = 0, j = n - 1;
        while (i < m && j >= 0) {
            if (matrix[i][j] == target) {
                return true;
            } else if (matrix[i][j] > target) {
                --j;
            } else {
                ++i;
            }
        }
        return false;
    }
};
```

![image-20220126075519571](README.assets/image-20220126075519571.png)

今天这道题比较简单，因为每行每列都是增序的，因此只要判断，若当前值小于target那么就向下移动一位，如果大于则向左移动一位。如果直到右下角都没有，那么就不存在。



### 20220127 [769. Max Chunks To Make Sorted](https://leetcode-cn.com/problems/max-chunks-to-make-sorted/) 

```c++
class Solution {
public:
    int maxChunksToSorted(vector<int>& arr) {
        int chunks = 0, cur_max = 0;
        for (int i = 0; i < arr.size(); ++i) {
            cur_max = max (cur_max, arr[i]);
            if (cur_max == i) {
                ++chunks;
            }
        }
        return chunks;
    }
};
```

![image-20220127081813456](README.assets/image-20220127081813456.png)

一开始没读懂题目要干什么,原来是要拆分出一些逆序数字。如果当前最大值大于数组的标号，则说明右边一定有小于数组位置的数字。

### 20220128 [232. Implement Queue using Stacks](https://leetcode-cn.com/problems/implement-queue-using-stacks/)

```c++
class MyQueue {
    stack<int> in, out;
public:
    MyQueue() {}
    void inToOut() {
        if (out.empty()) {
            while (!in.empty()) {
                int x = in.top();
                in.pop();
                out.push(x);
            }
        }
    }
    
    void push(int x) {
        in.push(x);
    }
    
    int pop() {
        inToOut();
        int x = out.top();
        out.pop();
        return x;
    }
    
    int peek() {
        inToOut();
        return out.top();
    }
    
    bool empty() {
        return in.empty() && out.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
```

![image-20220128084747815](README.assets/image-20220128084747815.png)

尝试用栈实现一个队列（怎么会有这么怪的需求）

用了两个栈，一个表示输入一个表示输出，用一个inToOut函数把栈内的数据反转存放在out。

### 20220131 [155. Min Stack](https://leetcode-cn.com/problems/min-stack/)

```c++
class MinStack {
    stack<int> s, min_s;
public:
    MinStack() {}
    
    void push(int val) {
        s.push(val);
        if (min_s.empty() || min_s.top() >= val) {
            min_s.push(val);
        }
    }
    
    void pop() {
        if (!min_s.empty() && min_s.top() == s.top()) {
            min_s.pop();
        }
        s.pop();
    }
    
    int top() {
        return s.top();
    }
    
    int getMin() {
        return min_s.top();
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
```

![image-20220131094257627](README.assets/image-20220131094257627.png)

相当于额外建立了一个栈`min_s`，`min_s`的栈顶是原来栈里所有值里的最小值。

## 2022年2月

### 20220201 [20. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)

```c++
class Solution {
public:
    bool isValid(string s) {
        stack<char> parsed;
        for (int i = 0; i < s.length(); ++i) {
            if (s[i] == '{' || s[i] == '[' || s[i] == '(') {
                parsed.push(s[i]);
            } else {
                if (parsed.empty()) {
                    return false;
                }
                char c = parsed.top();
                if ((s[i] == '}' && c == '{') ||
                    (s[i] == ']' && c == '[') ||
                    (s[i] == ')' && c == '(')) {
                        parsed.pop();
                    } else {
                        return false;
                    }
            }
        }
        return parsed.empty();
    }
};
```

![image-20220201172102990](README.assets/image-20220201172102990.png)

一道典型的应用栈的题目，遇到左括号时候入栈，遇到右括号时候出栈。这个可能也就是形式语言与自动机的CFL/PDA吧。

### 20220202 [739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)

```c++
class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        int n = temperatures.size();
        vector<int> ans(n);
        stack<int> indices;
        for (int i = 0; i < n; ++i) {
            while (!indices.empty()) {
                int pre_index = indices.top();
                if (temperatures[i] <= temperatures[pre_index]) {
                    break;
                }
                indices.pop();
                ans[pre_index] = i - pre_index;
            }
            indices.push(i);
        }
        return ans;
    }
};
```

![image-20220202090214022](README.assets/image-20220202090214022.png)

这道题用单调栈解决。用单调递减的栈表示每天的温度。

### 20220203 [23. Merge k Sorted Lists](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

```c++
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    struct Comp {
        bool operator() (ListNode* l1, ListNode* l2) {
            return l1->val > l2->val;
        }
    };
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if (lists.empty()) return nullptr;
        priority_queue<ListNode*, vector<ListNode*>, Comp>  q;
        for (ListNode* list: lists) {
            if (list) {
                q.push(list);
            }
        }
        ListNode* dummy = new ListNode(0), *cur = dummy;
        while (!q.empty()) {
            cur -> next = q.top();
            q.pop();
            cur = cur -> next;
            if (cur->next) {
                q.push(cur->next);
            }
        }
        return dummy->next;
    }
};
```

![image-20220203073122301](README.assets/image-20220203073122301.png)

今天这道题乍看上去觉得不怎么难，但是想要速度快一点，就需要考虑用到优先队列。考虑用最小堆实现。

### 20220204 [218. The Skyline Problem](https://leetcode-cn.com/problems/the-skyline-problem/)

```c++
class Solution {
public:
    vector<vector<int>> getSkyline(vector<vector<int>>& buildings) {
        auto cmp = [](const pair<int, int>& a, const pair<int, int>& b) -> bool { return a.second < b.second; };
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> que(cmp);

        vector<int> boundaries;
        for (auto& building : buildings) {
            boundaries.emplace_back(building[0]);
            boundaries.emplace_back(building[1]);
        }
        sort(boundaries.begin(), boundaries.end());

        vector<vector<int>> ret;
        int n = buildings.size(), idx = 0;
        for (auto& boundary : boundaries) {
            while (idx < n && buildings[idx][0] <= boundary) {
                que.emplace(buildings[idx][1], buildings[idx][2]);
                idx++;
            }
            while (!que.empty() && que.top().first <= boundary) {
                que.pop();
            }

            int maxn = que.empty() ? 0 : que.top().second;
            if (ret.size() == 0 || maxn != ret.back()[1]) {
                ret.push_back({boundary, maxn});
            }
        }
        return ret;
    }
};
```

![image-20220204081728808](README.assets/image-20220204081728808.png)

自己按照优先队列扫描，存储每个建筑物的高度和右端，最后只过了35/40个测试用例。怎么调试都找不出问题，最后抄了官方题解。

### 20220207 [239. Sliding Window Maximum](https://leetcode-cn.com/problems/sliding-window-maximum/)

```c++
class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        deque<int> dq;
        vector<int> ans;
        for (int i = 0; i < nums.size(); ++i) {
            if (!dq.empty() && dq.front() == i - k) {
                dq.pop_front();
            }
            while (!dq.empty() && nums[dq.back()] < nums[i]) {
                dq.pop_back();
            }
            dq.push_back(i);
            if (i >= k - 1) {
                ans.push_back(nums[dq.front()]);
            }
        }
        return ans;
    }
};
```

![image-20220207070455924](README.assets/image-20220207070455924.png)

使用单调队列即可。push时要把前面比自己小的元素都删掉，直到遇到更大的元素才停止删除。难点在于如何写这样一个单调队列。

### 20220208 [1. Two Sum](https://leetcode-cn.com/problems/two-sum/)

```c++
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> hash;
        vector<int> ans;
        for (int i = 0; i < nums.size(); ++i) {
            int num = nums[i];
            auto pos = hash.find(target - num);
            if (pos == hash.end()) {
                hash[num] = i;
            } else {
                ans.push_back(pos->second);
                ans.push_back(i);
                break;
            }
        }
        return ans;
    }
};
```

![image-20220208074217827](README.assets/image-20220208074217827.png)

通过哈希表存储遍历过的值，每次遍历到`i`的时候，查看hashtable里有没有`target-nums[i]`

### 20220209 [128. Longest Consecutive Sequence](https://leetcode-cn.com/problems/longest-consecutive-sequence/)

```c++
class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        unordered_set<int> hash;
        for (const int & num: nums) {
            hash.insert(num);
        }
        int ans = 0;
        while (!hash.empty()) {
            int cur = *(hash.begin());
            hash.erase(cur);
            int next = cur + 1, prev = cur - 1;
            while (hash.count(next)) {
                hash.erase(next++);
            }
            while (hash.count(prev)) {
                hash.erase(prev--);
            }
            ans = max(ans, next - prev -1);
        }
        return ans;
    }
}
```

![image-20220209073312206](README.assets/image-20220209073312206.png)

今天的题目可以用hash表解决，在hash table不空的条件下，判断hash table的cur指针的前后项是否存在，若存在则计数。

### 20220210 [149. Max Points on a Line](https://leetcode-cn.com/problems/max-points-on-a-line/)

```c++
//leetcode official solution
class Solution {
public:
    int gcd(int a, int b) {
        return b ? gcd(b, a % b) : a;
    }

    int maxPoints(vector<vector<int>>& points) {
        int n = points.size();
        if (n <= 2) {
            return n;
        }
        int ret = 0;
        for (int i = 0; i < n; i++) {
            if (ret >= n - i || ret > n / 2) {
                break;
            }
            unordered_map<int, int> mp;
            for (int j = i + 1; j < n; j++) {
                int x = points[i][0] - points[j][0];
                int y = points[i][1] - points[j][1];
                if (x == 0) {
                    y = 1;
                } else if (y == 0) {
                    x = 1;
                } else {
                    if (y < 0) {
                        x = -x;
                        y = -y;
                    }
                    int gcdXY = gcd(abs(x), abs(y));
                    x /= gcdXY, y /= gcdXY;
                }
                mp[y + x * 20001]++;
            }
            int maxn = 0;
            for (auto& [_, num] : mp) {
                maxn = max(maxn, num + 1);
            }
            ret = max(ret, maxn);
        }
        return ret;
    }
};
```

![image-20220210085357025](README.assets/image-20220210085357025.png)

今天这道题和题解的思路一模一样，但是就是通不过，找了半天也没找到原因，粘贴了官方题解的代码顺利通过了。怪事情。

### 20220211 [332. Reconstruct Itinerary](https://leetcode-cn.com/problems/reconstruct-itinerary/)

```c++
class Solution {
private:
    unordered_map<string, map<string, int>> targets;
    bool backtracking(int ticketNum, vector<string>& result) {
        if (result.size() == ticketNum + 1) {
            return true;
        }
        for (pair<const string, int>& target : targets[result[result.size() - 1]]) {
            if (target.second > 0) {
                result.push_back(target.first);
                target.second--;
                if (backtracking(ticketNum, result)) return true;
                result.pop_back();
                target.second++;
            }
        }
        return false;
    }

public:
    vector<string> findItinerary(vector<vector<string>>& tickets) {
        targets.clear();
        vector<string> result;
        for (const vector<string>& vec : tickets) {
            targets[vec[0]][vec[1]]++;
        }
        result.push_back("JFK");
        backtracking(tickets.size(), result);
        return result;
    }
};
```

![image-20220211080416055](README.assets/image-20220211080416055.png)

深搜使用回溯的思想解决。难点在于如何处理死循环。

### 20220214 [303. Range Sum Query - Immutable](https://leetcode-cn.com/problems/range-sum-query-immutable/)

```c++
class NumArray {
    vector<int> psum;
public:
    NumArray(vector<int>& nums): psum(nums.size() + 1, 0){
        partial_sum(nums.begin(), nums.end(), psum.begin() + 1);
    }
    
    int sumRange(int left, int right) {
        return psum[right+1] - psum[left];
    }
};

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray* obj = new NumArray(nums);
 * int param_1 = obj->sumRange(left,right);
 */
```

![image-20220214091945500](README.assets/image-20220214091945500.png)

建立数组`psum`存储`nums`每个位置之前所有数字的和。

### 20220215 [304. Range Sum Query 2D - Immutable](https://leetcode-cn.com/problems/range-sum-query-2d-immutable/)

```c++
class NumMatrix {
    vector<vector<int>> integral;
public:
    NumMatrix(vector<vector<int>>& matrix) {
        int m = matrix.size(), n = m > 0? matrix[0].size(): 0;
        integral = vector<vector<int>>(m + 1, vector<int>(n + 1,0));
        for (int i = 1; i <= m; ++i) {
            for (int j = 1; j <= n; ++j) {
                integral[i][j] = matrix[i-1][j-1] + integral[i-1][j] +  integral[i][j-1]- integral[i-1][j-1];
            }
        }
    }
    
    int sumRegion(int row1, int col1, int row2, int col2) {
        return integral[row2+1][col2+1] - integral[row2+1][col1] - integral[row1][col2+1] + integral[row1][col1]; 
    }
};

/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix* obj = new NumMatrix(matrix);
 * int param_1 = obj->sumRegion(row1,col1,row2,col2);
 */
```



![image-20220215083135174](README.assets/image-20220215083135174.png)

这道题用到了积分图，也就是昨天那道题使用的前缀和的二维版本。可以维护一个二维 `preSum` 数组，专门记录以原点为顶点的矩阵的元素之和，就可以用几次加减运算算出任何一个子矩阵的元素和。
