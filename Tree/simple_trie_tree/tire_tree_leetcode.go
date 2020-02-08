package leetcodetrie

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

func NewTrie() *Trie {
	t := &Trie{}
	t.Root = NewTrieNode()
	return t
}

func NewTrieNode() *TrieNode {
	n := &TrieNode{}
	n.Children = make(map[rune]*TrieNode)
	n.End = false
	return n
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.Root = NewTrieNode()
	return t
}

func (t *Trie) Insert(word string) {
	if len(word) < 1 {
		return
	}
	chars := []rune(word)
	slen := len(chars)
	node := t.Root
	for i := 0; i < slen; i++ { // 将word的每个词分割,作为key, 遍历word
		if _, exist := node.Children[chars[i]]; !exist {
			node.Children[chars[i]] = NewTrieNode() // 如果不存在则用这个key新建结点
		}
		node = node.Children[chars[i]] // 将node指针 移到子节点上
	}
	node.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	chars := []rune(word)
	slen := len(chars)
	node := this.Root
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; !exists { // 有字符不匹配
			return false
		}
		node = node.Children[chars[i]]       // 将node指针移到子节点上
		if node.End == true && i == slen-1 { // 执行到最后
			return true
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	chars := []rune(prefix)
	slen := len(chars)
	node := this.Root
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; !exists {
			return false // 只要有一个不匹配 则说明不是prefix
		}
		node = node.Children[chars[i]]
	}
	return true
}
