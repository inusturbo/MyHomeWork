# 代码随想录笔记

## Intro

这个文档是《代码随想录——跟着Carl学算法》的学习笔记。

[toc]

## 20220404 3.1-3.2

### 3.1 20220404

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

### 3.2 [704. Binary Search](https://leetcode-cn.com/problems/binary-search/) 20220404

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

### 3.3 [27. Remove Element](https://leetcode-cn.com/problems/remove-element/) 20220405

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

### 3.4 [209. Minimum Size Subarray Sum](https://leetcode-cn.com/problems/minimum-size-subarray-sum/) 20220406

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

### 3.5 [59. Spiral Matrix II](https://leetcode-cn.com/problems/spiral-matrix-ii/) 20220407

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

### 4.1-4.2 [203. Remove Linked List Elements](https://leetcode-cn.com/problems/remove-linked-list-elements/) 20220408

今天是一道基础的删除链表中的节点的题。通过设置一个虚拟的头节点，讲所有删除的操作都统一起来。

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
    ListNode* removeElements(ListNode* head, int val) {
        ListNode* dummyHead = new ListNode(0);
        dummyHead->next = head;
        ListNode* cur = dummyHead;
        while (cur->next != NULL) {
            if (cur->next->val == val) {
                ListNode* tmp = cur->next;
                cur->next = cur->next->next;
                delete tmp;
            } else {
                cur = cur->next;
            }
        }
        head = dummyHead->next;
        delete dummyHead;
        return head;

    }
};
```

![image-20220408073349804](coderandom.assets/image-20220408073349804.png)

### 20220411 4.3 [707. Design Linked List](https://leetcode-cn.com/problems/design-linked-list/)

```c++
class MyLinkedList {
public:
    struct LinkedNode {
        int val;
        LinkedNode* next;
        LinkedNode(int val) : val(val), next(nullptr) { }
    };

    MyLinkedList() {
        _dummyHead = new LinkedNode(0);
        _size = 0;
    }
    
    int get(int index) {
        if (index > (_size - 1) || index < 0) {
            return -1;
        }
        LinkedNode* cur = _dummyHead->next;
        while(index--) {
            cur = cur->next;
        }
        return cur->val;
    }
    
    void addAtHead(int val) {
        LinkedNode* newNode = new LinkedNode(val);
        newNode->next = _dummyHead->next;
        _dummyHead->next = newNode;
        _size++;
    }
    
    void addAtTail(int val) {
        LinkedNode* newNode = new LinkedNode(val);
        LinkedNode* cur = _dummyHead;
        while (cur->next != nullptr) {
            cur = cur->next;
        }
        cur->next = newNode;
        _size++;
    }
    
    void addAtIndex(int index, int val) {
        if (index > _size) {
            return;
        }
        LinkedNode* newNode = new LinkedNode(val);
        LinkedNode* cur = _dummyHead;
        while (index--) {
            cur = cur->next;
        }
        newNode->next = cur->next;
        cur->next = newNode;
        _size++;
    }
    
    void deleteAtIndex(int index) {
        if (index >= _size || index < 0) {
            return;
        }
        LinkedNode* cur = _dummyHead;
        while (index--) {
            cur = cur->next;
        }
        LinkedNode* tmp = cur->next;
        cur->next = cur->next->next;
        delete tmp;
        _size--;
    }
 private:
    int _size;
    LinkedNode* _dummyHead;
};

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList* obj = new MyLinkedList();
 * int param_1 = obj->get(index);
 * obj->addAtHead(val);
 * obj->addAtTail(val);
 * obj->addAtIndex(index,val);
 * obj->deleteAtIndex(index);
 */
