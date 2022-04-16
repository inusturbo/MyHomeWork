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
