# LRU Cache

## 描述

```txt

Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.



get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.


Follow up:
Could you do both operations in O(1) time complexity?

Example:
LRUCache cache = new LRUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4


```

## 题解

```go
package algorithms

import "container/list"

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (l *LRUCache) Get(key int) int {
	element, has := l.cache[key]
	if !has {
		return -1
	}
	l.list.MoveBefore(element, l.list.Front())
	return element.Value.([]int)[1]
}

func (l *LRUCache) Put(key int, value int) {
	element, has := l.cache[key]
	if has {
		element.Value = []int{key, value}
		l.list.MoveBefore(element, l.list.Front())
		return
	}
	if l.list.Len()+1 <= l.capacity {
		element := l.list.PushFront([]int{key, value})
		l.cache[key] = element
		return
	}
	back := l.list.Back()
	front := l.list.PushFront([]int{key, value})
	l.cache[key] = front
	l.list.Remove(back)
	delete(l.cache, back.Value.([]int)[0])
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

```