```

![image-20220411073226288](coderandom.assets/image-20220411073226288.png)

今天这道题相当于自己完整写一个链表和处理逻辑，整体难度不高，但是有一些细节需要处理，如对于非法值的处理等等。

### 20220412 4.4 [206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)

```c++
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode* tmp;
        ListNode* cur = head;
        ListNode* pre = nullptr;
        while (cur) {
            tmp = cur->next;
            cur->next = pre;
            pre = cur;
            cur = tmp;
        }
        return pre;

    }
};
```

![image-20220412071749755](coderandom.assets/image-20220412071749755.png)

今天是一道反转链表的题目，只要注意处理好循环中值得条件即可，难度不大。1⃣️让tmp=cur的下一个节点2⃣️让cur的下一个节点变成pre3⃣️将pre移动至cur4⃣️将cur移动至tmp

### 20220413 4.5 [19. Remove Nth Node From End of List](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

```c++
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* dummyHead = new ListNode(0);
        dummyHead->next = head;
        ListNode* slow = dummyHead;
        ListNode* fast = dummyHead;
        while(n-- && fast != nullptr) {
            fast = fast->next;
        }
        fast = fast->next;
        while (fast != nullptr){
            fast = fast->next;
            slow = slow->next;
        }
        slow->next = slow->next->next;
        return dummyHead->next;
    }
};
```

![image-20220413073208088](coderandom.assets/image-20220413073208088.png)

今天这道题做法很神奇，使用双指针，先让fast移动N步，随后让fast和slow同时移动，直到fast到达链表末尾。

### 20220414 4.6 [142. Linked List Cycle II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

```c++
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        ListNode* fast = head;
        ListNode* slow = head;
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
            if (slow == fast) {
                ListNode* index1 = fast;
                ListNode* index2 = head;
                while (index1 != index2) {
                    index1 = index1->next;
                    index2 = index2->next;
                }
                return index2;
            }
        }
        return nullptr;
    }
};
```

![image-20220414074340660](coderandom.assets/image-20220414074340660.png)

这道题还需要一些数学知识。首先使用快慢指针判断是否存在环。然后通过简单的数学计算相遇位置和head位置的相对关系。 详情见书P66。

### 20220415 5.1 5.2 [242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)

```c++
class Solution {
public:
    bool isAnagram(string s, string t) {
        int record[26] = {0};
        for (int i = 0; i < s.size(); i++) {
            record[s[i] - 'a']++;
        }
        for (int i = 0; i < t.size(); i++) {
            record[t[i] - 'a']--;
        }
        for (int i = 0; i < 26; i++) {
            if (record[i] != 0) {
                return false;
            }
        }
        return true;
    }
};
```

![image-20220415073712585](coderandom.assets/image-20220415073712585.png)

今天这道题蛮简单的，就是拉一个hash table看看有没有重复和冲突。没有就返回true

### 20220418 5.3 [349. Intersection of Two Arrays](https://leetcode-cn.com/problems/intersection-of-two-arrays/)

```c++
class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        unordered_set<int> result_set;
        unordered_set<int> nums_set(nums1.begin(), nums1.end());
        for (int num: nums2) {
            if (nums_set.find(num)!= nums_set.end()){
                result_set.insert(num);
            }
        }
        return vector<int>(result_set.begin(),result_set.end());
    }
};
```

![image-20220418074016075](coderandom.assets/image-20220418074016075.png)

今天这道题利用hash table就可以很方便的做出来。

### 20220419 5.4 [1. Two Sum](https://leetcode-cn.com/problems/two-sum/)

```c++
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> map;
        for (int i = 0; i < nums.size(); i++) {
            auto iter = map.find(target - nums[i]);
            if (iter != map.end()) {
                return {iter->second, i};
            }
            map.insert(pair<int, int> (nums[i],i));
        }
        return {};
    }
};
```

![image-20220419074133214](coderandom.assets/image-20220419074133214.png)

本题用map而不用set 的原因：set是一个集合，里面放的元素只能是一个key，但是本题不光要判断y是否存在，还要记住y的位置，map是<key, value>的结构，可以保存值的同时也保留下标。

### 20220420 5.5 [454. 4Sum II](https://leetcode-cn.com/problems/4sum-ii/)

```c++
class Solution {
public:
    int fourSumCount(vector<int>& nums1, vector<int>& nums2, vector<int>& nums3, vector<int>& nums4) {
        unordered_map<int, int> umap;
        for (int a : nums1) {
            for (int b : nums2) {
                umap[a + b]++;
            }
        }
        int count = 0;
        for (int c : nums3) {
            for (int d : nums4) {
                if (umap.find(0 - (c + d)) != umap.end()) {
                    count += umap[0 - (c + d)];
                }
            }
        }
        return count;
    }
};
```

![image-20220420072452105](coderandom.assets/image-20220420072452105.png)

把nums1、nums2算成第一组，nums3、nums4算成第二组，这样就把这个题转换成第一1题的解决方法。

### 20220421 5.6 [15. 3Sum](https://leetcode-cn.com/problems/3sum/)

```c++
class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
        sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] > 0) {
                return result;
            }
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            int left = i + 1;
            int right = nums.size() - 1;
            while (right > left) {
                if (nums[i] + nums[left] + nums[right] > 0) {
                    right--;
                } else if (nums[i] + nums[left] + nums[right] < 0) {
                    left++;
                } else {
                    result.push_back(vector<int>{nums[i], nums[left], nums[right]});
                    while (right >left && nums[right] == nums[right - 1]) {
                        right--;
                    }
                    while (right >left && nums[left] == nums[left + 1]) {
                        left++;
                    }
                    right--;
                    left++;
                }
            }
        }
        return result;
    }
};
```

![image-20220421072045390](coderandom.assets/image-20220421072045390.png)

这道题和偶数个数字求和有所不同，如果仍然用hash法就会非常复杂。因此本题使用了双指针法。首先将整个数组排序。如果nums[i]+nums[left]+nums[right]大了，那就让right往左移动。如果小了，就让left往右移动，知道left和right相遇。将i遍历一遍之后，就可以获得所有规定的数字了。

### 20220422 5.7 [18. 4Sum](https://leetcode-cn.com/problems/4sum/)

```c++
class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        vector<vector<int>> result;
        sort(nums.begin(), nums.end());
        for (int k = 0; k < nums.size(); k++) {
            if (k > 0 && nums[k] == nums[k - 1]) {
                continue;
            }
            for (int i = k + 1; i < nums.size(); i++) {
                if (i > k + 1 && nums[i] == nums[i - 1]) {
                    continue;
                }
                int left = i + 1;
                int right = nums.size() - 1;
                while (right > left) {
                    if (nums[k] + nums[i] > target - (nums[left] + nums[right])) {
                        right--;
                    } else if (nums[k] + nums[i] < target - (nums[left] + nums[right])) {
                        left++;
                    } else {
                        result.push_back(vector<int>{nums[k], nums[i], nums[left], nums[right]});
                        while (right > left && nums[right] == nums[right - 1]) right--;
                        while (right > left && nums[left] == nums[left + 1]) left++;
                        right--;
                        left++;
                    }
                }
            }
        }
        return result;
    }
};
```

![image-20220422073358768](coderandom.assets/image-20220422073358768.png)

这道题和昨天那道题有点像，都可以用双指针解决，不同的是，这次套了两次循环。

### 20220425 6.1 6.2 [344. Reverse String](https://leetcode-cn.com/problems/reverse-string/)

```c++
class Solution {
public:
    void reverseString(vector<char>& s) {
        for (int i = 0, j = s.size() - 1; i < s.size()/2; i++, j--) {
            swap(s[i], s[j]);
        }
    }
};
```

![image-20220425065306061](coderandom.assets/image-20220425065306061.png)

类似前几天的链表反转，但是字符串原则是数组，直接按下表形成“双指针”，来反转就可以了。

### 20220426 6.3 [541. Reverse String II](https://leetcode-cn.com/problems/reverse-string-ii/)

```c++
class Solution {
public:
    string reverseStr(string s, int k) {
        for (int i = 0; i < s.size(); i += (2 * k)) {
            if ( i + k <= s.size()) {
                reverse(s.begin() + i, s.begin() + i + k);
                continue;
            }
            reverse(s.begin() + i, s.begin() + s.size());
        }
        return s;
    }
};
```

![image-20220426075515314](coderandom.assets/image-20220426075515314.png)

这道题就是把昨天的字符串反转在局部实现了。在for上进行控制就可以。

### 20220427 6.4 [151. Reverse Words in a String](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

```c++
class Solution {
public:
    void reverse(string& s, int start, int end) {
        for (int i = start, j = end; i < j; i++, j--) {
            swap(s[i], s[j]);
        }
    }
    void removeExtraSpaces(string& s) {
        int slowIndex = 0, fastIndex = 0;
        while (s.size() > 0 && fastIndex < s.size() && s[fastIndex] == ' ') {
            fastIndex++;
        }
        for (; fastIndex < s.size(); fastIndex++) {
            if (fastIndex - 1 > 0 && s[fastIndex - 1] == s[fastIndex] && s[fastIndex] == ' ') {
                continue;
            } else {
                s[slowIndex++] = s[fastIndex];
            }
        }
        if (slowIndex - 1 > 0 && s[slowIndex - 1] == ' ') {
            s.resize(slowIndex - 1);
        } else {
            s.resize(slowIndex);
        }
    }
    string reverseWords(string s) {
        removeExtraSpaces(s);
        reverse(s, 0, s.size() - 1);
        int start = 0;
        int end = 0;
        bool entry = false;
        for (int i = 0; i < s.size(); i++) {
            if ((!entry) || (s[i] != ' ' && s[i - 1] == ' ')) {
                start = i;
                entry = true;
            }
            if (entry && s[i] == ' ' && s[i - 1] != ' ') {
                end = i - 1;
                entry = false;
                reverse(s, start, end);
            }
            if (entry && (i == (s.size() - 1)) && s[i] != ' ') {
                end = i;
                entry = false;
                reverse(s, start, end);
            }
        }
        return s;

    }
};
```

![image-20220427074604502](coderandom.assets/image-20220427074604502.png)

思路：1.删除多余的空格。2.字符串反转。3.单词反转

### 20220428 6.5 KMP算法理论

前缀表：

前缀：包含首字母，不包含尾字母的所有子串

后缀：只包含尾字母，不包含首字母的所有子串

求：最长相等前后缀

例如：aabaaf的最长相等前后缀

| 字符串 | 最长相等前后缀长度 |
| ------ | ------------------ |
| a      | 0                  |
| aa     | 1                  |
| aab    | 0                  |
| aaba   | 1                  |
| aabaa  | 2                  |
| aabaaf | 0                  |

模式串

| 0    | 1    | 2    | 3    | 4    | 5    |
| ---- | ---- | ---- | ---- | ---- | ---- |
| a    | a    | b    | a    | a    | f    |
| 0    | 1    | 0    | 1    | 2    | 0    |

### 20220429 6.6 [28. Implement strStr()](https://leetcode-cn.com/problems/implement-strstr/)

```c++
class Solution {
public:
    void getNext(int* next, const string& s) {
        int j = 0;
        next[0] = 0;
        for(int i = 1; i < s.size(); i++) {
            while (j > 0 && s[i] != s[j]) {
                j = next[j - 1];
            }
            if (s[i] == s[j]) {
                j++;
            }
            next[i] = j;
        }
    }
    int strStr(string haystack, string needle) {
        if (needle.size() == 0) {
            return 0;
        }
        int next[needle.size()];
        getNext(next, needle);
        int j = 0;
        for (int i = 0; i < haystack.size(); i++) {
            while (j > 0 && haystack[i] != needle[j]) {
                j = next[j - 1];
            }
            if (haystack[i] == needle[j]) {
                j++;
            }
            if (j == needle.size()) {
                return (i - needle.size() + 1);
            }
        }
        return -1;
    }
};
```

![image-20220429092014817](coderandom.assets/image-20220429092014817.png)

今天这道题实践了昨天的KMP算法。

### 20220502 6.7 [459. Repeated Substring Pattern](https://leetcode-cn.com/problems/repeated-substring-pattern/)

```c++
class Solution {
public:
    void getNext(int* next, const string& s) {
        next[0] = -1;
        int j = -1;
        for (int i = 1; i < s.size(); i++) {
            while (j > 0 && s[i] != s[j + 1]) {
                j = next[j];
            }
            if (s[i] == s[j + 1]) {
                j++;
            }
            next[i] = j;
        }
    }
    bool repeatedSubstringPattern(string s) {
        if (s.size() == 0) {
            return false;
        }
        int next[s.size()];
        getNext(next, s);
        int len = s.size();
        if (next[len - 1] != -1 && len % (len - (next[len - 1] + 1)) == 0) {
            return true;
        }
        return false;
    }
};
```

![image-20220502070459167](coderandom.assets/image-20220502070459167.png)

这道题乍看上去想不到和KMP算法的联系，但是KMP算法的本质不就是用来寻找重复的子串吗。因此用KMP算法可以方便的解决本题。

### 20220503 7.1 7.2 [232. Implement Queue using Stacks](https://leetcode-cn.com/problems/implement-queue-using-stacks/)

```c++
class MyQueue {
public:
    stack<int> stIn;
    stack<int> stOut;

