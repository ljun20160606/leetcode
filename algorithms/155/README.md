# Min Stack

## 描述

```txt

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.


push(x) -- Push element x onto stack.


pop() -- Removes the element on top of the stack.


top() -- Get the top element.


getMin() -- Retrieve the minimum element in the stack.




Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.


```

## 题解

```go
package algorithms

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	min := x
	if len(m.stack) > 0 && m.GetMin() < x {
		min = m.GetMin()
	}
	m.stack = append(m.stack, item{min: min, x: x})
	return
}

func (m *MinStack) Pop() {
	length := len(m.stack)
	m.stack = m.stack[:length-1]
	return
}

func (m *MinStack) Top() int {
	length := len(m.stack)
	return m.stack[length-1].x
}

func (m *MinStack) GetMin() int {
	length := len(m.stack)
	return m.stack[length-1].min
}

```