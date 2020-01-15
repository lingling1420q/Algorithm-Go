package trietree

type Trie struct {
	node   map[rune]*Trie // 节点是一个map存储结构 这样是否会有一个问题，占用内存太大?
	val    string         // 最后一个节点存储这个Trie分支表示的word
	isLast bool
}

func Constructor() Trie {
	return Trie{node: make(map[rune]*Trie, 26), isLast: false}
}

func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.node[v] == nil { //没找到该节点
			vt := &Trie{node: make(map[rune]*Trie, 26), isLast: false}
			t.node[v] = vt
		}
		t = t.node[v] // 找到节点 跳到下一个节点
	}
	t.val = word
	t.isLast = true
}

func (t *Trie) Search(word string) bool {
	for _, v := range word {
		if t.node[v] == nil { // 没有找到该节点
			return false
		}
		t = t.node[v] // 找到节点，跳到下一个节点
	}
	return t.isLast // 如果是最后一个节点  说明word在这个Trie的一个分支上
}

// 这是前缀search，不需要到最后一个节点，只要能匹配对应的前缀说明成功
func (t *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if t.node[v] == nil { //没找到该节点
			return false
		}
		t = t.node[v] //找到节点，跳到下一个节点
	}
	return true
}

// 查找数据
func (t *Trie) SearchNode(key string) (res []interface{}) {
	root := t
	for _, v := range key {
		if v, ok := root.node[v]; ok {
			root.node = v.node
		} else {
			return
		}
	}

	var queue []*Trie
	queue = append(queue, root)
	for len(queue) > 0 {
		var childQueue []*Trie
		for _, v := range queue { // 遍历这层的数据
			if v.isLast {
				res = append(res, v.val)
			}
			for _, vi := range v.node { //把下一次的数据进行遍历
				childQueue = append(childQueue, vi)
			}
		}
		queue = childQueue
	}
	return
}