    MyQueue() {

    }
    
    void push(int x) {
        stIn.push(x);
    }
    
    int pop() {
        if (stOut.empty()) {
            while (!stIn.empty()) {
                stOut.push(stIn.top());
                stIn.pop();
            }
        }
        int result = stOut.top();
        stOut.pop();
        return result;
    }
    
    int peek() {
        int res = this->pop();
        stOut.push(res);
        return res;
    }
    
    bool empty() {
        return stIn.empty() && stOut.empty();
    }
};
```

![image-20220503092139178](coderandom.assets/image-20220503092139178.png)

在工业级应用中，一定要注意代码的复用，尽量避免重复造轮子。

### 20220504 7.3 [225. Implement Stack using Queues](https://leetcode-cn.com/problems/implement-stack-using-queues/)

```c++
class MyStack {
public:
    queue<int> que1;
    queue<int> que2;
    MyStack() {

    }
    
    void push(int x) {
        que1.push(x);
    }
    
    int pop() {
        int size = que1.size();
        size--;
        while (size--) {
            que2.push(que1.front());
            que1.pop();
        }
        int result = que1.front();
        que1.pop();
        que1 = que2;
        while (!que2.empty()) {
            que2.pop();
        }
        return result;
    }
    
    int top() {
        return que1.back();
    }
    
