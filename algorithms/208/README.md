# Implement Trie (Prefix Tree)

## 描述

```txt

Implement a trie with insert, search, and startsWith methods.



Note:
You may assume that all inputs are consist of lowercase letters a-z.

```

## 题解

```go
package algorithms

//Implement a trie with insert, search, and startsWith methods.
//
//Note:
//You may assume that all inputs are consist of lowercase letters a-z.

type Trie struct {
	Children [26]*Trie
	End      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	trie := t
	for _, v := range word {
		node := &trie.Children[v-'a']
		if *node == nil {
			constructor := Constructor()
			*node = &constructor
		}
		trie = *node
	}
	trie.End = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	trie := t
	for _, v := range word {
		node := trie.Children[v-'a']
		if node == nil {
			return false
		}
		trie = node
	}
	return trie.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	trie := t
	for _, v := range prefix {
		node := trie.Children[v-'a']
		if node == nil {
			return false
		}
		trie = node
	}
	return true
}

```