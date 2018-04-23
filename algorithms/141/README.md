# Linked List Cycle

## 描述

```txt

Given a linked list, determine if it has a cycle in it.



Follow up:
Can you solve it without using extra space?

```

## 题解

```c++
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode *fast=head, *slow=head;
        while(fast != nullptr && fast->next != nullptr){
            slow = slow->next;
            fast = fast->next->next;
            if(slow == fast)
                return true;
        }
        return false;
    }
};
```