    bool empty() {
        return que1.empty();
    }
};
```

![image-20220504071539717](coderandom.assets/image-20220504071539717.png)

利用队列实现栈

### 20220505 7.4 [20. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)

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

![image-20220505092059560](coderandom.assets/image-20220505092059560.png)

### 20220506 7.5 [150. Evaluate Reverse Polish Notation](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)

```c++
class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<int> st;
        for (int i = 0; i < tokens.size(); i++) {
            if (tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/") {
                int num1 = st.top();
                st.pop();
                int num2 = st.top();
                st.pop();
                if (tokens[i] == "+") st.push(num2 + num1);
                if (tokens[i] == "-") st.push(num2 - num1);
                if (tokens[i] == "*") st.push(num2 * num1);
                if (tokens[i] == "/") st.push(num2 / num1);
            } else {
                st.push(stoi(tokens[i]));
            }
        }
        int result = st.top();
        st.pop();
        return result;

    }
};
```

![image-20220506072044907](coderandom.assets/image-20220506072044907.png)

### 20220509 7.6 [239. Sliding Window Maximum](https://leetcode.cn/problems/sliding-window-maximum/)

```c++
class Solution {
private:
    class MyQueue {
    public:
        deque<int> que;
        void pop(int value) {
            if (!que.empty() && value == que.front()) {
                que.pop_front();
            }
        }
        void push(int value) {
            while (!que.empty() && value > que.back()) {
                que.pop_back();
            }
            que.push_back(value);
        }
        int front() {
            return que.front();
        }
    };
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        MyQueue que;
        vector<int> result;
        for (int i = 0; i < k; i++) {
            que.push(nums[i]);
        }
        result.push_back(que.front());
        for (int i = k; i <nums.size(); i++) {
            que.pop(nums[i - k]);
            que.push(nums[i]);
            result.push_back(que.front());
        }
        return result;
    }
};
```

![image-20220509072309125](coderandom.assets/image-20220509072309125.png)

### 20220510 7.7 [347. Top K Frequent Elements](https://leetcode.cn/problems/top-k-frequent-elements/)

```c++
class Solution {
public:
    class mycomparison {
    public:
        bool operator() (const pair<int, int>& lhs, const pair<int, int>& rhs) {
            return lhs.second > rhs.second;
        }
    };
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> map;
        for (int i = 0; i < nums.size(); i++) {
            map[nums[i]]++;
        }
        priority_queue<pair<int, int>, vector<pair<int, int>>, mycomparison> pri_que;
        for (unordered_map<int, int>::iterator it = map.begin(); it != map.end(); it++) {
            pri_que.push(*it);
            if (pri_que.size() > k) {
                pri_que.pop();
            }
        }
        vector<int> result(k);
        for (int i = k - 1; i >= 0; i--) {
            result[i] = pri_que.top().first;
            pri_que.pop();
        }
        return result;
    }
};
```

![image-20220510073315314](coderandom.assets/image-20220510073315314.png)

在C++中优先级队列`priority_queue`和`max_heap`的底层实现都是大顶堆。

一开始想到要用大顶堆，但是定义一个大小为k的大顶堆，每次更新的时候都把最大值弹出了，无法保留前k个高频元素。因此考虑使用小顶堆。

### 20220511 [42. Trapping Rain Water](https://leetcode.cn/problems/trapping-rain-water/)

```c++
class Solution {
public:
    int trap(vector<int>& height) {
        if (height.size() <= 2) return 0;
        vector<int> maxLeft(height.size(), 0);
        vector<int> maxRight(height.size(), 0);
        int size = maxRight.size();
        maxLeft[0] = height[0];
        for (int i = 1; i < size; i ++) {
            maxLeft[i] = max(height[i], maxLeft[i - 1]);
        }
        maxRight[size - 1] = height[size - 1];
        for (int i = size - 2; i >= 0; i--) {
            maxRight[i] = max(height[i], maxRight[i + 1]);
        }
        int sum = 0;
        for (int i = 0; i < size; i++) {
            int count = min(maxLeft[i], maxRight[i]) - height[i];
            if (count > 0) sum += count;
        }
        return sum;
    }
};
```

![image-20220511073213899](coderandom.assets/image-20220511073213899.png)

### 20220512 8.1

二叉树：

如果用顺序存储二叉树，如果父节点的下标是i，

那么他的左孩子是$i\times 2 + 1$,右孩子是$i\times 2+2$

遍历：经常会用递归的方式来实现深度优先遍历，也就是实现前序、中序、后序遍历。广度优先遍历主要使用队列的结构实现。

### 20220513 8.2 前中后序的递归遍历

1. 确定递归函数的参数和返回值

   `void traversal(TreeNode* cur, vector<int>& vec)`

   前一个参数是树的当前节点，后一个参数是存放节点数据的数组。

2. 确定终止条件

   当遍历节点为空时，则直接返回。

   `if (cur == NULL) return;`

3. 确定单层递归的逻辑

   ```c++
   //前序
   vec.push_back(cur->val);		//中
   traversal(cur->left, vec);	//左
   traversal(cur->right, vec);	//右
   //中序
   traversal(cur->left, vec);	//左
   vec.push_back(cur->val);		//中
   traversal(cur->right, vec);	//右
   //后序
   traversal(cur->left, vec);	//左
   traversal(cur->right, vec);	//右
   vec.push_back(cur->val);		//中
   ```

   

### 20220516 8.3 前中后序的迭代遍历

前序 [144. Binary Tree Preorder Traversal](https://leetcode.cn/problems/binary-tree-preorder-traversal/)

```c++
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        stack<TreeNode*> st;
        vector<int> result;
        if (root == NULL) return result;
        st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            st.pop();
            result.push_back(node->val);
            if (node->right) st.push(node->right);
            if (node->left) st.push(node->left);
        }
        return result;

    }
};
```

![image-20220516065549651](coderandom.assets/image-20220516065549651.png)

中序 [94. Binary Tree Inorder Traversal](https://leetcode.cn/problems/binary-tree-inorder-traversal/)

```c++
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        TreeNode* cur = root;
        while (cur != NULL || !st.empty()) {
            if (cur != NULL) {
                st.push(cur);
                cur = cur->left;
            } else {
                cur = st.top();
                st.pop();
                result.push_back(cur->val);
                cur = cur->right;
            }
        }
        return result;
    }
};
```

![image-20220516070934760](coderandom.assets/image-20220516070934760.png)

后序 [145. Binary Tree Postorder Traversal](https://leetcode.cn/problems/binary-tree-postorder-traversal/)

```c++
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        stack<TreeNode*> st;
        vector<int> result;
        if (root == NULL) return result;
        st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            st.pop();
            result.push_back(node->val);
            if (node->left) st.push(node->left);
            if (node->right) st.push(node->right);
        }
        reverse(result.begin(), result.end());
        return result;
    }
};
```

![image-20220516071259749](coderandom.assets/image-20220516071259749.png)

### 20220517 8.4 统一的迭代遍历

中序

```c++
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        if (root != NULL) st.push(root);
      	while (!st.empty()) {
        	TreeNode* node = st.top();
          if (node != NULL) {
            st.pop();
            if (node->right) st.push(node->right);
            st.push(node);
            st.push(NULL);
            if (node->left) st.push(node->left);
          } else {
            st.pop();
            node = st.top();
            st.pop();
            result.push_back(node->val);
          }
      	}
      return result;
    }
};
```

前序

```c++
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        if (root != NULL) st.push(root);
      	while (!st.empty()) {
        	TreeNode* node = st.top();
          if (node != NULL) {
            st.pop();
            if (node->right) st.push(node->right);
            if (node->left) st.push(node->left);
            st.push(node);
            st.push(NULL);
          } else {
            st.pop();
            node = st.top();
            st.pop();
            result.push_back(node->val);
          }
      	}
      return result;
    }
};
```

中序

```c++
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        if (root != NULL) st.push(root);
      	while (!st.empty()) {
        	TreeNode* node = st.top();
          if (node != NULL) {
            st.pop();
            st.push(node);
            st.push(NULL);
            if (node->right) st.push(node->right);
            if (node->left) st.push(node->left);
          } else {
            st.pop();
            node = st.top();
            st.pop();
            result.push_back(node->val);
          }
      	}
      return result;
    }
};
```

### 20220518 8.5 二叉树的层序遍历

```c++
class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        queue<TreeNode*> que;
        if (root != NULL) que.push(root);
        vector<vector<int>> result;
        while (!que.empty()) {
            int size = que.size();
            vector<int> vec;
            for (int i = 0; i < size; ++i) {
                TreeNode* node = que.front();
                que.pop();
                vec.push_back(node->val);
                if (node->left) que.push(node->left);
                if (node->right) que.push(node->right);
            }
            result.push_back(vec);
        }
        return result;

    }
};
```

![image-20220518064819545](coderandom.assets/image-20220518064819545.png)

层序遍历还是比较简单的，建立一个队列即可。

### 20220519 8.6 [226. Invert Binary Tree](https://leetcode.cn/problems/invert-binary-tree/)

```c++
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        if (root == NULL) return root;
        stack<TreeNode*> st;
        st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            st.pop();
            swap(node->left, node->right);
            if(node->right) st.push(node->right);
            if(node->left) st.push(node->left);
        }
        return root;
    }
};
```

![image-20220519192947544](coderandom.assets/image-20220519192947544.png)

倒转二叉树时，只需要把前序遍历或者后序遍历的中间节点的处理部分更换成`swap(node->left, node->right);`

### 20220520 8.7 [101. Symmetric Tree](https://leetcode.cn/problems/symmetric-tree/)

```c++
class Solution {
public:
    bool compare(TreeNode* left, TreeNode* right) {
        if (left == NULL && right != NULL) return false;
        else if (left != NULL && right == NULL) return false;
        else if (left == NULL && right == NULL) return true;
        else if (left->val != right->val) return false;
        bool outside = compare(left->left, right->right);
        bool inside = compare(left->right, right->left);
        bool isSame = outside && inside;
        return isSame;
    }
    bool isSymmetric(TreeNode* root) {
        if (root == NULL) return true;
        return compare(root->left, root->right);
    }
};
```

![image-20220520072724825](coderandom.assets/image-20220520072724825.png)

这道题很有趣，需要后序遍历但是要注意判断是否相同要通过内层和外层综合判断。

### 20220523 8.8 [104. Maximum Depth of Binary Tree](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)

```c++
class Solution {
public:
    int getDepth(TreeNode* node) {
        if (node == NULL) {
            return 0;
        }
        int leftDepth = getDepth(node->left);
        int rightDepth = getDepth(node->right);
        int depth = 1 + max(leftDepth, rightDepth);
        return depth;
    }
    int maxDepth(TreeNode* root) {
        return getDepth(root);
    }
};
```

![image-20220523064638262](coderandom.assets/image-20220523064638262.png)

使用递归法：

1. 确定递归函数的参数和返回值：参数为传入二叉树的根节点，返回值为树的深度
2. 确定终止条件：如果为空节点，则返回0
3. 确定单层递归逻辑：先求左子树深度，再求右子树深度，最后取左右子树最大值+1.+1是因为算上当前中间节点。

### 20220524 8.9 [111. Minimum Depth of Binary Tree](https://leetcode.cn/problems/minimum-depth-of-binary-tree/)

```c++
class Solution {
public:
    int getDepth(TreeNode* node) {
        if (node == NULL) return 0;
        int leftDepth = getDepth(node->left);
        int rightDepth = getDepth(node->right);
        if (node->left == NULL && node->right != NULL) {
            return 1 + rightDepth;
        }
        if (node->left != NULL && node->right == NULL) {
            return 1 + leftDepth;
        }
        int result = 1 + min(leftDepth, rightDepth);
        return result;
    }
    int minDepth(TreeNode* root) {
        return getDepth(root);
    }
};
```

![image-20220524071851680](coderandom.assets/image-20220524071851680.png)

这道题和昨天的题目比较类似。

### 20220525 8.10 [110. Balanced Binary Tree](https://leetcode.cn/problems/balanced-binary-tree/)

```c++
class Solution {
public:
    int getDepth(TreeNode* node) {
        if (node == NULL) {
            return 0;
        }
        int leftDepth = getDepth(node->left);
        if (leftDepth == -1) return -1;
        int rightDepth = getDepth(node->right);
        if (rightDepth == -1) return -1;
        return abs(leftDepth - rightDepth) > 1 ? -1 : 1 + max(leftDepth, rightDepth);
    }
    bool isBalanced(TreeNode* root) {
        return getDepth(root) == -1 ? false : true;
    }
};
```

![image-20220525065230677](coderandom.assets/image-20220525065230677.png)

原理还是要求出二叉树的高度和最小深度来比较。

1. 明确递归函数的参数和返回值

   参数为传入的节点指针，返回值是传入节点为根节点的二叉树高度

   如果左右子树的差值大于1，那就不是平衡二叉树的，直接返回-1

2. 明确终止条件：递归的过程中遇到空节点即种植。返回零，表示该节点树的高度为0

3. 明确单层递归逻辑：分别求出左右子树的高度，如果差值小于等于1，则返回二叉树的高度。否则返回-1.

### 20220526 8.11 [257. Binary Tree Paths](https://leetcode.cn/problems/binary-tree-paths/)

```c++
class Solution {
public:
    void traversal(TreeNode* cur, vector<int>& path, vector<string>& result) {
        path.push_back(cur->val);
        if (cur->left == NULL && cur->right == NULL) {
            string sPath;
            for (int i = 0; i < path.size() - 1; i++) {
                sPath += to_string(path[i]);
                sPath += "->";
            }
            sPath += to_string(path[path.size() - 1]);
            result.push_back(sPath);
            return;
        }
        if (cur->left) {
            traversal(cur->left, path, result);
            path.pop_back();
        }
        if (cur->right) {
            traversal(cur->right, path, result);
            path.pop_back();
        }
    }
    vector<string> binaryTreePaths(TreeNode* root) {
        vector<string> result;
        vector<int> path;
        if (root == NULL) return result;
        traversal(root, path, result);
        return result;

    }
};
```

![image-20220526065625041](coderandom.assets/image-20220526065625041.png)

### 20220527 8.12 [112. Path Sum](https://leetcode.cn/problems/path-sum/)

```c++
class Solution {
public:
    bool traversal(TreeNode* cur, int count) {
        if (!cur->left && !cur->right && count == 0) return true;
        if (!cur->left && !cur->right) return false;
        if (cur->left) {
            count -= cur->left->val;
            if (traversal(cur->left, count)) return true;
            count += cur->left->val;
        }
        if (cur->right) {
            count -= cur->right->val;
            if (traversal(cur->right, count)) return true;
            count += cur->right->val;
        }
        return false;
    }
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (root == NULL) return false;
        return traversal(root, targetSum - root->val);

    }
};
```

![image-20220527065843148](coderandom.assets/image-20220527065843148.png)

使用回溯递归法。

### 20220530 8.13 [106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

```c++
class Solution {
public:
    TreeNode* traversal (vector<int>& inorder, vector<int>& postorder) {
        if (postorder.size() == 0) return NULL;
        int rootValue = postorder[postorder.size() - 1];
        TreeNode* root = new TreeNode(rootValue);
        if (postorder.size() == 1) return root;

        int delimiterIndex;
        for (delimiterIndex = 0; delimiterIndex < inorder.size(); delimiterIndex++) {
            if (inorder[delimiterIndex] == rootValue) break;
        }
        vector<int> leftInorder(inorder.begin(), inorder.begin() + delimiterIndex);
        vector<int> rightInorder(inorder.begin() + delimiterIndex + 1, inorder.end());
        postorder.resize(postorder.size() - 1);
        vector<int> leftPostorder(postorder.begin(), postorder.begin() + leftInorder.size());
        vector<int> rightPostorder(postorder.begin() + leftInorder.size(), postorder.end());
        root->left = traversal(leftInorder, leftPostorder);
        root->right = traversal(rightInorder, rightPostorder);
        return root;
    }
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        if (inorder.size() == 0 || postorder.size() == 0) return NULL;
        return traversal(inorder, postorder);
    }
};
```

![image-20220530070734489](coderandom.assets/image-20220530070734489.png)

原理：以 后序数组的最后一个元素为切割点，先切割中序数组，再根据中序数组，反过来再切割后序数组。一层一层切，每次后序数组的最后一个元素就是节点元素。

说到一层一层切，就想到了递归

第一步：如果数组长度为0，则说明是空节点

第二步：如果数组不为空，那么将后序数组的最后一个元素作为节点元素

第三步：找到后序数组的最后一个元素在中序数组中的位置并将其作为切割点

第四步：切割中序数组，切成中序左数组和中序右数组

第五步：切割后序数组，切成后序左数组和后序右数组

第六步：递归处理左区间和右区间

### 20220531 8.14 [617. Merge Two Binary Trees](https://leetcode.cn/problems/merge-two-binary-trees/)

```c++
class Solution {
public:
    TreeNode* mergeTrees(TreeNode* root1, TreeNode* root2) {
        if (root1 == NULL) return root2;
        if (root2 == NULL) return root1;
        root1->val += root2->val;
        root1->left = mergeTrees(root1->left, root2->left);
        root1->right = mergeTrees(root1->right, root2->right);
        return root1;
    }
};
```

![image-20220531064623565](coderandom.assets/image-20220531064623565.png)

### 20220601 8.15 [700. Search in a Binary Search Tree](https://leetcode.cn/problems/search-in-a-binary-search-tree/)

```c++
class Solution {
public:
    TreeNode* searchBST(TreeNode* root, int val) {
        if (root == NULL || root->val == val) return root;
        if (root->val > val) return searchBST(root->left, val);
        if (root->val < val) return searchBST(root->right, val);
        return NULL;
    }
};
```

![image-20220601064656139](coderandom.assets/image-20220601064656139.png)

第五行和第六行直接return了，因为如果不return就变成了遍历整个树了。

### 20220602 8.16 [98. Validate Binary Search Tree](https://leetcode.cn/problems/validate-binary-search-tree/)

```c++
class Solution {
public:
    vector<int> vec;
    void traversal(TreeNode* root) {
        if (root == NULL) return;
        traversal(root->left);
        vec.push_back(root->val);
        traversal(root->right);
    }
    bool isValidBST(TreeNode* root) {
        vec.clear();
        traversal(root);
        for (int i = 1; i < vec.size(); i++) {
            if (vec[i] <= vec[i-1]) {
                return false;
            }
        }
        return true;
    }
};
```

![image-20220602065357314](coderandom.assets/image-20220602065357314.png)

### 20220603 8.17 [530. Minimum Absolute Difference in BST](https://leetcode.cn/problems/minimum-absolute-difference-in-bst/)

```c++
class Solution {
public:
    vector<int> vec;
    void traversal(TreeNode* root) {
        if (root == NULL) return;
        traversal(root->left);
        vec.push_back(root->val);
        traversal(root->right);
    }
    int getMinimumDifference(TreeNode* root) {
        vec.clear();
        traversal(root);
        if (vec.size() < 2) return 0;
        int result = INT_MAX;
        for (int i = 1; i < vec.size(); i++) {
            result = min(result, vec[i] - vec[i-1]);
        }
        return result;
    }
};
```

![image-20220603065656694](coderandom.assets/image-20220603065656694.png)

### 20220606 8.18 [501. Find Mode in Binary Search Tree](https://leetcode.cn/problems/find-mode-in-binary-search-tree/)

```c++
class Solution {
public:
    int maxCount;
    int count;
    TreeNode* pre;
    vector<int> result;
    void searchBST(TreeNode* cur) {
        if (cur == NULL) return;
        searchBST(cur->left);
        if(pre == NULL) {
            count = 1;
        } else if (pre->val == cur->val) {
            count++;
        } else {
            count = 1;
        }
        pre = cur;
        if (count == maxCount) {
            result.push_back(cur->val);
        }
        if (count > maxCount) {
            maxCount = count;
            result.clear();
            result.push_back(cur->val);
        }
        searchBST(cur->right);
        return;
    }
    vector<int> findMode(TreeNode* root) {
        count = 0;
        maxCount = 0;
        TreeNode* pre = NULL;
        result.clear();
        searchBST(root);
        return result;
    }
};
```

![image-20220606080524069](coderandom.assets/image-20220606080524069.png)

### 20220607 8.19 [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)

```c++
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == q || root == p || root == NULL) return root;
        TreeNode* left = lowestCommonAncestor(root->left, p, q);
        TreeNode* right = lowestCommonAncestor(root->right, p, q);
        if (left != NULL && right != NULL) return root;
        if (left == NULL && right != NULL) return right;
        else if (left != NULL && right == NULL) return left;
        else return NULL;
    }
};
```

![image-20220607064722799](coderandom.assets/image-20220607064722799.png)

```c++
class Solution {
public:
    TreeNode* traversal(TreeNode* cur, TreeNode* p, TreeNode* q) {
        if (cur == NULL) return cur;
        if (cur->val > p->val && cur->val > q->val) {
            TreeNode* left = traversal(cur->left, p, q);
            if (left != NULL) {
                return left;
            }
        }
        if (cur->val < p->val && cur->val < q->val) {
            TreeNode* right = traversal(cur->right, p, q);
            if (right != NULL) {
                return right;
            }
        }
        return cur;
    }
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        return traversal(root, p, q); 
    }
};
```

![image-20220607065343469](coderandom.assets/image-20220607065343469.png)

### 20220608 8.20 [701. Insert into a Binary Search Tree](https://leetcode.cn/problems/insert-into-a-binary-search-tree/)

```c++
class Solution {
public:
    TreeNode* parent;
    void traversal(TreeNode* cur, int val) {
        if (cur == NULL) {
            TreeNode* node = new TreeNode(val);
            if (val > parent->val) parent->right = node;
            else parent->left = node;
            return;
        }
        parent = cur;
        if (cur->val > val) traversal(cur->left, val);
        if (cur->val < val) traversal(cur->right, val);
        return;
    }
    TreeNode* insertIntoBST(TreeNode* root, int val) {
        parent = new TreeNode(0);
        if (root == NULL) {
            root = new TreeNode(val);
        }
        traversal(root, val);
        return root;
    }
};
```

![image-20220608065845053](coderandom.assets/image-20220608065845053.png)

### 20220609 8.21 [450. Delete Node in a BST](https://leetcode.cn/problems/delete-node-in-a-bst/)

```c++
class Solution {
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        if (root == nullptr) return root;
        if (root->val == key) {
            if (root->left == nullptr) return root->right;
            else if (root->right == nullptr) return root->left;
            else {
                TreeNode* cur = root->right;
                while (cur->left != nullptr) {
                    cur = cur->left;
                }
                cur->left = root->left;
                TreeNode* tmp = root;
                root = root->right;
                delete tmp;
                return root; 
            }
        }
        if (root->val > key) root->left = deleteNode(root->left, key);
        if (root->val < key) root->right = deleteNode(root->right, key);
        return root;
    }
};
```

![image-20220609064726939](coderandom.assets/image-20220609064726939.png)

### 20220610 8.22 [669. Trim a Binary Search Tree](https://leetcode.cn/problems/trim-a-binary-search-tree/)

```c++
class Solution {
public:
    TreeNode* trimBST(TreeNode* root, int low, int high) {
        if (root == nullptr) return nullptr;
        if (root->val < low) {
            TreeNode* right = trimBST(root->right, low, high);
            return right;
        }
        if (root->val > high) {
            TreeNode* left = trimBST(root->left, low, high);
            return left;
        }
        root->left = trimBST(root->left, low, high);
        root->right = trimBST(root->right, low, high);
        return root;
    }
};
```

![image-20220610065112210](coderandom.assets/image-20220610065112210.png)

### 20220613 8.23 [108. Convert Sorted Array to Binary Search Tree](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/)

```c++
class Solution {
public:
    TreeNode* traversal(vector<int>& nums, int left, int right) {
        if (left > right) return nullptr;
        int mid = left +((right - left) / 2);
        TreeNode* root = new TreeNode(nums[mid]);
        root->left = traversal(nums, left, mid - 1);
        root->right = traversal(nums, mid + 1, right);
        return root;
    }
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        TreeNode* root = traversal(nums, 0, nums.size() - 1);
        return root;
    }
};
```

![image-20220613065121345](coderandom.assets/image-20220613065121345.png)

### 20220614 9.1 9.2 [77. Combinations](https://leetcode.cn/problems/combinations/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(int n, int k, int startIndex) {
        if (path.size() == k) {
            result.push_back(path);
            return;
        }
        for (int i = startIndex; i <= n - (k - path.size()) + 1; i++) {
            path.push_back(i);
            backtracking(n, k, i + 1);
            path.pop_back();
        }
    }
public:
    vector<vector<int>> combine(int n, int k) {
        backtracking(n, k, 1);
        return result;
    }
};
```

