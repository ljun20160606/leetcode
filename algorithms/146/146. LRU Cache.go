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