![image-20220614070256901](coderandom.assets/image-20220614070256901.png)

### 20220615 9.3 [216. Combination Sum III](https://leetcode.cn/problems/combination-sum-iii/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(int targetSum, int k, int sum, int startIndex) {
        if (path.size() == k) {
            if (sum == targetSum) result.push_back(path);
            return;
        }
        for (int i = startIndex; i <= 9; i++) {
            sum += i;
            path.push_back(i);
            backtracking(targetSum, k, sum, i + 1);
            sum -= i;
            path.pop_back();
        }
    }
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        result.clear();
        path.clear();
        backtracking(n, k, 0, 1);
        return result;
    }
};
```

![image-20220615065905043](coderandom.assets/image-20220615065905043.png)

### 20220616 9.4 [17. Letter Combinations of a Phone Number](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)

```c++
class Solution {
private:
    const string letterMap[10] = {
        "",
        "",
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz",
    };
public:
    vector<string> result;
    string s;
    void backtracking(const string& digits, int index) {
        if (index == digits.size()) {
            result.push_back(s);
            return;
        }
        int digit = digits[index] - '0';
        string letters =letterMap[digit];
        for (int i = 0; i < letters.size(); i++) {
            s.push_back(letters[i]);
            backtracking(digits, index + 1);
            s.pop_back();
        }
    }
    vector<string> letterCombinations(string digits) {
        s.clear();
        result.clear();
        if (digits.size() == 0) {
            return result;
        }
        backtracking(digits, 0);
        return result;
    }
};
```

![image-20220616075740100](coderandom.assets/image-20220616075740100.png)

### 20220617 9.5 [39. Combination Sum](https://leetcode.cn/problems/combination-sum/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(vector<int>& candidates, int target, int sum, int startIndex) {
        if (sum == target) {
            result.push_back(path);
            return;
        }
        for (int i = startIndex; i < candidates.size() && sum + candidates[i] <= target; i++) {
            sum += candidates[i];
            path.push_back(candidates[i]);
            backtracking(candidates, target, sum, i);
            sum -= candidates[i];
            path.pop_back();
        }
    }
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        result.clear();
        path.clear();
        sort(candidates.begin(), candidates.end());
        backtracking(candidates, target, 0, 0);
        return result;
    }
};
```

![image-20220617064714694](coderandom.assets/image-20220617064714694.png)

### 20220620 9.6 [40. Combination Sum II](https://leetcode.cn/problems/combination-sum-ii/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(vector<int>& candidates, int target, int sum, int startIndex, vector<bool>& used) {
        if (sum == target) {
            result.push_back(path);
            return;
        }
        for (int i = startIndex; i < candidates.size() && sum + candidates[i] <= target; i++) {
            if (i > 0 && candidates[i] == candidates[i - 1] && used[i - 1] == false) {
                continue;
            }
            sum += candidates[i];
            path.push_back(candidates[i]);
            used[i] = true;
            backtracking(candidates, target, sum, i + 1, used);
            used[i] = false;
            sum -= candidates[i];
            path.pop_back();
        }
    }
public:
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        vector<bool> used(candidates.size(), false);
        path.clear();
        result.clear();
        sort(candidates.begin(), candidates.end());
        backtracking(candidates, target, 0, 0, used);
        return result;
    }
};
```

![image-20220620065221811](coderandom.assets/image-20220620065221811.png)

### 20220621 9.7 [131. Palindrome Partitioning](https://leetcode.cn/problems/palindrome-partitioning/)

```c++
class Solution {
private:
    vector<vector<string>> result;
    vector<string> path;
    void backtracking (const string& s, int startIndex) {
        if (startIndex >= s.size()) {
            result.push_back(path);
            return;
        }
        for (int i = startIndex; i < s.size(); i++) {
            if (isPalindrome(s, startIndex, i)) {
                string str = s.substr(startIndex, i - startIndex + 1);
                path.push_back(str);
            } else {
                continue;
            }
            backtracking(s, i + 1);
            path.pop_back();
        }
    }
    bool isPalindrome(const string& s, int start, int end) {
        for (int i = start, j = end; i < j; i++, j--) {
            if (s[i] != s[j]) {
                return false;
            }
        }
        return true;
    }
public:
    vector<vector<string>> partition(string s) {
        result.clear();
        path.clear();
        backtracking(s, 0);
        return result;
    }
};
```

![image-20220621065750508](coderandom.assets/image-20220621065750508.png)

### 20220622 9.8 [93. Restore IP Addresses](https://leetcode.cn/problems/restore-ip-addresses/)

```c++
class Solution {
private:
    vector<string> result;
    void backtracking(string& s, int startIndex, int pointNum) {
        if (pointNum == 3) {
            if (isValid(s, startIndex, s.size() - 1)) {
                result.push_back(s);
            }
            return;
        }
        for (int i = startIndex; i < s.size(); i++) {
            if (isValid(s, startIndex, i)) {
                s.insert(s.begin() + i + 1, '.');
                pointNum++;
                backtracking(s, i + 2, pointNum);
                pointNum--;
                s.erase(s.begin() + i + 1);
            } else break;
        }
    }
    bool isValid(const string& s, int start, int end) {
        if (start > end) return false;
        if (s[start] == '0' && start != end) return false;
        int num = 0;
        for (int i = start; i <= end; i++) {
            if (s[i] > '9' || s[i] < '0') return false;
            num = num * 10 + (s[i] - '0');
            if (num > 255) return false;
        }
        return true;
    }
public:
    vector<string> restoreIpAddresses(string s) {
        result.clear();
        if (s.size() > 12) return result;
        backtracking(s, 0, 0);
        return result;
    }
};
```



![image-20220622070634519](coderandom.assets/image-20220622070634519.png)

### 20220623 9.9 [78. Subsets](https://leetcode.cn/problems/subsets/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(vector<int>& nums, int startIndex) {
        result.push_back(path);
        if (startIndex >= nums.size()) return;
        for (int i = startIndex; i < nums.size(); i++) {
            path.push_back(nums[i]);
            backtracking(nums, i + 1);
            path.pop_back();
        }
    }
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        result.clear();
        path.clear();
        backtracking(nums, 0);
        return result;
    }
};
```

![image-20220623050737658](coderandom.assets/image-20220623050737658.png)

### 20220624 9.10 [90. Subsets II](https://leetcode.cn/problems/subsets-ii/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(vector<int>& nums, int startIndex, vector<bool>& used) {
        result.push_back(path);
        for (int i = startIndex; i < nums.size(); i++) {
            if (i > 0 && nums[i] == nums[i - 1] && used[i - 1] == false) {
                continue;
            }
            path.push_back(nums[i]);
            used[i] = true;
            backtracking(nums, i + 1, used);
            used[i] = false;
            path.pop_back();
        }
    }
public:
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        result.clear();
        path.clear();
        vector<bool> used(nums.size(), false);
        sort(nums.begin(), nums.end());
        backtracking(nums, 0, used);
        return result;
    }
};
```

![image-20220624111407312](coderandom.assets/image-20220624111407312.png)

### 20220627 9.11 [491. Increasing Subsequences](https://leetcode.cn/problems/increasing-subsequences/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(vector<int>& nums, int startIndex) {
        if (path.size() > 1) {
            result.push_back(path);
        }
        int used[201] = {0};
        for (int i = startIndex; i < nums.size(); i++) {
            if ((!path.empty() && nums[i] < path.back())
                || used[nums[i] + 100] == 1) {
                continue;
            }
            used[nums[i] + 100] = 1;
            path.push_back(nums[i]);
            backtracking(nums, i + 1);
            path.pop_back();
        }
    }
public:
    vector<vector<int>> findSubsequences(vector<int>& nums) {
        result.clear();
        path.clear();
        backtracking(nums, 0);
        return result;
    }
};
```

![image-20220627083613235](coderandom.assets/image-20220627083613235.png)

### 20220628 9.12 [46. Permutations](https://leetcode.cn/problems/permutations/)

```c++
class Solution {
public:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking(vector<int>& nums, vector<bool>& used) {
        if (path.size() == nums.size()) {
            result.push_back(path);
            return;
        }
        for (int i = 0; i < nums.size(); i++) {
            if (used[i] == true) continue;
            used[i] = true;
            path.push_back(nums[i]);
            backtracking(nums, used);
            path.pop_back();
            used[i] = false;
        }
    }
    vector<vector<int>> permute(vector<int>& nums) {
        result.clear();
        path.clear();
        vector<bool> used(nums.size(), false);
        backtracking(nums, used);
        return result;
    }
};
```

![image-20220628084255963](coderandom.assets/image-20220628084255963.png)

### 20220629 9.13 [47. Permutations II](https://leetcode.cn/problems/permutations-ii/)

```c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;
    void backtracking (vector<int>& nums, vector<bool>& used) {
        if (path.size() == nums.size()) {
            result.push_back(path);
            return;
        }
        for (int i = 0; i < nums.size(); i++) {
            if (i > 0 && nums[i] == nums[i - 1] && used[i - 1] == false) {
                continue;
            }
            if (used[i] == false) {
                used[i] = true;
                path.push_back(nums[i]);
                backtracking(nums, used);
                path.pop_back();
                used[i] = false;
            }
        }
    }
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        result.clear();
        path.clear();
        sort(nums.begin(), nums.end());
        vector<bool> used(nums.size(), false);
        backtracking(nums, used);
        return result;
    }
};
```

![image-20220629105227420](coderandom.assets/image-20220629105227420.png)

### 20220630 9.14 [51. N-Queens](https://leetcode.cn/problems/n-queens/)

```c++
class Solution {
private:
    vector<vector<string>> result;
    void backtracking(int n, int row, vector<string>& chessboard) {
        if (row == n) {
            result.push_back(chessboard);
            return;
        }
        for (int col = 0; col < n; col++) {
            if (isValid(row, col, chessboard, n)) {
                chessboard[row][col] = 'Q';
                backtracking(n, row + 1, chessboard);
                chessboard[row][col] = '.';
            }
        }
    }
    bool isValid(int row, int col, vector<string>& chessboard, int n) {
        int count = 0;
        for (int i = 0; i < row; i++) {
            if (chessboard[i][col] == 'Q') {
                return false;
            }
        }
        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
            if (chessboard[i][j] == 'Q') {
                return false;
            }
        }
        for (int i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++) {
            if (chessboard[i][j] == 'Q') {
                return false;
            }
        }
        return true;
    }
public:
    vector<vector<string>> solveNQueens(int n) {
        result.clear();
        vector<string> chessboard(n, string(n, '.'));
        backtracking(n, 0, chessboard);
        return result;
    }
};
```

![image-20220630101004892](coderandom.assets/image-20220630101004892.png)

### 20220701 9.15 [37. Sudoku Solver](https://leetcode.cn/problems/sudoku-solver/)

```c++
class Solution {
private:
    bool backtracking(vector<vector<char>>& board) {
        for (int i = 0; i < board.size(); i++) {
            for (int j = 0; j < board[0].size(); j++) {
                if (board[i][j] != '.') continue;
                for (char k = '1'; k <= '9'; k++) {
                    if (isValid(i, j, k, board)) {
                        board[i][j] = k;
                        if (backtracking(board)) return true;
                        board[i][j] = '.';
                    }
                }
                return false;
            }
        }
        return true;
    }
    bool isValid(int row, int col, char val, vector<vector<char>>& board) {
        for (int i = 0; i < 9; i++) {
            if (board[row][i] == val) {
                return false;
            }
        }
        for (int j = 0; j < 9; j++) {
            if (board[j][col] == val) {
                return false;
            }
        }
        int startRow = (row / 3) * 3;
        int startCol = (col / 3) * 3;
        for (int i = startRow; i < startRow + 3; i++) {
            for (int j = startCol; j < startCol + 3; j++) {
                if (board[i][j] == val) {
                    return false;
                }
            }
        }
        return true;
    }
public:
    void solveSudoku(vector<vector<char>>& board) {
        backtracking(board);
    }
};
```

![image-20220701091203662](coderandom.assets/image-20220701091203662.png)

### 20220704 10.2 [455. Assign Cookies](https://leetcode.cn/problems/assign-cookies/)

```c+++
class Solution {
public:
    int findContentChildren(vector<int>& g, vector<int>& s) {
        sort(g.begin(), g.end());
        sort(s.begin(), s.end());
        int index = s.size() - 1;
        int result = 0;
        for (int i = g.size() - 1; i >= 0; i--) {
            if (index >= 0 && s[index] >= g[i]) {
                result++;
                index--;
            }
        }
        return result;

    }
};
```

![image-20220704083633139](coderandom.assets/image-20220704083633139.png)

### 20220705 10.3 [376. Wiggle Subsequence](https://leetcode.cn/problems/wiggle-subsequence/)

```c++
class Solution {
public:
    int wiggleMaxLength(vector<int>& nums) {
        if (nums.size() <= 1) return nums.size();
        int curDiff = 0;
        int preDiff = 0;
        int result = 1;
        for (int i = 0; i < nums.size() - 1; i++) {
            curDiff = nums[i + 1] - nums[i];
            if ((curDiff > 0 && preDiff <= 0) || (preDiff >= 0 && curDiff < 0)) {
                result++;
                preDiff = curDiff;
            }
        }
        return result;
    }
};
```

![image-20220705082854839](coderandom.assets/image-20220705082854839.png)

### 20220706 10.4 [53. Maximum Subarray](https://leetcode.cn/problems/maximum-subarray/)

```c++
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int curSum = nums[0];
        int maxSum = nums[0];
        for(int i = 1; i < nums.size(); ++i){
            curSum = max(nums[i], curSum + nums[i]);
            maxSum = max(curSum, maxSum);
        }
        return maxSum;
    }
};
```

![image-20220706065158059](coderandom.assets/image-20220706065158059.png)

### 20220711 10.5 [122. Best Time to Buy and Sell Stock II](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)

```c++
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int result = 0;
        for (int i = 1; i < prices.size(); i++) {
            result += max(prices[i] - prices[i - 1], 0);
        }
        return result;
    }
};
```

![image-20220711064932984](coderandom.assets/image-20220711064932984.png)

### 20220711 10.6 [55. Jump Game](https://leetcode.cn/problems/jump-game/)

```c++
class Solution {
public:
    bool canJump(vector<int>& nums) {
        int cover = 0;
        if (nums.size() == 1) return true;
        for(int i = 0; i <= cover; i++) {
            cover = max(i + nums[i], cover);
            if (cover >= nums.size() - 1) return true;
        }
        return false;
    }
};
```

![image-20220711070141659](coderandom.assets/image-20220711070141659.png)

### 20220711 10.7 [45. Jump Game II](https://leetcode.cn/problems/jump-game-ii/)

```c++
class Solution {
public:
    int jump(vector<int>& nums) {
        int curDistance = 0;
        int ans = 0;
        int nextDistance = 0;
        for (int i = 0; i <nums.size() - 1; i++) {
            nextDistance = max(nums[i] + i, nextDistance);
            if (i == curDistance) {
                curDistance = nextDistance;
                ans++;
            }
        }
        return ans;
    }
};
```

![image-20220711070614322](coderandom.assets/image-20220711070614322.png)

### 20220712 10.8 [134. Gas Station](https://leetcode.cn/problems/gas-station/)

```c++
class Solution {
public:
    int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
        int curSum = 0;
        int totalSum = 0;
        int start = 0;
        for (int i = 0; i < gas.size(); i++) {
            curSum += gas[i] - cost[i];
            totalSum += gas[i] - cost[i];
            if (curSum < 0) {
                start = i + 1;
                curSum = 0;
            }
        }
        if (totalSum < 0) return -1;
        return start;
    }
};
```

![image-20220712065307172](coderandom.assets/image-20220712065307172.png)

### 20220713 10.9 [135. Candy](https://leetcode.cn/problems/candy/)

```c++
class Solution {
public:
    int candy(vector<int>& ratings) {
        vector<int> candyVec(ratings.size(), 1);
        for (int i = 1; i < ratings.size(); i++) {
            if (ratings[i] > ratings[i - 1]) candyVec[i] = candyVec[i - 1] + 1;
        }
        for (int i = ratings.size() - 2; i >= 0; i--) {
            if (ratings[i] > ratings[i + 1]) {
                candyVec[i] = max(candyVec[i], candyVec[i + 1] + 1);
            }
        }
        int result = 0;
        for (int i = 0; i < candyVec.size(); i++) result += candyVec[i];
        return result;
    }
};
```

![image-20220713065754140](coderandom.assets/image-20220713065754140.png)

### 20220714 10.10 [860. Lemonade Change](https://leetcode.cn/problems/lemonade-change/)

```c++
class Solution {
public:
    bool lemonadeChange(vector<int>& bills) {
        int five = 0, ten = 0;
        for (int bill : bills) {
            if (bill == 5) {
                five++;
            }
            if (bill == 10) {
                if (five <= 0) return false;
                ten++;
                five--;
            }
            if (bill == 20) {
                if (five > 0 && ten > 0) {
                    five--;
                    ten--;
                } else if (five >= 3) {
                    five -= 3;
                } else return false;
            }
        }
        return true;
    }
};
```

![image-20220714065153553](coderandom.assets/image-20220714065153553.png)

### 20220715 10.11 [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/)

```c++
class Solution {
private:
    static bool cmp(const vector<int>& a, const vector<int>& b) {
        return a[0] < b[0];
    }
public:
    int findMinArrowShots(vector<vector<int>>& points) {
        if (points.size() == 0) return 0;
        sort(points.begin(), points.end(), cmp);
        int result = 1;
        for (int i = 1; i < points.size(); i++) {
            if (points[i][0] > points[i - 1][1]) {
                result++;
            } else {
                points[i][1] = min(points[i - 1][1], points[i][1]);
            }
        }
        return result;
    }
};
```

![image-20220715074236491](coderandom.assets/image-20220715074236491.png)

### 20220718 10.12 [56. Merge Intervals](https://leetcode.cn/problems/merge-intervals/)

```c++
class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector<vector<int>> result;
        if (intervals.size() == 0) return result;
        sort(intervals.begin(), intervals.end(), [](const vector<int>& a, const vector<int>& b) {
            return a[0] < b[0];
        });

        result.push_back(intervals[0]);
        for (int i = 1; i < intervals.size(); i++) {
            if (result.back()[1] >= intervals[i][0]) {
                result.back()[1] = max(result.back()[1], intervals[i][1]);
            } else {
                result.push_back(intervals[i]);
            }
        }
        return result;
    }
};
```

![image-20220718075833916](coderandom.assets/image-20220718075833916.png)

### 20220719 10.13 [738. Monotone Increasing Digits](https://leetcode.cn/problems/monotone-increasing-digits/)

```c++
class Solution {
public:
    int monotoneIncreasingDigits(int n) {
        string strNum = to_string(n);
        int flag = strNum.size();
        for (int i = strNum.size() - 1; i > 0; i--) {
            if (strNum[i - 1] > strNum[i]) {
                flag = i;
                strNum[i - 1]--;
            }
        }
        for (int i = flag; i < strNum.size(); i++) {
            strNum[i] = '9';
        }
        return stoi(strNum);
    }
};
```

![image-20220719070458775](coderandom.assets/image-20220719070458775.png)

### 20220710 11.1 11.2 [509. Fibonacci Number](https://leetcode.cn/problems/fibonacci-number/)

```c++
class Solution {
public:
    int fib(int n) {
        if (n < 2) return n;
        return fib(n - 1) + fib(n - 2);
    }
};
```

![image-20220720074702641](coderandom.assets/image-20220720074702641.png)
